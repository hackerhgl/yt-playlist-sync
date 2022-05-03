package main

import (
	"fmt"
	"os/exec"
	"strings"

	"google.golang.org/api/youtube/v3"
)

func DownloadVideo(item *youtube.PlaylistItem) *exec.Cmd {
	commandStr := fmt.Sprintf("-f bestaudio --embed-thumbnail --add-metadata --extract-audio --audio-format mp3 --audio-quality 0 -o TITLE -- %s", item.ContentDetails.VideoId)
	commands := strings.Split(commandStr, " ")
	for i, command := range commands {
		if command == "TITLE" {
			commands[i] = item.Snippet.Title
		}
	}
	cmd := exec.Command("yt-dlp", commands...)
	cmd.Dir = ("./songs")
	return cmd
}
