package main

import (
	"log"
	"sync"
)

func main() {

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

	// driveClient, err := DriveClient()
	// if err != nil {
	// 	log.Fatal("Failed to init drive client")
	// }

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
			DummyShell(index, start, end, db)
			wg.Done()
		}(index)
	}

	wg.Wait()
}
