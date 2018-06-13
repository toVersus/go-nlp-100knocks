package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

// TODO: improve O(n) implementation...
// https://stackoverflow.com/questions/17193176/searching-in-values-of-a-redis-db
func main() {
	var area string
	flag.StringVar(&area, "area", "", "specify the area")
	flag.Parse()

	client, err := newRedisClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	count, err := countArtistsInArea(client, area)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to count artists in specified area:\n  %s\n", err)
	}
	fmt.Printf("found %d keys\n", count)
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

func countArtistsInArea(client *redis.Client, area string) (int, error) {
	keys, err := client.Keys("*").Result()
	if err != nil {
		return 0, err
	}

	count := 0
	for _, key := range keys {
		val, err := client.Get(key).Result()
		if err != nil {
			return 0, err
		}
		if val != area {
			continue
		}
		count++
	}
	return count, nil
}
