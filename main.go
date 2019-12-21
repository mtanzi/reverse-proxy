package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
)

func getListenAddress() string {
	return ":1338"
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url, err := url.Parse("https://127.0.0.1:443")
	if err != nil {
		log.Fatal(err)
	}

	req.Host = url.Host
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.RequestURI = ""

	forwardedIP, _, _ := net.SplitHostPort(req.RemoteAddr)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}

	response.Header.Set("X-Forwarded-For", forwardedIP)

	copyHeader(res.Header(), response.Header)

	res.WriteHeader(response.StatusCode)
	io.Copy(res, response.Body)
}

func copyHeader(destination, source http.Header) {
	for key, values := range source {
		for _, value := range values {
			destination.Add(key, value)
		}
	}
}

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}
