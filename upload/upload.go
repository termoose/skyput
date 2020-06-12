package upload

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

func Do(path, portalUrl string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("file open: %v", err)
	}
	defer file.Close()

	filename := filepath.Base(path)
	fileInfo, _ := file.Stat()

	// prepare formdata
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return fmt.Errorf("create form file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("io copy: %v", err)
	}
	writer.Close()

	url := fmt.Sprintf("%s/%s?dryrun=false&filename=%s", strings.TrimRight(portalUrl, "/"),
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
		return fmt.Errorf("create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("🚫 %s\n", resp.Status)
	}

	var apiResponse Reponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&apiResponse)

	if err != nil {
		return fmt.Errorf("json decode: %v", err)
	}

	bar.Finish()

	c := color.New(color.FgGreen)
	c.Printf("clipboard 💥 ")

	skyLink := fmt.Sprintf("%s/%s/%s", portalUrl, apiResponse.Skylink, filename)

	clipboard.WriteAll(skyLink)


	fmt.Println(skyLink)

	return nil
}
