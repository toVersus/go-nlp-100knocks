package main

import (
	"testing"

	"github.com/go-test/deep"
)

var countArtistsTests = []struct {
	name string
	area string
	want int
}{
	{
		name: "should count the number of artists lived in specified area",
		area: "Minneapolis–Saint Paul",
		// actual count is 9 but that artist name is duplicated and overwritten by the one in another area.
		want: 8,
	},
	{
		name: "should return zero",
		area: "null",
		want: 0,
	},
}

func TestGetKeyVal(t *testing.T) {
	for _, testcase := range countArtistsTests {
		t.Log(testcase.name)

		client, err := newRedisClient()
		if err != nil {
			t.Error(err)
		}

		result, err := countArtistsInArea(client, testcase.area)
		if diff := deep.Equal(result, testcase.want); diff != nil {
			t.Error(diff)
		}
	}
}
