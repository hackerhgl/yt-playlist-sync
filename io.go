package main

import (
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/api/youtube/v3"
)

func SavePlaylistToJSON(items []*youtube.PlaylistItem)  {
	b, err :=  json.Marshal(items)
	if err != nil {
        fmt.Println(err)
        return
    }
	file, err := os.Create("db/playlist.json")
	file.Write(b)
	if err != nil {
        fmt.Println(err)
        return
    }
}