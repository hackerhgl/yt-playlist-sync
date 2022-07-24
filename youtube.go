package main

import (
	"context"
	"log"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetPlayList(client *youtube.Service) (parsedList []ParsedItem, data []*youtube.PlaylistItem, count int) {
	call := client.PlaylistItems.List([]string{"contentDetails,id,snippet,status"})
	call.MaxResults(50)
	call.PlaylistId(PLAYLIST)

	var items []*youtube.PlaylistItem
	nextPage := ""
	total := 0

	for {
		if nextPage != "" {
			call.PageToken(nextPage)
		}

		response, err := call.Do()
		if err != nil {
			log.Fatal(err.Error())
			panic(err)
		}

		total = int(response.PageInfo.TotalResults)
		items = append(items, response.Items...)
		size := len(items)
		nextPage = response.NextPageToken

		if nextPage == "" || size >= total {
			break
		}
	}

	var filtered []*youtube.PlaylistItem
	var parsed []ParsedItem
	for index, item := range items {
		title := items[index].Snippet.Title
		if strings.Contains(title, "Deleted") {
			continue
		}
		cleanTitle := strings.ReplaceAll(item.Snippet.Title, "/", "|") + ".mp3"
		items[index].Snippet.Title = cleanTitle

		filtered = append(filtered, item)

		pItem := &ParsedItem{
			ID:    item.ContentDetails.VideoId,
			Title: item.Snippet.Title,
		}
		parsed = append(parsed, *pItem)
	}

	return parsed, filtered, len(filtered)
}

func YoutubeClient() (*youtube.Service, error) {
	context := context.Background()

	return youtube.NewService(context, option.WithCredentialsFile("secret/credential.json"))
}
