package main

import (
	"github.com/e-mar404/gator/internal/config"
	"github.com/e-mar404/gator/internal/database"
)

type state struct {
	config *config.Config
	db *database.Queries
}
