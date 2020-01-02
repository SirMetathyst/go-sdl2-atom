package system

import (
	"github.com/SirMetathyst/atom"
)

// NewSDLSystemsWith ...
func NewSDLSystemsWith(e *atom.EntityManager) []atom.S {
	return []atom.S{
		NewSDLInitSystemWith(e),
		NewSDLCreateWindowSystemWith(e),
		NewSDLQuitSystemWith(e),
	}
}

// NewSDLSystems ...
func NewSDLSystems() []atom.S {
	return NewSDLSystemsWith(atom.Default())
}
