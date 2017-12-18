package main

import "image/png"
import "golang.org/x/image/bmp"
import "image/jpeg"
import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	argc := len(os.Args)
	if argc < 3 {
		fmt.Println("Too few argument, please put intput and output")
		return
	}
	fmt.Println("Reading image...")
	/* Image Read */
	imageinput, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	defer imageinput.Close()

	/* Detect Image format */

	filenameSplit := strings.Split(os.Args[1], ".")
	format := filenameSplit[len(filenameSplit)-1]
	fmt.Println("The image format is " + strings.ToLower(format))
	var src image.Image
	/* Decode Image */
	switch strings.ToLower(format) {
	case "png":
		src, err = png.Decode(imageinput)
	case "jpg", "jpeg":
		src, err = jpeg.Decode(imageinput)
	default:
		fmt.Println("The " + format + " we don't support to convert")
		os.Exit(1)
	}

	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}

	/* Covert to bmp*/

	fmt.Println("Converting image...")
	outfile, err := os.Create(os.Args[2])
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}
	defer outfile.Close()

	err = bmp.Encode(outfile, src)
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}

	fmt.Println("Convert Success!")

}
