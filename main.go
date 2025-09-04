package main

import (
	"database/sql"
	"github.com/AymaneIsmail/rss-gator/internal/config"
	"github.com/AymaneIsmail/rss-gator/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal("Error while connecting to the database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	commands := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	commands.register("login", loginHandler)
	commands.register("register", registerHandler)
	commands.register("reset", resetHandler)
	commands.register("users", listUsersHandler)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = commands.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
