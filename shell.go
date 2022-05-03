package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/youtube/v3"
)

func WorkerShell(worker int, start int, end int, playlist []*youtube.PlaylistItem, ignoreChan chan []string, drive *drive.Service) {
	path := fmt.Sprintf("logs/worker-%d.log", worker)
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		println("Error: initWorkerLogs")
	}
	logger := zerolog.New(logFile).With().Timestamp().Logger()
	var ignores []string
	for i := start; i <= end; i++ {
		index := i - 1
		item := playlist[index]
		cmd := DownloadVideo(item)
		// cmd := exec.Command("./scripts/echo.sh", strconv.Itoa(worker), strconv.Itoa(i))

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			println("PIPEOUTERROR")
			logger.Fatal().Err(err).Msg(err.Error())
			// return 0, err
		}

		err = cmd.Start()

		if err != nil {
			fmt.Printf("ERROR WORKER %d %d\n%s\n", worker, i, err.Error())
			logger.Fatal().Err(err).Msg(err.Error())
		}

		in := bufio.NewScanner(stdout)

		for in.Scan() {
			text := in.Text()
			println(text)
			logger.Info().Msg(text)
		}
		if err := in.Err(); err != nil {
			println("err", err.Error())
			logger.Fatal().Err(err).Msg(err.Error())
		}
		err = UploadAudio(drive, item.Snippet.Title)
		fileNotExist := err != nil && os.IsNotExist(err)
		if fileNotExist {
			ignores = append(ignores, item.ContentDetails.VideoId)
		} else {
			logger.Fatal().Err(err).Msg(err.Error())
		}

	}
	ignoreChan <- ignores
	logFile.Close()
	fmt.Printf("Worker %d finished\n", worker)
}

// func initWorkerLogs(worker int, logFile *os.File ) {
// 	path := fmt.Sprintf("logs/worker-%d.log", worker)
// 	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		println("Error: initWorkerLogs")
// 		return
//     }
// 	log.SetOutput(logFile)
// }

// func closeWorkerLogs( logFile *os.File) {
// 	logFile.Close()
// }
