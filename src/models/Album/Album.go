package Album

import (
	"fmt"
    "github.com/agus24/learngolang/src/DB"
)

var db = DB.Init()

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

func AlbumsByArtist(name string) ([]Album, error) {
    var albums []Album

	if db == nil {
		fmt.Println("db is nil")
	}
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()

    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}
