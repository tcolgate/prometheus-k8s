package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const CONFIG_FILE = "/etc/prometheus/prometheus.yml"
const BIND = "0.0.0.0:8080"

func check(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func main() {
	watcher, err := WatchFile(CONFIG_FILE, time.Second, func() {
		http.PostForm("http://prometheus:9090/-/reload", nil)
	})
	check(err)

	// Clean up
	defer func() {
		watcher.Close()
	}()

	fmt.Printf("Listening on '%s'....\n", BIND)
	err = http.ListenAndServe(BIND, nil)
	check(err)
}
