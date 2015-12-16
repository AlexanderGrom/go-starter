
## Starter - Start/Stop daemon in Golang

This is for UNIX based systems only.

## Install

```bash
$ go get github.com/AlexanderGrom/go-starter
```

## Simple Example

```go
package main

import (
	"github.com/AlexanderGrom/go-starter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!\n"))
	})

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error Listen:", err)
	}

	// Registration function: Close listening port
	starter.Bind(func() {
		l.Close()
	})

	go func() {
		if err := http.Serve(l, mux); err != nil {
			log.Println("Error Serve:", err)
		}
	}()

	// Registration function: Close cache
	starter.Bind(func() {
		time.Sleep(1 * time.Second) // Simulation
	})

	// Registration function: Close database connection
	starter.Bind(func() {
		time.Sleep(1 * time.Second) // Simulation
	})

	// Registration function: Close logs
	starter.Bind(func() {
		time.Sleep(1 * time.Second) // Simulation
	})

	// Wait until all functions completes
	starter.Wait()
}
```

Start programm:
```bash
$ ./myapp start
```

Correct Stop programm:
```bash
$ ./myapp stop
```

## -~- THE END -~-
