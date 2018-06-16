package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	var artistName string
	flag.StringVar(&artistName, "name", "", "specify the artist name")
	flag.StringVar(&artistName, "n", "", "specify the artist name")
	flag.Parse()

	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to mongodb://localhost:\n  %s", err)
		os.Exit(1)
	}

	c := session.DB("MusicBrainz").C("artist")
	query := c.Find(bson.M{"name": artistName})
	artist, _, err := getQueryTime(query)
	fmt.Println(artist)
}

func readBSON(r io.Reader) (Artists, error) {
	var artists Artists
	reader := bufio.NewReader(r)
	for {
		artist := Artist{}
		buf, readErr := reader.ReadBytes('\n')
		if (readErr != nil) && (readErr != io.EOF) {
			return nil, fmt.Errorf("could not read json file:\n  %s", readErr)
		}

		if err := bson.UnmarshalJSON(buf, &artist); err != nil && readErr != io.EOF {
			return nil, fmt.Errorf("could not parse json format:\n  %s", err)
		}

		artists = append(artists, &artist)

		if readErr == io.EOF {
			break
		}
	}
	return artists, nil
}

func getQueryTime(query *mgo.Query) (*Artist, time.Duration, error) {
	artist := &Artist{}

	start := time.Now()
	if err := query.One(&artist); err != nil {
		return nil, 0, err
	}
	return artist, time.Now().Sub(start), nil
}
