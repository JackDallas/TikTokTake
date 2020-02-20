package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	tracker, err := NewTracker()
	if err != nil {
		log.Fatalf("Error starting Tracker: %v", err)
	}

	// socketServer, err := NewSocketServer(&tracker, 3000)
	// if err != nil {
	// 	log.Fatalf("Error starting Socket Server: %v", err)
	// }

	startHTTPServer(&tracker, 6070)
}
