package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Artist struct {
	ID       int      `json:"id"`
	GID      string   `json:"gid"`
	Name     string   `json:"name"`
	SortName string   `json:"sort_name"`
	Area     string   `json:"area"`
	Aliases  []*Alias `json:"aliases"`
	Begin    *Begin   `json:"begin"`
	End      *End     `json:"end"`
	Tags     []*Tag   `json:"tags"`
	Rating   *Rating  `json:"rating"`
}

type Artists []*Artist

type Alias struct {
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
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/result/", resultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		r.ParseForm()
		fmt.Println(r.Form)
		return
	}
	t, err := template.ParseFiles("search.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find a specified file: %s", err)
		os.Exit(1)
	}
	t.Execute(w, nil)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := session.DB("MusicBrainz").C("artist")

	searchOpt := bson.M{}
	r.ParseForm()
	if r.Form.Get("name") != "" {
		searchOpt["name"] = r.Form.Get("name")
	}
	if r.Form.Get("alias") != "" {
		searchOpt["aliases.name"] = r.Form.Get("alias")
	}
	if r.Form.Get("tag") != "" {
		searchOpt["tags.value"] = r.Form.Get("tag")
	}

	artists := Artists{}
	if len(searchOpt) <= 0 {
		return
	}
	if err = c.Find(searchOpt).Sort("-rating.count").All(&artists); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	t, err := template.ParseFiles("result.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find a specified file: %s", err)
		os.Exit(1)
	}
	t.Execute(w, artists)
}
