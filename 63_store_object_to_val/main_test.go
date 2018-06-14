package main

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
)

var getTagsTests = []struct {
	name       string
	text       string
	artistName string
	want       []Tag
}{
	{
		name: "should get the proper tags",
		text: `{"name": "WIK▲N", "tags": [{"count": 1, "value": "sillyname"}], "sort_name": "WIK▲N", "ended": true, "gid": "8972b1c1-6482-4750-b51f-596d2edea8b1", "id": 805192}
{"name": "Gustav Ruppke", "sort_name": "Gustav Ruppke", "ended": true, "gid": "b4f76788-7e6f-41b7-ac7b-dfb67f66282e", "type": "Person", "id": 578352}`,
		artistName: "WIK▲N",
		want: []Tag{
			Tag{Count: 1, Value: "sillyname"},
		},
	},
}

func TestGetTags(t *testing.T) {
	for _, testcase := range getTagsTests {
		t.Log(testcase.name)

		r := strings.NewReader(testcase.text)
		artists, err := readJSON(r)
		if err != nil {
			t.Errorf("could not parse a JSON file:\n  %s", err)
		}

		client, err := newRedisClient()
		if err != nil {
			t.Error(err)
		}

		err = artists.rpushTags(client)
		if err != nil {
			t.Error(err)
		}

		results, err := getTags(client, testcase.artistName)
		if err != nil {
			t.Error(err)
		}

		if diff := deep.Equal(results, testcase.want); diff != nil {
			t.Error(diff)
		}

		for _, artist := range artists {
			_, err := client.Del(artist.Name).Result()
			if err != nil {
				t.Error(err)
			}
		}
	}
}
