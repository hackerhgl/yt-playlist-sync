package main

import (
	"os/exec"
	"strconv"
)

func DummyShell(worker int, start int, end int)  {
	for i:=start; i<=end; i++ {
		exec.Command("./scripts/echo.sh", strconv.Itoa(worker), strconv.Itoa(start)).Output()
	}
}