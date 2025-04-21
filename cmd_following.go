package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("there should no be any arguments")
	}
	
	userName := sql.NullString {
		String: s.config.CurrentUserName,
		Valid: true,
	}
	user, err := s.db.GetUser(context.Background(), userName) 
	if err != nil {
		return err
	}
	
	userID := uuid.NullUUID {
		UUID: user.ID,
		Valid: true,
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), userID)

	fmt.Printf("Feeds following for %s:\n\n", s.config.CurrentUserName)
	for _, feedFollow := range feedFollows {
		fmt.Printf(" * %s\n", feedFollow.Name.String)
	}

	return nil
}
