package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func DownloadVideo(item ParsedItem) *exec.Cmd {
	commandStr := fmt.Sprintf("-f bestaudio --embed-thumbnail --add-metadata --extract-audio --audio-format mp3 --audio-quality 0 -o TITLE -- %s", item.ID)
	commands := strings.Split(commandStr, " ")
	for i, command := range commands {
		if command == "TITLE" {
			commands[i] = item.Title
		}
	}
	cmd := exec.Command("yt-dlp", commands...)
	cmd.Dir = ("./songs")
	return cmd
}
