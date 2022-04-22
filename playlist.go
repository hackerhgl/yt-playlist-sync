package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func FakeCount() int {
	return 203
}


func Count() int {
	commandStr := fmt.Sprint("--playlist-items 0-1 -J -i %s", PLAYLIST)
	commands := strings.Split(commandStr, " ")
	cmd := exec.Command("yt-dlp", commands...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error());
		return 0
	}

	file, err := os.Create("db/info.json");

	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	file.Write(stdout)
	file.Close()
	playlist := Playlist{}
	if err := json.Unmarshal(stdout, &playlist); err != nil {
		fmt.Println("Unmarshal error", err.Error())
		return 0
	}

	return playlist.Count;

}