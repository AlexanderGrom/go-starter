package starter

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

var cmdStop = &command{
	Name: "stop",
	Run:  cmdStopFunc,
}

func cmdStopFunc() int {
	pidFile := newPIDFile(pidPath)
	if err := pidFile.Lock(); err == nil {
		fmt.Println("App not running!")
		return 0
	}
	pid, err := pidFile.Get()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	p, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("Process not found!")
		return 1
	}
	err = p.Signal(syscall.SIGTERM)
	if err != nil {
		fmt.Println("Signall not sent!")
		return 1
	}
	fmt.Println("Please wait...")
	for {
		if err := p.Signal(syscall.Signal(0)); err != nil {
			break
		}
		time.Sleep(500 * time.Microsecond)
	}
	fmt.Println("App stoped!")
	return 0
}
