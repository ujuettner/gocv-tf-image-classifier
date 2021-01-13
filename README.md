A simple example using [GoCV](https://gocv.io/) (package docs are [here](https://pkg.go.dev/gocv.io/x/gocv)) and a [pre-trained](https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip) [TensorFlow](https://www.tensorflow.org/) model to classify an image.

## Get the Model Files

1. `wget https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip`
2. `unzip inception5h.zip`

## Development

GoCV is a wrapper around [OpenCV](https://opencv.org/) using cgo. Therefore, installing GoCV means compiling a lot of C code and installing libraries. In order to avoid cluttering your local development environment, you can use a prepared Docker container:

1. `docker pull gocv/opencv`
2. `docker run --tty --interactive --rm --volume "${PWD}":/src --name gocv-dev gocv/opencv bash`
