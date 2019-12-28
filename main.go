package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mtanzi/reverse-proxy/cmd"
	"github.com/mtanzi/reverse-proxy/config"
	"github.com/mtanzi/reverse-proxy/proxy"
)

var command cmd.Cmd
var cfg config.Config

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getListenAddress() string {
	var port = "443"
	if command.SSL == "false" {
		port = "8080"
	}

	return ":" + port
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	t := proxy.NewProxyServer(res, req, cfg)
	t.ServeHTTP()
}

func main() {
	command = cmd.ParseCmd()
	cfg = config.InitConfig()

	http.HandleFunc("/", handleRequestAndRedirect)

	if command.SSL == "true" {
		log.Printf("Server listening on... https://localhost%v\n", getListenAddress())
		if err := http.ListenAndServeTLS(getListenAddress(), "certs/server.crt", "certs/server.key", nil); err != nil {
			log.Fatal("ListenAndServeTLS: ", err)
		}
	} else {
		log.Printf("Server listening on... http://localhost%v\n", getListenAddress())
		if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}

}
