package main

import (
	"context"
	"log"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

var rootDirId string

func InitRootDir(service *drive.Service) {
	call := service.Files.List()
	fields := []googleapi.Field{"files/webViewLink", "files/name", "files/kind", "files/id", "files/mimeType"}
	call.Fields(fields...)
	call.Q("mimeType='application/vnd.google-apps.folder' and name='Music'")
	result, err := call.Do()
	if err != nil {
		log.Fatalln(err.Error())
		log.Fatalln("Failed execute drive query")
		return
	}
	rootDirId = result.Files[0].Id
}

func DriveClient() (*drive.Service, error) {
	context := context.Background()
	credentials := option.WithCredentialsFile("secret/credential.json")
	scopes := option.WithScopes("https://www.googleapis.com/auth/drive")
	return drive.NewService(context, credentials, scopes)
}
