package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const portalUrl = "https://siasky.net"
const portalUploadPath = "/skynet/skyfile"

type UploadReponse struct {
	Skylink string `json:"skylink"`
}

func main() {
	if len(os.Args) < 2 {
		c := color.New(color.FgYellow)
		c.Printf("🏥 Usage: %s [filename]\n", os.Args[0])
		return
	}
	path := os.Args[1]

	// open the file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	filename := filepath.Base(path)
	fileInfo, _ := file.Stat()

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
	defer writer.Close()

	url := fmt.Sprintf("%s/%s?dryrun=true&filename=%s", strings.TrimRight(portalUrl, "/"),
		strings.TrimLeft(portalUploadPath, "/"), filename)

	tmpl := `{{ green "uploading ⏳" }} {{ bar . "[" "-" (cycle . "↖" "↗" "↘" "↙" ) "." "]"}} {{speed . "%s/s" | green }} {{percent .}}`

	bar := pb.New(int(fileInfo.Size()))
	bar.SetTemplateString(tmpl)
	bar.Set(pb.SIBytesPrefix, true)
	bar.SetWidth(80)
	bar.Start()
	reader := bar.NewProxyReader(body)

	req, err := http.NewRequest("POST", url, reader)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var apiResponse UploadReponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&apiResponse)

	bar.Finish()

	c := color.New(color.FgGreen)
	c.Printf("clipboard 💥 ")

	skyLink := fmt.Sprintf("https://siasky.net/%s/%s", apiResponse.Skylink, filename)

	clipboard.WriteAll(skyLink)

	if err != nil {
		panic(err)
	}

	fmt.Println(skyLink)
}
