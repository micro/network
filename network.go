package network

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"tailscale.com/tsnet"
	"tailscale.com/types/logger"
)

type Network struct {
	Options Options
	Server  *tsnet.Server
}

type Options struct {
	// The network tkoen
	Token string
	// Network name
	Name string
}

type Option func(*Options)

func New(opts ...Option) *Network {
	options := Options{
		Name:  "network",
		Token: os.Getenv("MICRO_NETWORK_TOKEN"),
	}

	for _, o := range opts {
		o(&options)
	}

	srv := &tsnet.Server{
		Dir:       "./.network",
		Logf:      logger.Discard,
		Hostname:  options.Name,
		Ephemeral: true,
		AuthKey:   options.Token,
	}

	return &Network{
		Options: options,
		Server:  srv,
	}
}

func (n *Network) Close() error {
	return n.Server.Close()
}

func (n *Network) Connect() error {
	ln, err := n.Server.ListenFunnel("tcp", ":443")
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Printf("Listening on https://%v\n", n.Server.CertDomains()[0])

	// Parse the target URL
	apiURL, err := url.Parse("http://localhost:8080")
	if err != nil {
		return err
	}

	// Create the api proxy
	apiProxy := httputil.NewSingleHostReverseProxy(apiURL)

	// Parse the target URL
	webURL, err := url.Parse("http://localhost:8082")
	if err != nil {
		return err
	}

	// Create the web proxy
	webProxy := httputil.NewSingleHostReverseProxy(webURL)

	return http.Serve(ln, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// serve the api
		if strings.HasPrefix(r.URL.Path, "/api") {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
			apiProxy.ServeHTTP(w, r)
			return
		}

		// set the header
		r.Header.Set("Micro-API", "/api")

		if strings.HasPrefix(r.URL.Path, "/") {
			// serve the web proxy
			webProxy.ServeHTTP(w, r)
		}
	}))
}

// WithName sets the network name
func WithName(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// WithToken sets the network token
func WithToken(t string) Option {
	return func(o *Options) {
		o.Token = t
	}
}
