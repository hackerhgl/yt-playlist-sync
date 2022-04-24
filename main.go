package main

func main() {

	// client, err := Client()

	// if err != nil {
	// 	log.Fatal(err.Error())
	// 	panic(err)
	// }

	// items, total := GetPlayList(client)

	// SavePlaylistToJSON(items)

	total := 52

	safeBatches := total / MIN_PER_BATCH

	if safeBatches > MAX_BATCHES {
		safeBatches = MAX_BATCHES
	}

	perBatch := total / safeBatches
	var batches []int

	for i:=0; i <safeBatches; i++ {
		batches = append(batches, perBatch)
	}

	if (total % perBatch) != 0 {
		batches = append(batches, total - (perBatch * safeBatches))
	}

	for index, value := range batches {
		multiplier := (index * perBatch)
		start := multiplier + 1 
		end := multiplier + value
		println("START",start,end)
		DummyShell(index, start, end)
	}
}
