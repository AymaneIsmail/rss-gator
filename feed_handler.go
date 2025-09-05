package main

import (
    "fmt"
    "context"
    "time"

    "github.com/google/uuid"
    "github.com/AymaneIsmail/rss-gator/internal/database"

)

func feedHandler(s *state, cmd command) error {
    feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
    if err != nil {
        return fmt.Errorf("❌ Error while fetching feed: %v", err)
    }

    fmt.Printf("\n✅ Feed loaded: %s\n", feed.Channel.Title)
    fmt.Printf("Description: %s\n", feed.Channel.Description)
    fmt.Printf("Link: %s\n", feed.Channel.Link)
    fmt.Printf("Items (%d):\n", len(feed.Channel.Item))

    for i, item := range feed.Channel.Item {
        fmt.Printf("\n%d. %s\n", i+1, item.Title)
        fmt.Printf("   Date: %s\n", item.PubDate)
        fmt.Printf("   Link: %s\n", item.Link)
        fmt.Printf("   Description: %s\n", item.Description)
    }

    return nil
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
        return fmt.Errorf("Cannot create feed")
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