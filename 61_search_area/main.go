package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	var artistName string
	flag.StringVar(&artistName, "name", "", "specify the artist name")
	flag.StringVar(&artistName, "n", "", "specify the artist name")
	flag.Parse()

	client, err := newRedisClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	area, err := client.Get(artistName).Result()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err == redis.Nil {
		return
	}
	fmt.Printf("(name, area) = (%s, %s)\n", artistName, area)
}

// newRedisClient returns client for connecting to Redis.
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
