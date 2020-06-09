package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	_ "io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const portalUrl = "https://siasky.net"
const portalUploadPath = "/skynet/skyfile"

type UploadReponse struct {
	Skylink string `json:"skylink"`
}

func OpenFile(path string) (*os.File, string, error) {
	handle, err := os.Open(path)
	filename := filepath.Base(path)

	if err != nil {
		return nil, "", err
	}

	return handle, filename, nil
}

func main() {
	path := os.Args[1]

	// open the file
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	filename := filepath.Base(path)
	//fileInfo, _ := file.Stat()

	// prepare formdata
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%s/%s?filename=%s", strings.TrimRight(portalUrl, "/"), strings.TrimLeft(portalUploadPath, "/"),
		filename)

	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		panic(err)
	}

	// upload the file to skynet
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// parse the response
	//body = &bytes.Buffer{}
	//_, err = body.ReadFrom(resp.Body)
	//if err != nil {
	//	panic(err)
	//}

	var apiResponse UploadReponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&apiResponse)

	if err != nil {
		panic(err)
	}

	fmt.Printf("https://siasky.net/%s/%s\n", apiResponse.Skylink, filename)
}
