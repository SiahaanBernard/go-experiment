package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	imgFile, err := os.Open("instance.png") // a QR code image

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// if you create a new image instead of loading from file, encode the image to buffer instead with png.Encode()

	// png.Encode(&buf, image)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	// Embed into an html without PNG file
	//fmt.Println(imgBase64Str)
	b64File, err := os.Open("file.base64")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer b64File.Close()
	imag, err := base64.StdEncoding.DecodeString(imgBase64Str)
	fmt.Println(imag)
}
