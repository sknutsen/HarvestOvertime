package main

import (
	"HarvestOvertime/gui"
	"net/http"
	"time"
)

var client *http.Client

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	gui.GetGui(client)
}
