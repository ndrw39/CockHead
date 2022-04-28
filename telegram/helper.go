package telegram

import (
	"bytes"
	"cockHead/common"
	"cockHead/conf"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func SendMessage(userId int, message interface{}) {
	values := map[string]interface{}{
		"chat_id": strconv.Itoa(userId),
		"text":    message,
	}

	jsonValue, _ := json.Marshal(values)
	buffer := bytes.NewBuffer(jsonValue)
	resp, err := http.Post(conf.ApiUrl+"/sendMessage", "application/json", buffer)

	common.ErrorHandler(err, false)
	common.ErrorHandler(resp.Body.Close(), true)
}

func SendPhoto(userId int, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	client := &http.Client{Timeout: time.Second * 10}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	//chat_id
	fw, err := writer.CreateFormField("chat_id")
	if common.ErrorHandler(err, false) != false {
		return
	}
	_, err = io.Copy(fw, strings.NewReader(strconv.Itoa(userId)))
	if common.ErrorHandler(err, false) != false {
		return
	}

	//photo
	fw, err = writer.CreateFormFile("photo", path)
	if common.ErrorHandler(err, false) != false {
		return
	}
	file, err := os.Open(path)
	if common.ErrorHandler(err, false) != false {
		return
	}
	_, err = io.Copy(fw, file)
	if common.ErrorHandler(err, false) != false {
		return
	}

	err = writer.Close()
	if common.ErrorHandler(err, false) != false {
		return
	}

	// request
	req, err := http.NewRequest("POST", conf.ApiUrl+"/sendPhoto", bytes.NewReader(body.Bytes()))
	if common.ErrorHandler(err, false) != false {
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Println(rsp)
	}
}
