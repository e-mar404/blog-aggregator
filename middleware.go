package main

import (
	"context"
	"database/sql"
	"e-mar404/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		userName := sql.NullString {
			String: s.config.CurrentUserName,
			Valid: true,
		}
		user, err := s.db.GetUser(context.Background(), userName)
		if err != nil {
			return err
		}
		handler(s, cmd, user)
		return nil
	}
}
