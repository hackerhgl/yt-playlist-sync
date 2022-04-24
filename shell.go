package main

import (
	"os/exec"
	"strconv"
)

func DummyShell(worker int, start int, end int)  {
	for i:=start; i<=end; i++ {
		b, err := exec.Command("./scripts/echo.sh", strconv.Itoa(worker), strconv.Itoa(start)).Output()

		if err != nil {
			println("ERROR WORKER", worker, start, err.Error())
			return
		}

		println("SUCCESS")
		print(string(b))

	}
}