package starter

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdStart = &command{
	Name: "start",
	Run:  cmdStartFunc,
}

func cmdStartFunc() int {
	pidFile := newPIDFile(pidPath)
	if err := pidFile.Lock(); err != nil {
		fmt.Println("App already running!")
		return 0
	}
	pidFile.Unlock()
	prog := appPath
	args := os.Args[2:]
	cmd := exec.Command(prog, args...)
	cmd.Start()
	fmt.Println("App running!")
	return 0
}
