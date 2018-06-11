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

type Rating struct {
	Count int `json:"count"`
	Value int `json:"value"`
}

func main() {
	var filepath string
	flag.StringVar(&filepath, "file", "", "specify a file path")
	flag.StringVar(&filepath, "f", "", "specify a file path")
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
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, artist := range artists {
		err = client.Set(artist.Name, artist.Area, 0).Err()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func readJSON(r io.Reader) ([]*Artist, error) {
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
