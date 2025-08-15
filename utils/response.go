package utils

import (
	"net/http"
	
	"fmt"
	"io"
	"log"
)

func Response(url string) ([]byte) {
	for {
		fmt.Printf("    GET %s\n", url)

		response, err := http.Get(url)
		if err != nil {
			log.Print("Failed, retrying...")
			continue
		}

		data, _ := io.ReadAll(response.Body)
		return data
	}
	//return response, err

	//return string(responseData)
}