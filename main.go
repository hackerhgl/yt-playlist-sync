package main

import (
	"log"
	"sync"
)

func main() {
	err := InitDirs()
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	youtubeClient, err := YoutubeClient()

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	playlist, total := GetPlayList(youtubeClient)

	err = SavePlaylistToJSON(playlist, total)

	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	driveClient, err := DriveClient()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	InitRootDir(driveClient)

	total = 1

	files, err := GetDownloadedFiles(driveClient)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	batches, perBatch, batchesSize := CalculateBatches(total)
	var wg sync.WaitGroup
	wg.Add(batchesSize)

	for index, value := range batches {
		multiplier := (index * perBatch)
		start := multiplier + 1
		end := multiplier + value
		go func(index int) {
			WorkerShell(index, start, end, files, playlist, driveClient)
			wg.Done()
		}(index)
	}

	wg.Wait()
}
