package main

import (
	"fmt"
	"log"
)

func main() {
	count := FakeCount()
	perBatch := count / BATCH
	const batchSize = BATCH+1
	var batches [batchSize]int

	for i:=0; i <batchSize; i++ {
		batches[i] = perBatch
	}
	
	batches[BATCH] = count - (perBatch * BATCH)
	
	for index, value := range batches {
		start := index * perBatch
		end := start + value
		fmt.Println("START",start,end)
	}

	client, err := Client()

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	items, total := GetPlayList(client)

	abc := len(items)

	print("iaaa", abc, "--", total)



	SavePlaylistToJSON(items)

}
