package main

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var insertDocsTests = []struct {
	name string
	text string
	want []int
}{
	{
		name: "should get the IDs of artists",
		text: `{"name": "Sam James", "area": "United States", "gender": "Male", "sort_name": "James, Sam", "ended": true, "gid": "183da4be-0cb0-4e6d-ba6d-91e57b7a6780", "type": "Person", "id": 729749, "aliases": [{"name": "Sam James Vende", "sort_name": "Sam James Vende"}]}
{"name": "Norman Kolodziej", "area": "Germany", "gender": "Male", "sort_name": "Kolodziej, Norman", "ended": true, "gid": "5ff386f1-2c4e-4c1c-b5fc-668ec25e1b3e", "type": "Person", "id": 811484}
{"name": "Bass Cube", "sort_name": "Bass Cube", "ended": true, "gid": "f1568f36-152b-40da-aef3-3582636f88be", "type": "Group", "id": 6153}
{"name": "Medras", "sort_name": "Medras", "ended": true, "gid": "a7d007ec-8026-4e84-982d-b6306baa14df", "type": "Person", "id": 723542}
{"name": "Kalev Lindal", "area": "Estonia", "gender": "Male", "sort_name": "Lindal, Kalev", "ended": true, "gid": "8864f9e3-6a03-40a7-9acb-ac8386d404e7", "type": "Person", "id": 892318}
{"name": "Nick Flower", "sort_name": "Flower, Nick", "ended": true, "gid": "537b606c-d7e4-4c8a-8e39-91e22c2bf720", "type": "Person", "id": 725047}
{"name": "Danièle Forget", "sort_name": "Forget, Danièle", "ended": true, "gid": "5d6d5857-3d03-4b1b-a3ee-d789e883c1b2", "type": "Person", "id": 726135}`,
		want: []int{729749, 811484, 6153, 723542, 892318, 725047, 726135},
	},
}

func TestGetTags(t *testing.T) {
	for _, testcase := range insertDocsTests {
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

		t.Log("")
		t.Log("before indexing")
		for _, artist := range artists {
			query := c.Find(bson.M{"name": artist.Name})

			a, count, err := getQueryTime(query)
			if err != nil {
				t.Error(err)
			}
			t.Logf("%s found...%d\n", a.Name, count)
		}
		t.Log("")

		t.Log("after indexing")
		results := []int{}
		for _, artist := range artists {
			keys := []string{"name", "aliases.name", "tags.value", "rating.value"}
			for _, key := range keys {
				err = c.EnsureIndexKey(key)
				if err != nil {
					t.Error(err)
				}
			}

			query := c.Find(bson.M{"name": artist.Name})
			a, count, err := getQueryTime(query)
			if err != nil {
				t.Error(err)
			}
			t.Logf("%s found...%d\n", a.Name, count)
			results = append(results, a.ID)
		}

		if diff := deep.Equal(results, testcase.want); diff != nil {
			t.Error(diff)
		}

		if err = db.DropDatabase(); err != nil {
			t.Errorf("could not delete database\n  %s\n", err)
		}
	}
}
