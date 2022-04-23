package main

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func GetPlayList(client *youtube.Service) (data []*youtube.PlaylistItem , count int)  {
	call := client.PlaylistItems.List([]string{"contentDetails,id"})
	// call.Id(PLAYLIST)
	call.MaxResults(50)
	call.PlaylistId(PLAYLIST)
	
	// call.

	var items []*youtube.PlaylistItem
	nextPage := "" 
	total := 0

	for {
		response, err := call.Do()
		if err != nil {
			log.Fatal(err.Error())
			panic(err)
		}
		if (nextPage != "") {
			call.PageToken(nextPage)
		}

		items := append(items, response.Items...)
		size := len(items)
		nextPage = response.NextPageToken
		
		if nextPage == "" || size >= total {
			break
		}
	}

	return items, total
}

func Client() (*youtube.Service, error) {
	context := context.Background()

	return youtube.NewService(context, option.WithCredentialsFile("secret/credential.json"))
}