// Package registry uses the go-micro registry for selection
package registry

import (
	"github.com/micro/network/mucp/selector"
)

// NewSelector returns a new registry selector
func NewSelector(opts ...selector.Option) selector.Selector {
	return selector.NewSelector(opts...)
}
