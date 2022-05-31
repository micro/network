package transport

import (
	"github.com/micro/network/mucp/tunnel"
	"github.com/micro/network/mucp/transport"
)

type tunListener struct {
	l tunnel.Listener
}

func (t *tunListener) Addr() string {
	return t.l.Channel()
}

func (t *tunListener) Close() error {
	return t.l.Close()
}

func (t *tunListener) Accept(fn func(socket transport.Socket)) error {
	for {
		// accept connection
		c, err := t.l.Accept()
		if err != nil {
			return err
		}
		// execute the function
		go fn(c)
	}
}
