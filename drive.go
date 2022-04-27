package main

import (
	"context"
	"log"
	"os"

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

func UploadAudio(service *drive.Service, name string) error {
	mimeType := "application/vnd.google-apps.audio"
	path := name + ".mp3"
	fullPath := "songs/" + path

	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	call := service.Files.Create(&drive.File{
		Parents:  []string{rootDirId},
		MimeType: mimeType,
		Name:     path,
	})
	call.Media(file)
	defer file.Close()
	_, err = call.Do()
	if err != nil {
		return err
	}
	err = os.Remove(fullPath)
	if err != nil {
		return err
	}

	return nil
}
