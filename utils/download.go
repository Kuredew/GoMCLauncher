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

	var downloadingStr string
Loop:
	for {
		select {
		case <-t.C:
			currentDownloadingStr := fmt.Sprintf("   Downloading %v / %v bytes (%.2f%%)\n", resp.BytesComplete(), resp.Size(), 100*resp.Progress())

			if currentDownloadingStr != downloadingStr{
				downloadingStr = currentDownloadingStr
				fmt.Print(downloadingStr)
			}
		case <-resp.Done:
			break Loop
		}

	}
	if err := resp.Err(); err != nil {
		log.Printf("Downloading Failed. %v", err)
		return Download(savedPath, url)
	}
	log.Printf("Download saved to %v", savedPath)

	return true
}