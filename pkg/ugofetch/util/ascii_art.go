package util

import (
	"log"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func Generate_AA(filePath string) string {
	flags := aic_package.DefaultFlags()

	flags.Dimensions = []int{50, 25}
	flags.Colored = true

	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		log.Fatal("Failed to convert to ascii art: ", err)
	}
	return asciiArt
}
