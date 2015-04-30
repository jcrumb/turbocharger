package main

import (
	"bytes"
	"fmt"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	proxy.OnRequest(goproxy.DstHostIs("prod.ros.rockstargames.com")).DoFunc(modifyOrder)
	proxy.OnResponse(goproxy.DstHostIs("prod.ros.rockstargames.com")).DoFunc(viewResponse)
	log.Fatal(http.ListenAndServe(":8228", proxy))
}

func modifyOrder(request *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	log.Printf("Request received for: %s", request.RequestURI)
	log.Printf("Method: %s", request.Method)

	matched, err := regexp.MatchString("/GTA5/car/", request.RequestURI)
	if err != nil {
		log.Fatal(err)
	}

	if matched && request.Method == "POST" {
		body, err := ioutil.ReadAll(request.Body)
		bodyString := fmt.Sprintf("%s", body)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Body: %s", body)

		bodyString = strings.Replace(bodyString, "\"11\":\"0\"", "\"11\":\"1\"", 1)
		bodyString = strings.Replace(bodyString, "\"31\":\"0\"", "\"31\":\"1\"", 1)
		bodyString = strings.Replace(bodyString, "\"12\":\"0\"", "\"12\":\"4\"", 1)
		bodyString = strings.Replace(bodyString, "\"13\":\"0\"", "\"13\":\"3\"", 1)
		bodyString = strings.Replace(bodyString, "\"30\":\"0\"", "\"30\":\"5\"", 1)
		bodyString = strings.Replace(bodyString, "AINTCARE", "TURBOCHG", 1)

		modifiedBodyBuffer := bytes.NewBufferString(bodyString)
		newBody := ioutil.NopCloser(modifiedBodyBuffer)

		request.Body = newBody
	}
	return request, nil
}

func viewResponse(response *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	log.Printf("Response: %s", response.Status)

	return response
}
