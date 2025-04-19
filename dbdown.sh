#! /usr/bin/env bash

pushd sql/schema
goose postgres "postgres://emar:@localhost:5432/gator?sslmode=disable" down
popd
