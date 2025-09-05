# Gator

A CLI tool for managing and following RSS feeds from the command line.

## Prerequisites

Before using Gator, you need to have installed:

- **Docker** : To run the PostgreSQL database
- **Go** (version 1.21 or higher) : Programming language to compile and install the CLI

### Installing Go

Download and install Go from [golang.org](https://golang.org/download/)

## Installation

### Install the Gator CLI

```bash
go install github.com/AymaneIsmail/rss-go run .
```


## Setup

### 1. Run the initialization script

The project includes an `init.sh` script that will automatically set up PostgreSQL with Docker and create the configuration file.

**Make the script executable and run it:**
```bash
chmod +x init.sh
./init.sh
```

This script will:
- Start a PostgreSQL v15 container with Docker
- Create the `.gatorconfig.json` configuration file in your home directory
- Set up the database connection automatically

### 2. Update your username in the config

After running the script, edit the configuration file to set your username:

```bash
nano ~/.gatorconfig.json
```

Replace `votre-nom-utilisateur` with your desired username:
```json
{
  "db_url": "postgres://postgres:password@localhost:5432/go run .?sslmode=disable",
  "current_user_name": "your-username"
}
```

## Usage

### Example Usage

```bash
# reset user
go run . reset 

# Register a user
go run . register alice

# Login
go run . login alice

# Add a feed
go run . addfeed techcrunch https://techcrunch.com/feed/

# Follow the feed
go run . follow techcrunch

# Start aggregation (fetches new posts every 60 seconds)
go run . agg 60

# In another terminal, browse the latest posts
go run . browse 10
```