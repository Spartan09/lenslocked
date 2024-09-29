package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getTenCats() []string {
	url := "https://api.thecatapi.com/v1/images/search?limit=10"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil
	}

	var catImagesJson []map[string]interface{}
	if err := json.Unmarshal(body, &catImagesJson); err != nil {
		fmt.Printf("Error unmarshalling response body: %v\n", err)
		return nil
	}

	var catImages []string
	for _, img := range catImagesJson {
		catImages = append(catImages, img["url"].(string))
	}

	return catImages
}

func GetCats() []string {
	var cats []string
	for i := 0; i < 2; i++ {
		cats = append(cats, getTenCats()...)
	}
	return cats
}
