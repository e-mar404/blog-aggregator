package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %v", err)
	}

	for _, user := range users {
		if s.config.CurrentUserName == user.Name.String {
			fmt.Printf("* %s (current)\n", user.Name.String)
		} else {
			fmt.Printf("* %s\n", user.Name.String)
		}
	}

	return nil
}
