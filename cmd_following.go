package main

import (
	"context"
	"fmt"

	"github.com/e-mar404/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 0 {
		return fmt.Errorf("there should no be any arguments")
	}
	
	userID := uuid.NullUUID {
		UUID: user.ID,
		Valid: true,
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), userID)
	if err != nil {
		return err
	}

	fmt.Printf("Feeds following for %s:\n\n", s.config.CurrentUserName)
	for _, feedFollow := range feedFollows {
		fmt.Printf(" * %s\n", feedFollow.Name.String)
	}

	return nil
}
