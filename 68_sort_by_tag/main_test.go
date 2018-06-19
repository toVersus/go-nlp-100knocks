package main

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var sortByTagTests = []struct {
	name string
	text string
	tag  string
	want []string
}{
	{
		name: "should get the ranking of artists retrieved by tag name",
		text: `{"name": "Sam James", "area": "United States", "gender": "Male", "sort_name": "James, Sam", "ended": true, "gid": "183da4be-0cb0-4e6d-ba6d-91e57b7a6780", "type": "Person", "id": 729749, "aliases": [{"name": "Sam James Vende", "sort_name": "Sam James Vende"}]}
{"name": "Norman Kolodziej", "area": "Germany", "gender": "Male", "sort_name": "Kolodziej, Norman", "ended": true, "gid": "5ff386f1-2c4e-4c1c-b5fc-668ec25e1b3e", "type": "Person", "id": 811484}
{"name": "Bass Cube", "sort_name": "Bass Cube", "ended": true, "gid": "f1568f36-152b-40da-aef3-3582636f88be", "type": "Group", "id": 6153}
{"name": "Medras", "sort_name": "Medras", "ended": true, "gid": "a7d007ec-8026-4e84-982d-b6306baa14df", "type": "Person", "id": 723542}
{"name": "Kalev Lindal", "area": "Estonia", "gender": "Male", "sort_name": "Lindal, Kalev", "ended": true, "gid": "8864f9e3-6a03-40a7-9acb-ac8386d404e7", "type": "Person", "id": 892318}
{"name": "Nick Flower", "sort_name": "Flower, Nick", "ended": true, "gid": "537b606c-d7e4-4c8a-8e39-91e22c2bf720", "type": "Person", "id": 725047}
{"name": "Danièle Forget", "sort_name": "Forget, Danièle", "ended": true, "gid": "5d6d5857-3d03-4b1b-a3ee-d789e883c1b2", "type": "Person", "id": 726135}
{"name": "Anika Paris", "tags": [{"count": 1, "value": "production music"}], "sort_name": "Paris, Anika", "ended": true, "gid": "752a5dba-e08c-4d44-8f2b-af8144fe4fd6", "id": 57821}
{"rating": {"count": 15, "value": 84}, "begin": {"month": 3, "year": 1985}, "name": "Guns N’ Roses", "area": "United States", "tags": [{"count": 1, "value": "glam metal"}, {"count": 1, "value": "pop/rock"}, {"count": 1, "value": "classic pop and rock"}, {"count": 1, "value": "heavy metal"}, {"count": 1, "value": "metal"}, {"count": 2, "value": "rock"}, {"count": 1, "value": "american"}, {"count": 5, "value": "hard rock"}, {"count": 1, "value": "patience"}, {"count": 1, "value": "band"}], "sort_name": "Guns N’ Roses", "ended": true, "gid": "eeb1195b-f213-4ce1-b28c-8565211f8e43", "type": "Group", "id": 1440, "aliases": [{"name": "ガンズ・アンド・ローゼズ", "sort_name": "ガンズ・アンド・ローゼズ"}, {"name": "Guns and Roses", "sort_name": "Guns and Roses"}, {"name": "Guns -N- Roses", "sort_name": "Guns -N- Roses"}, {"name": "guns 'n' roses", "sort_name": "guns 'n' roses"}, {"name": "Guns'N Roses", "sort_name": "Guns'N Roses"}, {"name": "Guns'N'Roses", "sort_name": "Guns'N'Roses"}, {"name": "GNR", "sort_name": "GNR"}, {"name": "Guns´n Roses", "sort_name": "Guns´n Roses"}, {"name": "Gn'R", "sort_name": "Gn'R"}, {"name": "Guns 'N Roses", "sort_name": "Guns 'N Roses"}, {"name": "Guns N'Roses", "sort_name": "Guns N'Roses"}, {"name": "Guns N´ Roses", "sort_name": "Guns N´ Roses"}, {"name": "Guns N Roses", "sort_name": "Guns N Roses"}, {"name": "Guns & Roses", "sort_name": "Guns & Roses"}, {"name": "Guns 'N' Roses", "sort_name": "Guns 'N' Roses"}]}`,
		tag:  "glam metal",
		want: []string{"Guns N’ Roses"},
	},
}

func TestGetTags(t *testing.T) {
	for _, testcase := range sortByTagTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		artists, err := readBSON(r)
		if err != nil {
			t.Errorf("could not parse a JSON file: %s", err)
		}

		session, err := mgo.Dial("mongodb://localhost")
		if err != nil {
			t.Error(err)
		}

		db := session.DB("Testing")
		c := db.C("artist")

		for _, artist := range artists {
			err := c.Insert(artist)
			if err != nil {
				t.Error(err)
			}
		}

		keys := []string{"name", "aliases.name", "tags.value", "rating.value"}
		for _, key := range keys {
			err = c.EnsureIndexKey(key)
			if err != nil {
				t.Error(err)
			}
		}

		var results Artists
		if err := c.Find(bson.M{"tags.value": testcase.tag}).Sort("-rating.count").All(&results); err != nil {
			t.Error(err)
		}

		result := []string{}
		for _, artist := range results {
			result = append(result, artist.Name)
		}
		if diff := deep.Equal(result, testcase.want); diff != nil {
			t.Error(diff)
		}

		if err = db.DropDatabase(); err != nil {
			t.Errorf("could not delete database\n  %s\n", err)
		}
	}
}
