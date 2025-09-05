package main

import (
	"fmt"
	"context"
	"time"

	"github.com/google/uuid"
    "github.com/AymaneIsmail/rss-gator/internal/database"

)

func followHandler(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage <follow> <url>")
	}

	ctx := context.Background()
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("Error getting feed by url %s", url)
	}

	params := database.CreateFeedFollowParams{
		ID:      uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID: feed.ID,
		UserID: user.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(ctx, params)
	if err != nil {
		return fmt.Errorf("Cannot create feed follow: %v", err)
	}

	fmt.Printf("%v\n", feedFollow)

	return nil
}

func listFollowFeeds(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Cannot get followfeeds %v", err)
	}


	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
	}

	return nil
}

func unfollowFeedHandler(s *state, cmd command, user database.User) error {

	// if len(cmd.Args) != 1 {
	// 	return fmt.Errorf("need url params")
	// }

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error getting feed by url %s", url)
	}


	unfollowParams := database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	_, err = s.db.UnfollowFeed(context.Background(), unfollowParams)
	if err != nil {
		return fmt.Errorf("Cannot drop feed %s", feed.Name)
	}

	fmt.Printf("%s\n", feed.Name)

	return nil
}