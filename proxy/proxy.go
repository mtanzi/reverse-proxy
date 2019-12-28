package proxy

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/mtanzi/reverse-proxy/config"
)

// ProxyServer is the reverse proxy struct
type ProxyServer struct {
	config   config.Config
	response http.ResponseWriter
	request  *http.Request
}

// NewProxyServer return a new ProxyServer
func NewProxyServer(res http.ResponseWriter, req *http.Request, cfg config.Config) (t *ProxyServer) {
	t = &ProxyServer{
		config:   cfg,
		response: res,
		request:  req,
	}
	return t
}

func (t ProxyServer) buildURL() (*url.URL, error) {
	var defaultPort string
	// newConfig := config.InitConfig()

	for _, rule := range t.config.Rules {
		rex := regexp.MustCompile(rule.Matcher)
		match := rex.FindStringSubmatch(t.request.URL.Path)

		if len(match) == 0 {
			if defaultPort == "" {
				defaultPort = t.config.DefaultPort
			}
		} else {
			defaultPort = rule.DownstreamPort
		}
	}

	defaultURL := t.config.DefaultURL
	url, err := url.Parse("http://" + defaultURL + ":" + defaultPort)

	return url, err
}

// ServeHTTP forward the call to the server downstream
func (t ProxyServer) ServeHTTP() {
	newURL, err := t.buildURL()
	if err != nil {
		t.response.WriteHeader(http.StatusInternalServerError)
		log.Fatal(t.response, err)
		return
	}

	t.setDefaultHeaders(*newURL)
	log.Printf("Forwarding request to: %v \n", newURL)

	response, err := http.DefaultClient.Do(t.request)
	if err != nil {
		t.response.WriteHeader(http.StatusInternalServerError)
		log.Fatal(t.response, err)
		return
	}

	forwardedIP, _, _ := net.SplitHostPort(t.request.RemoteAddr)
	response.Header.Set("X-Forwarded-For", forwardedIP)

	t.copyHeader(response.Header)

	t.response.WriteHeader(response.StatusCode)
	io.Copy(t.response, response.Body)
}

func (t ProxyServer) setDefaultHeaders(downstreamURL url.URL) {
	t.request.Host = downstreamURL.Host
	t.request.URL.Host = downstreamURL.Host
	t.request.URL.Scheme = downstreamURL.Scheme
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
