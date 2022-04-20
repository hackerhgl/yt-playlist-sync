package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	commands := strings.Split("--playlist-items 0-1 -J -i PLoSjAzdJQCyfsUatRxBnaes_4yz3-ONm-", " ")
	cmd := exec.Command("yt-dlp", commands...)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}


	file, err := os.Create("db/info.json");

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	file.Write(stdout)
	file.Close()

	playlist := Playlist{}
	
	if err := json.Unmarshal(stdout, &playlist); err != nil {
		fmt.Println("Unmarshal error", err.Error())
		return
	}

	println(playlist.Count)


}
