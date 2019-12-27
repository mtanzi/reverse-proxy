package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mtanzi/reverse-proxy/proxy"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getListenAddress() string {
	port := getEnv("UPSTREAM_PORT", "443")
	return ":" + port
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	t := proxy.NewProxyServer(res, req)
	t.ServeHTTP()
}

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)

	log.Printf("Server listening on... https://localhost%v\n", getListenAddress())
	if err := http.ListenAndServeTLS(getListenAddress(), "certs/server.crt", "certs/server.key", nil); err != nil {
		panic(err)
	}
}
