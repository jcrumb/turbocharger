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

	// Only look at requests to the rockstar servers. Everything else is passed transparently
	proxy.OnRequest(goproxy.DstHostIs("prod.ros.rockstargames.com")).DoFunc(checkUrl)
	proxy.OnResponse(goproxy.DstHostIs("prod.ros.rockstargames.com")).DoFunc(viewResponse)
	log.Fatal(http.ListenAndServe(":8228", proxy))
}

func checkUrl(request *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	log.Printf("Request received for: %s", request.RequestURI)
	log.Printf("Method: %s", request.Method)

	matched, err := regexp.MatchString("/GTA5/car/", request.RequestURI)
	if err != nil {
		log.Fatal(err)
	}

	if matched && request.Method == "POST" {
		request = modifyRequest(request)
	}
	return request, nil
}

func viewResponse(response *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	log.Printf("Response: %s", response.Status)

	return response
}

func modifyRequest(request *http.Request) *http.Request {
	body, err := ioutil.ReadAll(request.Body)
	bodyString := fmt.Sprintf("%s", body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Body: %s", body)

	// Add armour, full ecu tune, brakes, and turbo
	// See https://docs.google.com/a/toxicedge.com/spreadsheet/ccc?key=0AixUkyNxN55gdF83LWI1MVFaeE9CY0ptdFEyYVFPV3c&usp=sharing#gid=0
	// for information on the fields/values required for this
	bodyString = strings.Replace(bodyString, "\"11\":\"0\"", "\"11\":\"1\"", 1)
	bodyString = strings.Replace(bodyString, "\"31\":\"0\"", "\"31\":\"1\"", 1)
	bodyString = strings.Replace(bodyString, "\"12\":\"0\"", "\"12\":\"4\"", 1)
	bodyString = strings.Replace(bodyString, "\"13\":\"0\"", "\"13\":\"3\"", 1)
	bodyString = strings.Replace(bodyString, "\"30\":\"0\"", "\"30\":\"5\"", 1)

	modifiedBodyBuffer := bytes.NewBufferString(bodyString)

	// ioutil.NopCloser gives us an io.Closer from a byte buffer, required by the request type
	newBody := ioutil.NopCloser(modifiedBodyBuffer)

	request.Body = newBody

	return request
}
