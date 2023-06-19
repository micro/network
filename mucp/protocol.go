package mucp

import (
	"net"
)

// Proto is a code representation of the MUCP protocol
type Proto struct {
	conn net.Conn
}

func (p *Proto) Connect() error {
	return nil
}

func (p *Proto) Close() error {
	return nil
}

func (p *Proto) Call(service, endpoint string, request interface{}) (response interface{}, err error) {
	return nil, nil
}

func (p *Proto) Stream(service, endpoint, request interface{}) (*Stream, error) {
	return nil, nil
}

func (p *Proto) Publish(topic string, message interface{}) error {
	return nil
}

