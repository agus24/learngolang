package main

import (
	"log"
	"fmt"
	models "github.com/agus24/learngolang/src/models/Album"
)

func main() {
	albums, err := models.AlbumsByArtist("Sarah Vaughan")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}
