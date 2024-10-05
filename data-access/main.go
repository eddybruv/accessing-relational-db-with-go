package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float64
}

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Println("Error loading .env file")
	}

	config := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("DB Connected 🙌")
	drakeAlbums, err := albumsByArtist("Kanye")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Albums found:", drakeAlbums)

	alb, err := albumByID(49)
	if err != nil {
		log.Fatal("err")
	}
	fmt.Println("Albums found:", alb)

	albID, err := addAlbum(Album{
		Title:  "Split Decision",
		Artist: "Dave",
		Price:  72.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
}

// albumsByArtist queries for albums that have the specified artist name
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("select * from album where artist like ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var tempAlbum Album
		if err := rows.Scan(&tempAlbum.ID, &tempAlbum.Title, &tempAlbum.Artist, &tempAlbum.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, tempAlbum)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

// albumByID queries for the album with the specified ID
func albumByID(id int64) (Album, error) {
	var tempAlbum Album

	row := db.QueryRow("select * from album where id=?", id)
	if err := row.Scan(&tempAlbum.ID, &tempAlbum.Title, &tempAlbum.Artist, &tempAlbum.Price); err != nil {
		// use errors.Is instead of err == sql.ErrNoRows
		if errors.Is(err, sql.ErrNoRows) {
			return tempAlbum, fmt.Errorf("albumsById %d: no such album", id)
		}
		return tempAlbum, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return tempAlbum, nil
}

// addAlbum
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT into album (title, artist, price) values (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
