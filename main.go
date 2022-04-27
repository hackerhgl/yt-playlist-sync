package main

import (
	"sync"
)

func main() {

	// client, err := Client()

	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	panic(err)
	// }

	// items, total := GetPlayList(client)

	// SavePlaylistToJSON(items)

	// driveClient, err := DriveClient()
	// if err != nil {
	// 	log.Fatal("Failed to init drive client")
	// }

	// FakeWrite(driveClient)

	total := 10
	batches, perBatch, batchesSize := CalculateBatches(total)
	var wg sync.WaitGroup
	wg.Add(batchesSize)

	for index, value := range batches {
		multiplier := (index * perBatch)
		start := multiplier + 1
		end := multiplier + value
		println("X", start, end)
		go func(index int) {
			DummyShell(index, start, end)
			wg.Done()
		}(index)
	}

	wg.Wait()
}
