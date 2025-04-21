package main

import (
	"context"
	"database/sql"
	"e-mar404/gator/internal/database"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("command should only have one argument: gator follow url")
	}

	userName := sql.NullString { 
		String: s.config.CurrentUserName,
		Valid: true,
	}
	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}

	feedUrl := sql.NullString { 
		String: cmd.arguments[0],
		Valid: true,
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams {
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

	_, err = s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("Following feed: %s\nFor cur user: %s\n", feed.Name.String, s.config.CurrentUserName)

	return nil
}
