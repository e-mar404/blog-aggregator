package main

import (
	"context"
	"fmt"
)

func handerlAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", feed)
	return nil
}
