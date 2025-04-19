package main

import (
	"e-mar404/gator/internal/config"
	"e-mar404/gator/internal/database"
)

type state struct {
	config *config.Config
	db *database.Queries
}
