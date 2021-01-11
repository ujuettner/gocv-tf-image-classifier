package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"gocv.io/x/gocv"
)

func main() {
	var modelFile *string = flag.StringP("model", "m", "", "TensorFlow model file")
	var descFile *string = flag.StringP("desc", "d", "", "description file containing the model's labels")
	var imageFile *string = flag.StringP("image", "i", "", "image file")
	flag.Parse()

	// check for required flags
	if *modelFile == "" {
		fmt.Printf("No model file given!")
		os.Exit(2)
	}
	if *descFile == "" {
		fmt.Println("No description file given!")
		os.Exit(2)
	}
	if *imageFile == "" {
		fmt.Println("No image file given!")
		os.Exit(2)
	}

	image := gocv.IMRead(*imageFile, gocv.IMReadColor)
	if image.Empty() {
		fmt.Printf("Error reading image from: %v\n", *imageFile)
		return
	}
	defer image.Close()

	fmt.Println("TBD!!!")
}
