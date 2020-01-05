package system

import (
	"github.com/SirMetathyst/zinc"
)

// NewSDLSystemsWith ...
func NewSDLSystemsWith(e *zinc.EntityManager) []zinc.S {
	return []zinc.S{
		NewSDLInitSystemWith(e),
		NewSDLCreateWindowSystemWith(e),
		NewSDLQuitSystemWith(e),
	}
}

// NewSDLSystems ...
func NewSDLSystems() []zinc.S {
	return NewSDLSystemsWith(zinc.Default())
}
