# Accessing a Relational DB with Go

This project demonstrates how to read from and write to a MySQL database using the Go programming language and the `go-sql-driver/mysql` package. It covers basic database interactions such as querying, inserting, and error handling.

## Project Structure

The main files of this project are located in the `data-access` directory, and the project makes use of environment variables to configure the database connection.

### Files:

- `main.go`: Contains the Go code to interact with the MySQL database.
- `create-table.sql`: SQL script to create the necessary table (`album`) in the database.
- `more-albums.sql`: SQL script to populate the `album` table with sample data.

## Database Information

- **Database Name**: `recording`
- **Table Created**: `album`
    - `id`: Auto-incremented integer serving as the primary key.
    - `title`: The title of the album.
    - `artist`: The artist of the album.
    - `price`: The price of the album.

## Setup

### Prerequisites

- Go installed on your machine.
- MySQL server running and accessible.
- The `go-sql-driver/mysql` package installed.
- The `github.com/joho/godotenv` package installed to load environment variables.

### Installation

1. Clone the repository to your local machine.
2. Install dependencies:
   ```bash
   go mod tidy
3. Create a .env file in the root of your project directory, and set your database credentials:
   ```bash
   go mod tidy
4. Run the SQL script `create-table.sql` to set up the `album` table in the MySQL database
   ```bash
   mysql -u your_username -p recording < create-table.sql
5. Populate the table using `more-albums.sql`
   ```bash
   mysql -u your_username -p recording < more-albums.sql

## Running the Project

Once everything is set up, you can run the Go program using:

```bash
go run data-access/main.go
```

### This will:

1. Connect to the recording database.
2. Query albums by a specific artist.
3. Retrieve an album by its ID.
4. Add a new album to the database.

Reference https://go.dev/doc/tutorial/database-access