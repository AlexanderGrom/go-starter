package starter

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"
)

type command struct {
	Name string
	Run  func() int
}

var commands = []*command{
	cmdStart,
	cmdStop,
}

var (
	mutex       sync.Mutex
	closerFuncs []func()       = make([]func(), 0)
	doneChan    chan bool      = make(chan bool, 1)
	waitChan    chan bool      = make(chan bool, 1)
	signalChan  chan os.Signal = make(chan os.Signal, 1)
	appPath     string         = os.Args[0]
	appName     string         = path.Base(appPath)
	pidPath     string         = "/var/run/" + appName + ".pid"
)

//
// Авто инициализация
//
func init() {
	if len(os.Args) > 1 {
		name := os.Args[1]
		for _, cmd := range commands {
			if cmd.Name == name {
				os.Exit(cmd.Run())
			}
		}
	}
	createPIDFileLockAndSet(pidPath, os.Getpid())
	signalListen(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
}

//
// Привязываем функцию завершения
//
func Bind(fn func()) {
	mutex.Lock()
	c := make([]func(), 0, len(closerFuncs)+1)
	c = append(c, fn)
	closerFuncs = append(c, closerFuncs...)
	mutex.Unlock()
}

//
// Ждем пока не будет обработан выход
//
func Wait() {
	<-waitChan
	fmt.Println("Please wait...")
	<-doneChan
	fmt.Println("App stoped!")
	os.Exit(0)
}

//
// Слушает сигналы завершения
//
func signalListen(signalChan chan os.Signal, siganls ...os.Signal) {
	signal.Notify(signalChan, siganls...)
	go func() {
		<-signalChan
		mutex.Lock()
		defer mutex.Unlock()
		waitChan <- true
		for _, fn := range closerFuncs {
			fn()
		}
		doneChan <- true
	}()
}

//
// Создаем pid файл, блокируем и сохраняем в него новый pid
//
func createPIDFileLockAndSet(path string, pid int) {
	pidFile := newPIDFile(path)
	if err := pidFile.Lock(); err != nil {
		fmt.Println("PID not lock!")
		os.Exit(1)
	}
	if err := pidFile.Set(pid); err != nil {
		fmt.Println("PID not set!")
		os.Exit(1)
	}
}

//
// Обертка над os.File для работы с PID файлом
//
type pidFile struct {
	*os.File
}

func newPIDFile(path string) *pidFile {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("PID file not open!")
		os.Exit(0)
	}
	return &pidFile{file}
}

func (file *pidFile) Set(pid int) error {
	file.Truncate(0)
	file.Seek(0, os.SEEK_SET)
	_, err := fmt.Fprint(file, pid)
	if err != nil {
		return fmt.Errorf("PID not save!")
	}
	return nil
}

func (file *pidFile) Get() (int, error) {
	var pid int = 0
	_, err := fmt.Fscan(file, &pid)
	if err != nil {
		return 0, fmt.Errorf("PID not read!")
	}
	return pid, nil
}

func (file *pidFile) Lock() error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

func (file *pidFile) Unlock() error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}
