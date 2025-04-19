package main

import (
	"context"
	"database/sql"
	"e-mar404/gator/internal/database"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("provide exactly one user to register")
	}

	params := database.CreateUserParams {
		ID: uuid.New(),
		Name: sql.NullString{
			String: cmd.arguments[0],
			Valid: true,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err := s.db.CreateUser(context.Background(), params) 
	if err != nil {
		return fmt.Errorf("error while creating user: %v", err)
	}
	
	fmt.Printf("user created: %v\n", user)
	s.config.SetUser(user.Name.String)
	fmt.Printf("user set successfully\n")

	return nil
}
