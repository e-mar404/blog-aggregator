package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/e-mar404/gator/internal/database"
	"github.com/google/uuid"
)

func handerlAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("please provide a time duration to wait during requests")
	}
	
	time_between_reqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}
	
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		if err = scrapeFeeds(s); err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, fetchErr := s.db.GetNextFeedToFetch(context.Background())
	if fetchErr != nil {
		if strings.Contains(fetchErr.Error(), "no rows in result set") {
			return nil
		}
		return fetchErr 
	}

	feedContent, err := fetchFeed(context.Background(), nextFeed.Url.String) 
	if err != nil {
		return err
	}

	fmt.Println("Feed fetched")

	params := database.MarkFeedFetchedParams {
		ID: nextFeed.ID,
		LastFetchedAt: sql.NullTime {
			Time: time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
	}
	if err = s.db.MarkFeedFetched(context.Background(), params); err != nil {
		fmt.Println(err)
		return err 
	}
	
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"
	for _, item := range feedContent.Channel.Item {
		pubDate, err := time.Parse(layout, item.PubDate)
		if err != nil {
			return err 
		}

		postParams := database.CreatePostParams {
			ID: uuid.New(),
			Title: sql.NullString {
				String: item.Title,
				Valid: true,
			},
			Url: sql.NullString {
				String: item.Link,
				Valid: true,
			},
			Description: sql.NullString {
				String: item.Description,
				Valid: true,
			},
			PublishedAt: sql.NullTime {
				Time: pubDate,
				Valid: true,
			},
			FeedID: uuid.NullUUID {
				UUID: nextFeed.ID,
				Valid: true,
			},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	
		_, err = s.db.CreatePost(context.Background(), postParams)
		if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint \"posts_url_key\"") {
			return err
		}
	}

	fmt.Println("Posts saved in database")

	return nil
}
