package main

import (
	"fmt"
	"github.com/Spartan09/lenslocked/models"
)

func main() {
	gs := models.GalleryService{}
	fmt.Println(gs.Images(2))
}
