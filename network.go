package network

import (
	"fmt"
	"net/http"
	"os"

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
	// Network handler
	Handler http.Handler
}

type Option func(*Options)

func New(opts ...Option) *Network {
	options := Options{
		Name: "network",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "<html><body><h1>Micro Network</h1>")
		}),
		Token: os.Getenv("MICRO_NETWORK_TOKEN"),
	}

	// parse the options
	for _, o := range opts {
		o(&options)
	}

	// create new network
	net := new(Network)

	// set the server
	net.Server = &tsnet.Server{
		Dir:       "./.network",
		Logf:      logger.Discard,
		Hostname:  options.Name,
		Ephemeral: true,
		AuthKey:   options.Token,
	}

	// set the options
	net.Options = options

	// return the network
	return net
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

	// serve the handler
	return http.Serve(ln, n.Options.Handler)
}

// WithHandler sets the handler used
func WithHandler(h http.Handler) Option {
	return func(o *Options) {
		o.Handler = h
	}
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
