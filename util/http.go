package util

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Get(url string) (body []byte) {
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	Resp, err := client.Get(url)
	if err != nil {
		log.Error(err)
	}
	defer Resp.Body.Close()
	body, errs := ioutil.ReadAll(Resp.Body)
	if errs != nil {
		log.Error(errs)
	}
	return body
}

//Get请求
func GetHttpResponse(method, urlVal string) (body []byte) {
	client := &http.Client{}
	req,err := http.NewRequest(method, urlVal, nil)
	if err != nil {
		log.Error(err)
	}
	//代码直接get不了，但浏览器可以访问
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200{
		s:=fmt.Sprintf("resp.StatusCode= %d",resp.StatusCode)
		log.Error(s)
	}
	body, errs:= ioutil.ReadAll(resp.Body)
	if errs!=nil{
		log.Error("Read resp.body failed!")
	}

	return body
}
//Post请求
func PostHttpResponse(url string, body string) (Body []byte) {
	payload := strings.NewReader(body)
	requests, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Error(err)
		return nil
	}
	requests.Header.Add("Accept", "application/json, text/plain, */*")
	requests.Header.Add("Content-Type", "application/json;charset=UTF-8")
	requests.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(requests)
	if err != nil {
		log.Error(err)
		return nil
	}
	if response.StatusCode != 200{
		s:=fmt.Sprintf("resp.StatusCode= %d",response.StatusCode)
		log.Error(s)
	}
	defer response.Body.Close()
	Body,errs:=ioutil.ReadAll(response.Body)
	if errs!=nil{
		log.Error("Read resp.body failed!")
	}
	return Body
}

