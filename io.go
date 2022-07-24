package main

import (
	"encoding/json"
	"os"
	"time"
)

func GetIgnores() ([]string, error) {
	b, err := os.ReadFile("db/ignores.json")
	if err != nil {
		return nil, err
	}

	var ignores []string
	if len(b) != 0 {
		err = json.Unmarshal(b, &ignores)
		if err != nil {
			return nil, err
		}
	}
	return ignores, nil
}

func SyncIgnores(ignores []string) error {
	file, err := os.Create("db/ignores.json")
	if ignores == nil {
		ignores = []string{}
	}
	if err != nil {
		return err
	}
	b, err := json.Marshal(&ignores)
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func TimeStamp() error {
	file, err := os.Create("db/timestamp.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(time.Now().String()))

	return nil
}

func SavePlaylistToJSON(items []ParsedItem, total int) error {
	b, err := json.MarshalIndent(items, "", "	")
	if err != nil {
		return err
	}
	file, err := os.Create("db/playlist.json")
	if err != nil {
		return err
	}
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
