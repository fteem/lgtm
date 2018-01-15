package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("lgtm", "Looks GIFFed To Me")

	// Adding
	addGif     = app.Command("add", "Add a gif to collections")
	addGifURL  = addGif.Flag("url", "URL to gif").Short('u').Required().URL()
	addGifTags = addGif.Flag("tag", "Tag").Short('t').Required().Strings()

	// Take
	takeGif    = app.Command("take", "Take one from collection")
	takeGifTag = takeGif.Flag("tag", "Tag").Short('t').Required().String()

	// List
	lsGif = app.Command("ls", "List all GIFs in collection")

	// Delete
	deleteGif   = app.Command("del", "Delete GIF from collection")
	deleteGifID = deleteGif.Flag("id", "ID of GIF").Short('i').Required().Int()
)

const (
	storagePath = "gifs.yml"
)

func main() {
	repo := NewRepo(storagePath)
	err := repo.Init()
	if err != nil {
		panic(err)
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case addGif.FullCommand():
		url := *addGifURL
		gif := GIF{
			URL:  url.String(),
			Tags: *addGifTags,
		}
		err := repo.Save(gif)
		if err != nil {
			panic(err)
		}
		fmt.Println("LGTM")
	case takeGif.FullCommand():
		gifs, err := repo.ByTag(*takeGifTag)
		if err != nil {
			panic(err)
		}
		randomGif := RandomItem(gifs)
		ToClipboard(randomGif.URL)
		fmt.Println("Link copied to clipboard!")
	case lsGif.FullCommand():
		gifs, err := repo.All()
		if err != nil {
			panic(err)
		}
		fmt.Println(gifs)
	case deleteGif.FullCommand():
		err := repo.Delete(*deleteGifID)
		if err != nil {
			panic(err)
		}
		fmt.Println("LGTM")
	}
}
