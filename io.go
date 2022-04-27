package main

import (
	"encoding/json"
	"errors"
	"os"

	"google.golang.org/api/youtube/v3"
)

func SavePlaylistToJSON(items []*youtube.PlaylistItem, total int) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	file, err := os.Create("db/playlist.json")
	file.Write(b)
	if err != nil {
		return err
	}

	rawData, err := os.ReadFile("db/sync.json")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	var db []SyncPlaylistItem
	if len(rawData) != 0 {
		err = json.Unmarshal(rawData, &db)
		if err != nil {
			return err
		}
	}

	for _, item := range items {
		contains := dbContainsItem(item, db)
		if !contains {
			db = append(db, SyncPlaylistItem{
				ID:         item.ContentDetails.VideoId,
				Title:      item.Snippet.Title + ".mp3",
				Downloaded: false,
			})
		}
	}

	dbJson, err := json.Marshal(&db)
	if err != nil {
		return err
	}

	dbFile, err := os.Create("db/sync.json")
	if err != nil {
		return err
	}

	dbFile.Write(dbJson)
	dbFile.Close()

	return nil
}

func dbContainsItem(item *youtube.PlaylistItem, db []SyncPlaylistItem) bool {
	for _, dbItem := range db {
		if dbItem.Title == item.Snippet.Title {
			return true
		}
	}
	return false
}
