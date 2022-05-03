package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

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

func GetDownloadedFiles(service *drive.Service) ([]*drive.File, error) {
	var files []*drive.File
	var nextPageToken string

	for {
		call := service.Files.List()
		fields := []googleapi.Field{"files/webViewLink", "files/name", "files/kind", "files/id", "files/mimeType"}
		call.Fields(fields...)
		if nextPageToken != "" {
			call.PageToken(nextPageToken)
		}
		query := fmt.Sprintf("mimeType='audio/mpeg' and parents='%s'", rootDirId)
		call.Q(query)
		call.PageSize(50)
		result, err := call.Do()
		if err != nil {
			log.Fatalln(err.Error())
			log.Fatalln("Failed execute drive query")
			return nil, err
		}
		files = append(files, result.Files...)
		nextPageToken = result.NextPageToken

		if result.NextPageToken == "" {
			break
		}
	}
	return files, nil
}

func DriveClient() (*drive.Service, error) {
	context := context.Background()
	credentials := option.WithCredentialsFile("secret/credential.json")
	scopes := option.WithScopes("https://www.googleapis.com/auth/drive")
	return drive.NewService(context, credentials, scopes)
}

func UploadAudio(service *drive.Service, name string) error {

	println("** UploadAudio **")
	// mimeType := "application/vnd.google-apps.audio"
	fullPath := filepath.Join("songs", name)
	file, err := os.Open(fullPath)
	// file, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return err
	}
	driveFile := &drive.File{
		Parents: []string{rootDirId},
		Name:    name,
	}

	call := service.Files.Create(driveFile).Media(file)
	// call := service.Files.Create(driveFile).Media(bytes.NewReader(file), googleapi.ContentType("audio/mp3"), googleapi.ChunkSize(0), googleapi.)
	// service.Cre
	call.SupportsAllDrives(true)
	// call.

	_, err = call.Do()
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	err = os.Remove(fullPath)
	if err != nil {
		return err
	}

	return nil
}
