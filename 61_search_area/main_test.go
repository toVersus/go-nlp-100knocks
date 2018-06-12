package main

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/go-test/deep"
)

var searchAreaTests = []struct {
	name string
	key  string
	want string
}{
	{
		name: "should get area of artist by using name key",
		key:  "The Silhouettes",
		want: "United States",
	},
	{
		name: "should not get area of artist due to empty string value",
		key:  "WIK▲N",
		want: "",
	},
}

func TestSearchArea(t *testing.T) {
	for _, testcase := range searchAreaTests {
		t.Log(testcase.name)

		client, err := newRedisClient()
		if err != nil {
			t.Error(err)
		}

		result, err := client.Get(testcase.key).Result()
		if err != nil {
			t.Error(err)
		}
		if err == redis.Nil {
			if diff := deep.Equal(result, testcase.want); diff != nil {
				t.Error(diff)
			}
		}
	}
}
