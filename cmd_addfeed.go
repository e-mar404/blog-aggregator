package main

import (
	"context"
	"database/sql"
	"e-mar404/gator/internal/database"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("expect exactly 2 arguments, name and url")
	}

	name := sql.NullString {
		String: s.config.CurrentUserName,
		Valid: true,
	}
	curUser, err := s.db.GetUser(context.Background(), name) 
	if err != nil {
		return fmt.Errorf("error getting cur user: %v", err)
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
			UUID: curUser.ID,
			Valid: true,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	feed, err := s.db.CreateFeed(context.Background(), params)	

	feedFollowParams := database.CreateFeedFollowParams {
		ID: uuid.New(),
		UserID: uuid.NullUUID {
			UUID: curUser.ID,
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
