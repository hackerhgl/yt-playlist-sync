package main

import (
	"log"
)

func main() {
	InitDirs()

	driveClient, err := DriveClient()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	InitRootDir(driveClient)
	err = UploadAudio(driveClient, "Adele - Skyfall (Lyric Video).mp3")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
