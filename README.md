# Juel Ratings API

A REST API for serving college football team ratings and statistics data.

## Overview

This Go-based API provides endpoints for accessing college football team data, including team information, conference affiliations, and geographic data. It uses SQLite for local database access.

## Tech Stack

- **Language**: Go 1.21+
- **Database**: SQLite
- **Database Driver**: mattn/go-sqlite3
- **Architecture**: Clean architecture with separate store and handler packages

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go           # Application entry point
├── internal/
│   ├── handlers/
│   │   └── teams.go          # HTTP handlers for teams
│   ├── models/
│   │   └── models.go         # Data models
│   └── store/
│       └── store.go          # Database operations
├── .env                      # Environment variables (not in git)
├── go.mod                    # Go module definition
└── go.sum                    # Go module checksums
```

## Setup

### Prerequisites

- Go 1.21 or later
- A SQLite database file containing the `teams` table
- `.env` file with the following variable:
  ```
  DATABASE_URL=/absolute/path/to/juel_ratings.db
  ```

### Installation

1. Clone the repository:
   ```bash
   git clone <your-repo-url>
   cd juel-ratings-api
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up your `.env` file with the SQLite database path.

4. Run the application:
   ```bash
   go run ./cmd/api
   ```

## API Endpoints

### Teams

- **GET /teams** - Returns all teams
  ```json
  [
    {
      "id": 1,
      "cfbd_id": 5,
      "school": "Alabama",
      "mascot": "Crimson Tide",
      "abbreviation": "ALA",
      "conference": "SEC",
      "division": "West",
      "classification": "FBS",
      "city": "Tuscaloosa",
      "state": "AL"
    }
  ]
  ```

- **GET /teams/{id}** - Returns a specific team by ID (Coming soon)

## Configuration

- `DATABASE_URL` should point to a SQLite database file using an absolute filesystem path.
- The application looks for `.env` in the current working directory and then walks up parent directories until it finds one.
- The SQLite database file must already exist before the API starts.

## Database Schema

The application expects a SQLite database with a `teams` table containing:

- `id` (INTEGER PRIMARY KEY)
- `cfbd_id` (INT NOT NULL UNIQUE)
- `school` (TEXT NOT NULL)
- `mascot` (TEXT)
- `abbreviation` (TEXT)
- `conference` (TEXT)
- `division` (TEXT)
- `classification` (TEXT)
- `city` (TEXT)
- `state` (TEXT)

The `GET /teams` endpoint selects these columns directly:

- `id`
- `cfbd_id`
- `school`
- `mascot`
- `abbreviation`
- `conference`
- `division`
- `classification`
- `city`
- `state`

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o api ./cmd/api
```

### Running the API

```bash
go run ./cmd/api
```

### Architecture

The project follows a clean architecture pattern:

- **Store Package**: Handles all database operations and abstracts SQL queries
- **Handlers Package**: HTTP handlers that use the store to fetch data
- **Models Package**: Shared data structures used across the application
- **Main**: Wire up dependencies (DB → Store → Handlers → HTTP Server)

## License

MIT

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
