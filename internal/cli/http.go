package cli

import (
	"bytes"
	"github.com/go-kratos/kratos/v2/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func GetHttpData(params url.Values, rawurl string) ([]byte, error) {
	Url, err := url.Parse(rawurl)
	if err != nil {
		panic(err.Error())

	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func PostHttpData(rawurl string, timeout time.Duration, jsonStr []byte, log *log.Helper) ([]byte, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	// client := &http.Client{}
	req, err := http.NewRequest("POST", rawurl, bytes.NewBuffer(jsonStr))
	if err != nil {
		// handle error
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Get data err:", err)
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
