package main

import (
	"context"
	"fmt"
)

func handlerFeeds (s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("there shouln't be any args for this command")
	}

	// list feeds
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Printf("There are no feeds registered, you can add a feed with `gator addfeed name url`\n")
	}

	for _, feed := range feeds {
		fmt.Printf("RSS Feed: (%s)\n", feed.FeedName.String)
		fmt.Printf("* URL: %s\n", feed.FeedUrl.String)
		fmt.Printf("* Created By: %s\n", feed.UserName.String)
		fmt.Printf("====================================\n")
	}

	return nil
}
