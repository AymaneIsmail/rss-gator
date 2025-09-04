package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/AymaneIsmail/rss-gator/internal/database"
	"github.com/google/uuid"
)

func loginHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	user, err := s.db.GetOneUserByName(context.Background(), name)
	if err != nil {
		fmt.Printf("Cannot login if username '%s' does not exists in database", name)
		os.Exit(1)
	}

	err = s.cfg.SetUserName(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully to: ", name)
	return nil
}

func registerHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.Name)
	}

	userName := cmd.Args[0]

	user, err := s.db.GetOneUserByName(context.Background(), userName)
	if err == nil {
		fmt.Printf("❌ Registration failed: User with name '%s' already exists.\n", userName)
		os.Exit(1)
	}
	if err != sql.ErrNoRows {
		return fmt.Errorf("Database query error while checking user existence: %v", err)
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}

	user, err = s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Database error while creating user: %v", err)
	}

	if err := s.cfg.SetUserName(user.Name); err != nil {
		return fmt.Errorf("Could not save user name to config: %v", err)
	}

	fmt.Printf(
		"✅ Registration successful!\nUser '%s' was created.\nFull user record:\n%+v\n",
		user.Name, user,
	)

	return nil
}
