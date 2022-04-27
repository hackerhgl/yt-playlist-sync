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
	// mimeType := "application/vnd.google-apps.audio"
	fullPath := "songs/" + name
	file, err := os.Open(fullPath)
	// file, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return err
	}

	println("XX")

	driveFile := &drive.File{
		Parents: []string{rootDirId},
		// MimeType: mimeType,
		Name: name,
	}
	call := service.Files.Create(driveFile).Media(file)
	// call := service.Files.Create(driveFile).Media(bytes.NewReader(file), googleapi.ContentType("audio/mp3"), googleapi.ChunkSize(0), googleapi.)
	// service.Cre
	call.SupportsAllDrives(true)
	// call.

	resp, err := call.Do()
	if err != nil {
		println("DOOO ERRROR")
		return err
	}

	println("resp")
	println(resp.Id)
	// err = os.Remove(fullPath)
	// if err != nil {
	// 	return err
	// }

	return nil
}
