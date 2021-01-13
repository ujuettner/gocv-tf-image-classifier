package main

import (
	"bufio"
	"fmt"
	"image"
	"os"

	flag "github.com/spf13/pflag"
	"gocv.io/x/gocv"
)

func main() {
	modelFile := flag.StringP("model", "m", "", "TensorFlow model file")
	labelsFile := flag.StringP("labels", "l", "", "file containing the model's labels")
	imageFile := flag.StringP("image", "i", "", "image file")
	flag.Parse()

	// check for required flags
	if *modelFile == "" {
		fmt.Println("No model file given!")
		os.Exit(2)
	}
	if *labelsFile == "" {
		fmt.Println("No labels file given!")
		os.Exit(2)
	}
	if *imageFile == "" {
		fmt.Println("No image file given!")
		os.Exit(2)
	}

	labels, err := readLabelsFromFile(*labelsFile)
	if err != nil {
		fmt.Printf("Error reading labels from file %v: %v\n", *labelsFile, err)
		return
	}

	imageDataArray := gocv.IMRead(*imageFile, gocv.IMReadColor)
	defer imageDataArray.Close()
	if imageDataArray.Empty() {
		fmt.Printf("Error reading image file %v\n", *imageFile)
		return
	}
	defer imageDataArray.Close()

	// scale the image to 224x224, swap the red and the blue color channels (may improve classification) and preserve aspect ratio
	imageBlob := gocv.BlobFromImage(imageDataArray, 1.0, image.Point{224, 224}, gocv.NewScalar(0, 0, 0, 0), true, true)
	defer imageBlob.Close()

	tfNet := gocv.ReadNetFromTensorflow(*modelFile)
	defer tfNet.Close()

	tfNet.SetInput(imageBlob, "input")

	// run forward pass to compute output of the last layer, i.e. run the network
	tfNetLayers := tfNet.GetLayerNames()
	netProbabilities := tfNet.Forward(tfNetLayers[len(tfNetLayers)-1])
	defer netProbabilities.Close()

	// flatten the network's result (mutli-channel) matrix into a (single-channel) 1 x n matrix
	probabilities := netProbabilities.Reshape(1, 1)
	defer probabilities.Close()

	_, maxProbability, _, maxLoc := gocv.MinMaxLoc(probabilities)
	identifiedLabel := "unknown"
	if maxLoc.X < len(labels) {
		identifiedLabel = labels[maxLoc.X]
	}

	fmt.Printf("Found a %v in image file %v with a probability of %v\n", identifiedLabel, *imageFile, maxProbability)
}

func readLabelsFromFile(labelsFile string) ([]string, error) {
	labelsFileHandle, err := os.Open(labelsFile)
	if err != nil {
		return nil, err
	}
	defer labelsFileHandle.Close()

	var labels []string
	scanner := bufio.NewScanner(labelsFileHandle)
	for scanner.Scan() {
		labels = append(labels, scanner.Text())
	}

	return labels, scanner.Err()
}
