package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/go-redis/redis"
)

type Artist struct {
	ID       int       `json:"id"`
	GID      string    `json:"gid"`
	Name     string    `json:"name"`
	SortName string    `json:"sort_name"`
	Area     string    `json:"area"`
	Aliases  []*Aliase `json:"aliases"`
	Begin    *Begin    `json:"begin"`
	End      *End      `json:"end"`
	Tags     []*Tag    `json:"tags"`
	Rating   *Rating   `json:"rating"`
}

type Artists []*Artist

func (artists Artists) rpushTags(client *redis.Client) error {
	for _, artist := range artists {
		if len(artist.Tags) == 0 {
			continue
		}

		for _, tag := range artist.Tags {
			err := client.RPush(artist.Name, tag).Err()
			if err != nil {
				return fmt.Errorf("could not push the lists\n  %s", err)
			}
		}
	}
	return nil
}

type Aliase struct {
	Name     string `json:"name"`
	SortName string `json:"sort_name"`
}

type Begin struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Date  int `json:"date"`
}

type End struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Date  int `json:"date"`
}

type Tag struct {
	Count int    `json:"count"`
	Value string `json:"value"`
}

// MarshalBinary interface must be implemeted to encode the Tag type to json format and set as list-typed value:
// https://stackoverflow.com/questions/44771474/how-to-set-array-document-into-redis-in-golang
func (t Tag) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

type Rating struct {
	Count int `json:"count"`
	Value int `json:"value"`
}

func main() {
	var filepath, artistName string
	flag.StringVar(&filepath, "file", "", "specify a file path")
	flag.StringVar(&filepath, "f", "", "specify a file path")
	flag.StringVar(&artistName, "name", "", "specify the artist name")
	flag.StringVar(&artistName, "n", "", "specify the artist name")
	flag.Parse()

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open a file: %s\n  %s", filepath, err)
		os.Exit(1)
	}
	defer f.Close()

	artists, err := readJSON(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not construct object from JSON scheme: %s", err)
		os.Exit(1)
	}

	client, err := newRedisClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = artists.rpushTags(client)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tags, err := getTags(client, artistName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, tag := range tags {
		fmt.Printf("tag.count: %d, tag.value: %s\n", tag.Count, tag.Value)
	}
}

func getTags(client *redis.Client, artistName string) ([]Tag, error) {
	rets, err := client.LRange(artistName, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("could not get the lists\n  %s", err)
	}
	if err == redis.Nil {
		return nil, fmt.Errorf("could not find the value...%s", artistName)
	}

	var (
		tag  Tag
		tags []Tag
	)
	for _, ret := range rets {
		json.Unmarshal([]byte(ret), &tag)
		tags = append(tags, tag)
	}
	return tags, nil
}

func readJSON(r io.Reader) (Artists, error) {
	var artists []*Artist
	reader := bufio.NewReader(r)
	for {
		artist := Artist{}
		buf, readErr := reader.ReadBytes('\n')
		if (readErr != nil) && (readErr != io.EOF) {
			return nil, fmt.Errorf("could not read bytes: %s", readErr)
		}

		if err := json.Unmarshal(buf, &artist); err != nil && readErr != io.EOF {
			return nil, fmt.Errorf("could not parse json format: %s", err)
		}

		artists = append(artists, &artist)

		if readErr == io.EOF {
			break
		}
	}
	return artists, nil
}

func newRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("could not connect to redis\n  %s", err)
	}

	return client, nil
}
