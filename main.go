package main

import (
	"Intersect/server"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {

	// Initialize Libraries

	// Initialize and defer Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}); err != nil {
		log.Fatalln("Sentry Init Error: ", err)
	}

	defer sentry.Flush(2 * time.Second)

	// Initialize Server
	server.Init()
}

///
