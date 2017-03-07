package starter

import (
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
		pidFile.Unlock()
		Println(appName + " not running!")
		return 0
	}
	pid, err := pidFile.Get()
	if err != nil {
		Errorln(err.Error())
		return 1
	}
	p, err := os.FindProcess(pid)
	if err != nil {
		Errorln("Process not found!")
		return 1
	}
	err = p.Signal(syscall.SIGTERM)
	if err != nil {
		Errorln("Signall not sent!")
		return 1
	}
	Print(appName + " stopping... ")
	for {
		if err := p.Signal(syscall.Signal(0)); err != nil {
			break
		}
		time.Sleep(500 * time.Microsecond)
	}
	Println("stopped!")
	return 0
}
