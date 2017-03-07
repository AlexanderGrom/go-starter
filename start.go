package starter

import (
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
		Println(appName + " already running!")
		return 0
	}
	pidFile.Unlock()
	prog := appPath
	args := os.Args[2:]
	exec.Command(prog, args...).Start()
	Println(appName + " starting... started!")
	return 0
}
