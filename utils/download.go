package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

func Download(savedPath string, url string) bool {
	client := grab.NewClient()
	req, _ := grab.NewRequest(savedPath, url)

	log.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("   Downloading %v / %v bytes (%.2f%%)\n", resp.BytesComplete(), resp.Size(), 100*resp.Progress())
		case <-resp.Done:
			break Loop
		}

	}
	if err := resp.Err(); err != nil {
		log.Fatalf("Downloading Failed. %v", err)
	}
	log.Printf("Download saved to %v", savedPath)

	return true
}