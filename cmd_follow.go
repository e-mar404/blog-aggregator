package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/e-mar404/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("command should only have one argument: gator follow url")
	}

	feedUrl := sql.NullString { 
		String: cmd.arguments[0],
		Valid: true,
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams {
		ID: uuid.New(),
		UserID: uuid.NullUUID {
			UUID: user.ID,
			Valid: true,
		},
		FeedID: uuid.NullUUID {
			UUID: feed.ID,
			Valid: true,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("Following feed: %s\nFor cur user: %s\n", feed.Name.String, s.config.CurrentUserName)
	fmt.Printf("%v\n", feed)

	return nil
}
