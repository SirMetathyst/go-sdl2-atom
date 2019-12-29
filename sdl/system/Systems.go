package system

import (
	"github.com/SirMetathyst/atom"
)

// NewSDLSystemsWith ...
func NewSDLSystemsWith(e *atom.EntityManager) []atom.System {
	return []atom.System{
		NewSDLInitSystemWith(e),
		NewSDLCreateWindowSystemWith(e),
		NewSDLQuitSystemWith(e),
	}
}

// NewSDLSystems ...
func NewSDLSystems() []atom.System {
	return NewSDLSystemsWith(atom.Default())
}
