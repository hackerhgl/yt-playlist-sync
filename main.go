package main

import (
	"log"
	"sync"
)

func main() {
	InitDirs()

	youtubeClient, err := YoutubeClient()

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	items, total := GetPlayList(youtubeClient)

	db, err := SavePlaylistToJSON(items, total)

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

	// FakeWrite(driveClient)

	total = 10
	batches, perBatch, batchesSize := CalculateBatches(total)
	var wg sync.WaitGroup
	wg.Add(batchesSize)

	for index, value := range batches {
		multiplier := (index * perBatch)
		start := multiplier + 1
		end := multiplier + value
		println("X", start, end)
		go func(index int) {
			DummyShell(index, start, end, db, driveClient)
			wg.Done()
		}(index)
	}

	wg.Wait()
	SyncDB(db)
}
