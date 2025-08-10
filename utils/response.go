package utils

import (
	"net/http"
	
	//"fmt"
	"io"
	"log"
)

func Response(url string) ([]byte, error) {
	log.Printf("Fetching %s", url)

	response, err := http.Get(url)
	if err != nil {
		return []byte(""), err
	}

	data, err := io.ReadAll(response.Body)
	return data, err

	//return response, err

	//return string(responseData)
}