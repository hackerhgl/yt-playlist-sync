package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"google.golang.org/api/drive/v3"
)

func DummyShell(worker int, start int, end int, db []SyncPlaylistItem, drive *drive.Service) {
	path := fmt.Sprintf("logs/worker-%d.log", worker)
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		println("Error: initWorkerLogs")
		return
	}
	logger := zerolog.New(logFile).With().Timestamp().Logger()
	for i := start; i <= end; i++ {
		index := i - 1
		item := db[index]
		if item.Downloaded {
			continue
		}
		cmd := DownloadVideo(item.ID)
		// cmd := exec.Command("./scripts/echo.sh", strconv.Itoa(worker), strconv.Itoa(i))

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			// return 0, err
		}

		err = cmd.Start()

		if err != nil {
			println("ERROR WORKER", worker, start, err.Error())
			logger.Fatal().Err(err).Msg(err.Error())
			return
		}

		in := bufio.NewScanner(stdout)

		for in.Scan() {
			text := in.Text()
			println(text)
			logger.Info().Msg(text)
		}
		if err := in.Err(); err != nil {
			logger.Fatal().Err(err).Msg(err.Error())
		}
		if err := UploadAudio(drive, item.Title); err != nil {
			// logger.Fatal().Msg("GOOGLE DRIVE ERROR")
			logger.Fatal().Err(err).Msg(err.Error())
		}

		db[index].Downloaded = true
	}
	// closeWorkerLogs(logFile)
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
