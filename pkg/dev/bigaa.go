package dev

import (
	"fmt"
	"log"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func generate_AA(filePath string) string {
	flags := aic_package.DefaultFlags()

	flags.Dimensions = []int{100, 50}
	flags.Colored = true

	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		log.Fatal("Failed to convert to ascii art: ", err)
	}
	return asciiArt
}

func BigAA(filePath string) {
	fmt.Printf("%v\n", generate_AA(filePath))
}
