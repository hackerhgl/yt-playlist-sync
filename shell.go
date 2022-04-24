package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)



func DummyShell(worker int, start int, end int)  {
	var logFile *os.File 
	initWorkerLogs(worker, logFile)
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
			log.Printf(in.Text()) // write each line to your log, or anything you need
		}
		if err := in.Err(); err != nil {
			log.Printf("error: %s", err)
		}
	}
	closeWorkerLogs(logFile)
}


func initWorkerLogs(worker int, logFile *os.File ) {
	path := fmt.Sprintf("logs/worker-%d.log", worker)
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		println("Error: initWorkerLogs")
		return
    }
	log.SetOutput(logFile)
}

func closeWorkerLogs( logFile *os.File) {
	logFile.Close()
}