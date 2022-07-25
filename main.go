package main

import (
	"log"
	"sync"
)

func main() {
	err := TimeStamp()
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	err = InitDirs()

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	// ignores, err := GetIgnores()
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	youtubeClient, err := YoutubeClient()

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	playlist, _, total := GetPlayList(youtubeClient)
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
	files, err := GetDownloadedFiles(driveClient)
	// for _, file := range files {
	// 	// fmt.Printf("Files %+v\n\n\n", file.Name)
	// 	println(file.Name)
	// }

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	playlist, total = GetFilteredPlaylist(playlist, files, []string{})

	if total == 0 {
		println("Playlist is synced")
		return
	}

	batches, perBatch, batchesSize := CalculateBatches(total)
	var wg sync.WaitGroup
	wg.Add(batchesSize)

	// ignoresChannels := GetIgnoreChannels(batchesSize)
	// for _, c := range ignoresChannels {
	// 	defer close(c)
	// }
	for index, value := range batches {
		multiplier := (index * perBatch)
		start := multiplier + 1
		end := multiplier + value
		go func(index int) {
			WorkerShell(index, start, end, playlist, driveClient)
			wg.Done()
		}(index)
	}
	// ignores = append(ignores, GetValuesFromIgnoreChannels(ignoresChannels)...)
	wg.Wait()
	// SyncIgnores(ignores)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
