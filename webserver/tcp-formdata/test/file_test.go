package test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)
import (
	fserver "../../tcp-formdata"
)

func TestFile(t *testing.T) {
	fserver.Start("http://localhost:7777/")
}
func TestSend(t *testing.T) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", "test.txt")
	if err != nil {
		log.Panicln(err)
	}

	//打开文件句柄操作
	fh, err := os.Open("test.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		log.Panicln(err)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post("http://127.0.0.1:7777/upload", contentType, bodyBuf)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))
}
