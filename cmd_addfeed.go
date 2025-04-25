package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/e-mar404/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("expect exactly 2 arguments, name and url")
	}

	params := database.CreateFeedParams {
		ID: uuid.New(),
		Name: sql.NullString{
			String: cmd.arguments[0],
			Valid: true,
		},
		Url: sql.NullString{
			String: cmd.arguments[1],
			Valid: true,
		},
		UserID: uuid.NullUUID{
			UUID: user.ID,
			Valid: true,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	feed, err := s.db.CreateFeed(context.Background(), params)	

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
