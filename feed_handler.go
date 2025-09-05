package main

import (
    "fmt"
    "context"
    "time"
	"strings"
    "log"

    "github.com/google/uuid"
    "github.com/AymaneIsmail/rss-gator/internal/database"
    "database/sql"

)

func feedHandler(s *state, cmd command) error {

    if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

    timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

    ticker := time.NewTicker(timeBetweenRequests)


    for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}


func addFeedHandler(s *state, cmd command, user database.User) error {
    if len(cmd.Args) < 2 {
        return fmt.Errorf("usage: addfeed <name> <url>")
    }

    ctx := context.Background()

    feedName := cmd.Args[0]
    url := cmd.Args[1]

    params := database.CreateFeedParams{
        ID: uuid.New(),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        Name:     feedName,
	    Url:      url,
	    UserID: user.ID,
    }

    feed, err := s.db.CreateFeed(ctx, params)
    if err != nil {
        return fmt.Errorf("Cannot create feed %v", err)
    }

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:      uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID: feed.ID,
		UserID: user.ID,
	}

	_, err = s.db.CreateFeedFollow(ctx, createFeedFollowParams)
	if err != nil {
		return fmt.Errorf("Cannot create feed follow: %v", err)
	}

    fmt.Printf("Feed created %v\n", feed)

    return nil
}

func listFeedsHandler(s *state, cmd command) error {
    feeds, err := s.db.GetUserFeeds(context.Background())
    if err != nil {
		return fmt.Errorf("❌ Error retrieving feeds: %v", err)
    }

    if len(feeds) == 0 {
        fmt.Println("ℹ️ No feeds found in the database.")
		return nil
    }

    for _, feed := range feeds {
        fmt.Printf("%s\n", feed.Name)
        fmt.Printf("%s\n", feed.Url)
         fmt.Printf("%s\n", feed.UserName)
    }

    return nil

}

func scrapeFeeds(s *state) {
    feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			FeedID:    feed.ID,
			Title:     item.Title,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			Url:         item.Link,
			PublishedAt: publishedAt,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}