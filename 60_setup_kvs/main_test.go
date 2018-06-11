package main

import (
	"strings"
	"testing"

	"github.com/go-redis/redis"

	"github.com/go-test/deep"
)

var readJSONTests = []struct {
	name string
	text string
	want []*Artist
}{
	{
		name: "should parse JSON from full text",
		text: `{"name": "WIK▲N", "tags": [{"count": 1, "value": "sillyname"}], "sort_name": "WIK▲N", "ended": true, "gid": "8972b1c1-6482-4750-b51f-596d2edea8b1", "id": 805192}
{"name": "Gustav Ruppke", "sort_name": "Gustav Ruppke", "ended": true, "gid": "b4f76788-7e6f-41b7-ac7b-dfb67f66282e", "type": "Person", "id": 578352}`,
		want: []*Artist{
			&Artist{
				ID:   805192,
				GID:  "8972b1c1-6482-4750-b51f-596d2edea8b1",
				Name: "WIK▲N",
				Tags: []*Tag{
					&Tag{
						Count: 1,
						Value: "sillyname",
					},
				},
				SortName: "WIK▲N",
			},
			&Artist{
				ID:       578352,
				GID:      "b4f76788-7e6f-41b7-ac7b-dfb67f66282e",
				Name:     "Gustav Ruppke",
				SortName: "Gustav Ruppke",
			},
		},
	},
	{
		name: "should not parse JSON from empty text",
		text: ``,
		want: []*Artist{
			&Artist{},
		},
	},
}

func TestReadJSON(t *testing.T) {
	for _, testcase := range readJSONTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		results, err := readJSON(r)
		if err != nil {
			t.Errorf("could not parse a JSON file: %s", err)
		}

		if diff := deep.Equal(results, testcase.want); diff != nil {
			t.Error(diff)
		}
	}
}

var getKeyValTests = []struct {
	name string
	text string
	want []string
}{
	{
		name: "should parse JSON from full text",
		text: `{"begin": {"year": 1956}, "end": {"year": 1993}, "name": "The Silhouettes", "area": "United States", "sort_name": "Silhouettes, The", "ended": true, "gid": "ca3f3ee1-c4a7-4bac-a16a-0b888a396c6b", "type": "Group", "id": 101060, "aliases": [{"name": "Silhouettes", "sort_name": "Silhouettes"}, {"name": "The Sihouettes", "sort_name": "The Sihouettes"}]}
{"name": "Gustav Ruppke", "sort_name": "Gustav Ruppke", "ended": true, "gid": "b4f76788-7e6f-41b7-ac7b-dfb67f66282e", "type": "Person", "id": 578352}`,
		want: []string{
			"United States",
			"",
		},
	},
	{
		name: "should not parse JSON from empty text",
		text: ``,
		want: []string{""},
	},
}

func TestGetKeyVal(t *testing.T) {
	for _, testcase := range getKeyValTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		artists, err := readJSON(r)
		if err != nil {
			t.Errorf("could not parse a JSON file: %s", err)
		}

		client, err := newRedisClient()
		if err != nil {
			t.Error(err)
		}

		var results []string
		for _, artist := range artists {
			err = client.Set(artist.Name, artist.Area, 0).Err()
			if err != nil {
				t.Error(err)
			}

			val, err := client.Get(artist.Name).Result()
			if err != nil {
				t.Error(err)
			}
			if err != redis.Nil {
				results = append(results, val)
			}
		}

		if diff := deep.Equal(results, testcase.want); diff != nil {
			t.Error(diff)
		}
	}
}
