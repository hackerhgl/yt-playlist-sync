package main

import (
	"encoding/json"
	"os"
	"strings"

	"google.golang.org/api/youtube/v3"
)

func SavePlaylistToJSON(items []*youtube.PlaylistItem, total int) error {
	for index, item := range items {
		items[index].Snippet.Title = strings.ReplaceAll(item.Snippet.Title, "/", "|") + ".mp3"
	}
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	file, err := os.Create("db/playlist.json")
	if err != nil {
		return err
	}
	file.Write(b)
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func InitDirs() error {
	paths := []string{"songs", "db", "logs"}

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0775)
		}
	}

	return nil
}
