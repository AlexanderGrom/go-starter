
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
	"context"
	"log"
	"net/http"
	"time"

	"github.com/AlexanderGrom/go-starter"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!\n"))
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Registration function: Close listening port
	starter.Bind(func() {
		ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
		srv.Shutdown(ctx)
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil {
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

Correct Restart programm:
```bash
$ ./myapp restart
```

## -~- THE END -~-
