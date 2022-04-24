package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/rs/zerolog"
)



func DummyShell(worker int, start int, end int)  {
	// var logFile *os.File 
	// initWorkerLogs(worker, logFile)
	path := fmt.Sprintf("logs/worker-%d.log", worker)
	// os.Remove(path)
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		println("Error: initWorkerLogs")
		return
	}
	logger := zerolog.New(logFile).With().Timestamp().Logger()
	for i:=start; i<=end; i++ {
		cmd := exec.Command("./scripts/echo.sh", strconv.Itoa(worker), strconv.Itoa(i))

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			// return 0, err
		}

		err = cmd.Start()

		if err != nil {
			println("ERROR WORKER", worker, start, err.Error())
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