package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: p4bm image.png")
		return
	}

	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer file.Close()

	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	imageType := filepath.Ext(fileName)
	imageType = strings.TrimPrefix(imageType, ".")
	if imageType == "jpeg" {
		imageType = "jpg"
	}

	base64String := base64.StdEncoding.EncodeToString(imageData)

	fmt.Printf("<img src=\"data:image/%s;base64,%s\" alt=\"%s\"/>\n", imageType, base64String, fileName)
}
