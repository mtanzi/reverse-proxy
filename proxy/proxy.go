package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

const (
	downstreamURL  = "http://127.0.0.1"
	downstreamPort = "1333"
)

// ProxyServer is the reverse proxy struct
type ProxyServer struct {
	downstreamURL url.URL
	response      http.ResponseWriter
	request       *http.Request
}

// NewProxyServer return a new ProxyServer
func NewProxyServer(res http.ResponseWriter, req *http.Request) (t *ProxyServer) {
	newURL, err := buildURL()
	if err != nil {
		log.Fatal(err)
	}

	t = &ProxyServer{
		downstreamURL: *newURL,
		response:      res,
		request:       req,
	}
	return
}

func buildURL() (*url.URL, error) {
	defaultURL := getEnv("DOWNSTREAM_URL", "http://127.0.0.1")
	defaultPort := getEnv("DOWNSTREAM_PORT", "1333")
	url, err := url.Parse(defaultURL + ":" + defaultPort)

	return url, err
}

func (t ProxyServer) GetDownstreamURL() *url.URL {
	newURL, err := buildURL()
	if err != nil {
		fmt.Fprint(t.response, err)
	}

	return newURL
}

// ServeHTTP forward the call to the server downstream
func (t ProxyServer) ServeHTTP() {
	t.setDefaultHeaders()

	response, err := http.DefaultClient.Do(t.request)
	if err != nil {
		t.response.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(t.response, err)
		return
	}

	forwardedIP, _, _ := net.SplitHostPort(t.request.RemoteAddr)
	response.Header.Set("X-Forwarded-For", forwardedIP)

	t.copyHeader(response.Header)

	t.response.WriteHeader(response.StatusCode)
	io.Copy(t.response, response.Body)
	fmt.Printf("Request forwarded to downstream server: %v\n", t.GetDownstreamURL())
}

func (t ProxyServer) setDefaultHeaders() {
	t.request.Host = t.downstreamURL.Host
	t.request.URL.Host = t.downstreamURL.Host
	t.request.URL.Scheme = t.downstreamURL.Scheme
	t.request.RequestURI = ""
}

func (t ProxyServer) copyHeader(source http.Header) {
	for key, values := range source {
		for _, value := range values {
			t.response.Header().Add(key, value)
		}
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
