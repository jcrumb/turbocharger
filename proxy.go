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

func startProxy(port string) {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	// Only look at requests to the rockstar servers. Everything else is denied to prevent abuse
	proxy.OnRequest(goproxy.DstHostIs("prod.ros.rockstargames.com")).DoFunc(checkUrl)
	proxy.OnResponse(goproxy.DstHostIs("prod.ros.rockstargames.com")).DoFunc(viewResponse)

	matchRockstar, _ := regexp.Compile(".*rockstar*.")
	proxy.OnRequest(goproxy.Not(goproxy.ReqHostMatches(matchRockstar))).DoFunc(denyRequest)

	log.Println("Proxy listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, proxy))
}

func denyRequest(request *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	log.Printf("Denying request for: %s", request.RequestURI)
	return request, goproxy.NewResponse(request,
		goproxy.ContentTypeText, http.StatusForbidden,
		"Please disconnect from the turbocharger proxy")
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
	if response == nil {
		log.Printf("Error: %s", ctx.Error)
	} else {
		log.Printf("Response: %s", response.Status)
	}
	return response
}

func modifyRequest(request *http.Request) *http.Request {
	body, err := ioutil.ReadAll(request.Body)
	bodyString := fmt.Sprintf("%s", body)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Body:\n%s", body)

	// The body has some extra lines ahead of the actual json data we want
	// So split on newlines, and then join back again after
	lines := strings.Split(bodyString, "\n")

	orderStruct := rockstarOrderFromString(lines[4])
	orderStruct = applyOrder(orderStruct)
	jsonString := rockstarOrderToString(orderStruct)

	lines[4] = jsonString
	bodyString = strings.Join(lines, "\n")

	modifiedBodyBuffer := bytes.NewBufferString(bodyString)

	// ioutil.NopCloser gives us an io.Closer from a byte buffer, required by the request type
	newBody := ioutil.NopCloser(modifiedBodyBuffer)

	request.Body = newBody
	request.ContentLength = int64(len(bodyString))
	return request
}

func findAndReplace(source string, target string, replacement string) string {
	return strings.Replace(source, target, replacement, 1)
}
