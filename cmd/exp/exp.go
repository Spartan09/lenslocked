package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://api.thecatapi.com/v1/images/search?limit=10"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	var catImagesJson []map[string]interface{}
	if err := json.Unmarshal(body, &catImagesJson); err != nil {
		fmt.Printf("Error unmarshalling response body: %v\n", err)
		return
	}

	var catImages []string
	for img := range catImagesJson {
		catImages = append(catImages, catImagesJson[img]["url"].(string))
	}

	fmt.Println(catImages)

}
