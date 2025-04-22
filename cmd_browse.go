package main

import (
	"context"
	"e-mar404/gator/internal/database"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) > 1 {
		return fmt.Errorf("expecting either no argument for a default limit of 2 or 1 argument for the desired limit")
	}

	postLimit := int32(2)
	if len(cmd.arguments) == 1 {
		postLimit64, err := strconv.ParseInt(cmd.arguments[0], 10, 32)
		postLimit = int32(postLimit64)
		if err != nil {
			return err
		}
	}

	params := database.GetPostsForUserParams {
		UserID: uuid.NullUUID{
			UUID: user.ID,
			Valid: true,
		},
		Limit: postLimit,
	}
	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}
	
	if len(posts) == 0 {
		fmt.Printf("No posts found\n")
		return nil
	}

	for _, post := range posts {
		fmt.Printf("Feed: %s\n", post.FeedName.String)
		fmt.Printf("Title: %s\n", post.Title.String)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url.String)
		fmt.Printf("=============================================\n\n")
	}

	return nil
}
