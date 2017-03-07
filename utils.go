package starter

import (
	"fmt"
	"os"
)

// Отправляет сообщение в Stdout
func Print(format string, v ...interface{}) error {
	_, err := fmt.Fprintf(os.Stdout, format, v...)
	return err
}

// Отправляет сообщение в Stdout и добавляет перевед строки
func Println(format string, v ...interface{}) error {
	_, err := fmt.Fprintf(os.Stdout, format+"\n", v...)
	return err
}

// Отправляет сообщение в Stderr
func Errorln(format string, v ...interface{}) error {
	_, err := fmt.Fprintf(os.Stderr, format+"\n", v...)
	return err
}

// Отправляет сообщение в Stderr и заверщает программу с кодом 1
func Fatalln(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", v...)
	os.Exit(1)
}
