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
			currentDownloadingStr := fmt.Sprintf("   Downloading (%.2f%%)\r",100*resp.Progress())

			if currentDownloadingStr != downloadingStr{
				downloadingStr = currentDownloadingStr
				fmt.Print(downloadingStr)
			}
		case <-resp.Done:
			break Loop
		}

	}
	if err := resp.Err(); err != nil {
		fmt.Printf("   Downloading Failed. %v\n", err)
		return Download(savedPath, url)
	}
	fmt.Printf("   Download saved to %v\n", savedPath)

	return true
}