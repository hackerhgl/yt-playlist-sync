package main

import (
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/youtube/v3"
)

func CalculateBatches(total int) (batchesA []int, perBatchA int, batchesSizeA int) {
	safeBatches := total / MIN_PER_BATCH

	if (total % MIN_PER_BATCH) != 0 {
		safeBatches++
	}

	if safeBatches > MAX_BATCHES {
		safeBatches = MAX_BATCHES
	}

	perBatch := total / safeBatches
	var batches []int

	for i := 0; i < safeBatches; i++ {
		batches = append(batches, perBatch)
	}

	if (total % perBatch) != 0 {
		batches = append(batches, total-(perBatch*safeBatches))
	}

	return batches, perBatch, len(batches)
}

func CheckIfItemExists(item *youtube.PlaylistItem, files []*drive.File) bool {
	for _, file := range files {
		if item.Snippet.Title == file.Name {
			return true
		}
	}
	return false
}
