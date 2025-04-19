package main

import (
	"context"
	"database/sql"
	"fmt"
)

func handelerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("provide exactly one username as an argument to login")
	}

	// check if user exits
	name := sql.NullString{
		String: cmd.arguments[0],
		Valid: true,
	}
	user, err := s.db.GetUser(context.Background(), name) 
	if err != nil {
		return fmt.Errorf("user does not exists")
	}

	if err := s.config.SetUser(user.Name.String); err != nil {
		return fmt.Errorf("error setting username: %v", err)
	}
	fmt.Printf("user has been set!\n")
	return nil
}
