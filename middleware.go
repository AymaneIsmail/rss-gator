package main

import (
	"context"
	"fmt"

	"github.com/AymaneIsmail/rss-gator/internal/database"
)

// Wrap handlers that require a logged-in user.
// Usage: cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
    return func(s *state, cmd command) error {
        user, err := s.db.GetOneUserByName(context.Background(), s.cfg.CurrentUserName)
        if err != nil {
            return fmt.Errorf("you must be logged in: %w", err)
        }
        return handler(s, cmd, user)
    }
}
