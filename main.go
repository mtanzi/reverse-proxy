package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mtanzi/reverse-proxy/cmd"
	"github.com/mtanzi/reverse-proxy/config"
	"github.com/mtanzi/reverse-proxy/proxy"
)

const (
	defaultConfigPath = "config.json"
	portSSL           = "443"
	portDefault       = "8080"
)

var command cmd.Cmd
var cfg config.Config

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	t := proxy.NewProxyServer(res, req, cfg)
	t.ServeHTTP()
}

func configPath() string {
	path := command.ConfigPath
	if path == "" {
		return defaultConfigPath
	}

	return path
}

func main() {
	command = cmd.ParseCmd()
	cfg = config.InitConfig(configPath())

	http.HandleFunc("/", handleRequestAndRedirect)

	if cfg.SSL == true {
		log.Printf("Server listening on... https://localhost:%v\n", portSSL)
		if err := http.ListenAndServeTLS(":"+portSSL, "certs/server.crt", "certs/server.key", nil); err != nil {
			log.Fatal("ListenAndServeTLS: ", err)
		}
	} else {
		log.Printf("Server listening on... http://localhost:%v\n", portDefault)
		if err := http.ListenAndServe(":"+portDefault, nil); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}

}
