package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/e-mar404/gator/internal/database"
	"github.com/google/uuid"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("there should be a rss url to unfollow")
	}

	url := sql.NullString {
		String: cmd.arguments[0],
		Valid: true,
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowsParams {
		UserID: uuid.NullUUID{
			UUID: user.ID,
			Valid: true,
		},
		FeedID: uuid.NullUUID{
			UUID: feed.ID,
			Valid: true,
		},
	}
	if err = s.db.DeleteFeedFollows(context.Background(), params); err != nil {
		return err 
	}

	fmt.Printf("successfully unfollowed: %s\n", cmd.arguments[0])

	return nil
}
