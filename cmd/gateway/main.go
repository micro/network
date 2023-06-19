package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/micro/network/mucp"
)

func handler() http.Handler {
	// Parse the target URL
	apiURL, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatal("Failed to parse api URL: ", err)
	}

	// Create the api proxy
	apiProxy := httputil.NewSingleHostReverseProxy(apiURL)

	// Parse the target URL
	grpcURL, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatal("Failed to parse grpc URL: ", err)
	}

	// Create the grpc proxy
	grpcProxy := httputil.NewSingleHostReverseProxy(grpcURL)

	// Parse the target URL
	webURL, err := url.Parse("http://localhost:8082")
	if err != nil {
		log.Fatal("Failed to parse web URL: ", err)
	}

	// Create the web proxy
	webProxy := httputil.NewSingleHostReverseProxy(webURL)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")

		// serve the api
		if strings.HasPrefix(ct, "application/json") {
			apiProxy.ServeHTTP(w, r)
			return
		}

		// serve grpc
		if strings.HasPrefix(ct, "application/grpc") {
			grpcProxy.ServeHTTP(w, r)
			return
		}

		// set the header
		r.Header.Set("Micro-API", "/api")

		if strings.HasPrefix(r.URL.Path, "/") {
			// serve the web proxy
			webProxy.ServeHTTP(w, r)
		}
	})
}

func main() {
	token := os.Getenv("MICRO_NETWORK_TOKEN")
	if len(token) == 0 {
		log.Fatal("Missing MICRO_NETWORK_TOKEN")
	}

	net := mucp.NewNetwork(
		mucp.WithName("gateway"),
		mucp.WithHandler(handler()),
		mucp.WithToken(token),
	)

	if err := net.Connect(); err != nil {
		log.Fatal("Connect error: ", err)
	}
	defer net.Close()
}
