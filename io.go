package main

import (
	"encoding/json"
	"os"
	"time"

	"google.golang.org/api/youtube/v3"
)

func TimeStamp() error {

	file, err := os.Create("db/timestamp.txt")

	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(time.Now().String()))

	return nil
}

func SavePlaylistToJSON(items []*youtube.PlaylistItem, total int) error {
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
