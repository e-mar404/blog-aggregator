package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	fmt.Printf("resetting users table...\n")
	if err := s.db.RestUsers(context.Background()); err != nil {
		return err
	}
	fmt.Printf("success!\n")
	return nil
}
