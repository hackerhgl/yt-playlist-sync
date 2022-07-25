package main

import (
	"strings"

	"google.golang.org/api/drive/v3"
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

func itemExistsInDrive(item ParsedItem, files []*drive.File) bool {
	for _, file := range files {
		compare := strings.Compare(item.Title, file.Name)
		if compare == 0 {
			return true
		}
	}
	return false
}

func GetIgnoreChannels(size int) []chan []string {
	channels := make([]chan []string, size)
	for i := 0; i < size; i++ {
		channels[i] = make(chan []string)
	}
	return channels
}

func GetValuesFromIgnoreChannels(channels []chan []string) []string {
	values := []string{}
	for _, channel := range channels {
		data := <-channel
		if data == nil {
			continue
		}
		values = append(values, data...)
	}
	return values
}

func isVideoIgnored(item ParsedItem, ignores []string) bool {
	for _, id := range ignores {
		if item.ID == id {
			return true
		}
	}
	return false
}

func GetFilteredPlaylist(playlist []ParsedItem, files []*drive.File, ignores []string) ([]ParsedItem, int) {
	// func GetFilteredPlaylist(playlist []*youtube.PlaylistItem, files []*drive.File, ignores []string) ([]*youtube.PlaylistItem, int) {
	var filtered []ParsedItem
	for _, item := range playlist {
		if !itemExistsInDrive(item, files) && !isVideoIgnored(item, ignores) {
			filtered = append(filtered, item)
		}
	}
	return filtered, len(filtered)
}
