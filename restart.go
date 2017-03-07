package starter

var cmdRestart = &command{
	Name: "restart",
	Run:  cmdRestartFunc,
}

func cmdRestartFunc() int {
	cmdStop.Run()
	cmdStart.Run()
	return 0
}
