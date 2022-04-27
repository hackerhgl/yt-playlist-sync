package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func DownloadVideo(id string) *exec.Cmd {
	commandStr := fmt.Sprintf("-f bestaudio --embed-thumbnail --add-metadata --extract-audio --audio-format mp3 --audio-quality 0 -o %%(title)s.%%(ext)s %s", id)
	commands := strings.Split(commandStr, " ")
	cmd := exec.Command("yt-dlp", commands...)
	cmd.Dir = ("./songs")
	return cmd
}
