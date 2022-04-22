package main

func main() {
	count := FakeCount()
	perBatch := count / BATCH
	const batchSize = BATCH+1
	var batches [batchSize]int

	for i:=0; i <batchSize; i++ {
		batches[i] = perBatch
	}
	batches[BATCH] = count - (perBatch * BATCH)
	
	println(batches[4])


}
