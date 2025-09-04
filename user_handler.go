package main

import (
	"context"
	"fmt"
)

func resetHandler(s *state, cmd command) error {
	if err := s.db.DropUsers(context.Background()); err != nil {
		return fmt.Errorf("❌ Failed to drop 'users' table: %v", err)
	}

	fmt.Println("✅ 'users' table dropped successfully.")

	return nil
}

// List all users with improved logging
func listUsersHandler(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("❌ Error retrieving users: %v", err)
	}

	if len(users) == 0 {
		fmt.Println("ℹ️ No users found in the database.")
		return nil
	}

	fmt.Printf("✅ Found %d user(s):\n", len(users))

	for _, user := range users {

		if s.cfg.CurrentUserName == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {

			fmt.Printf("* %s\n", user.Name)
		}

	}

	return nil
}
