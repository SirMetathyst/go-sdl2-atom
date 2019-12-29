package system

import (
	"github.com/SirMetathyst/atom"
)

// NewConvSystemsWith ...
func NewConvSystemsWith(e *atom.EntityManager) []atom.System {
	return []atom.System{
		NewConvCreateWindowSystemWith(e),
	}
}

// NewConvSystems ...
func NewConvSystems() []atom.System {
	return NewConvSystemsWith(atom.Default())
}
