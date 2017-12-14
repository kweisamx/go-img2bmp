package main

import  "image/png"
import "golang.org/x/image/bmp"
import (
	"fmt"
	"os"
)

func main() {
	argc := len(os.Args)
	if argc < 3 {
		fmt.Println("Too few argument, please put intput and output")
		return
	}
    fmt.Println("Reading image...")
	/* Image Read */
	imagein, err := os.Open(os.Args[1])
	if err != nil {
		panic(err.Error())
	}
	defer imagein.Close()

	src, err := png.Decode(imagein)
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
