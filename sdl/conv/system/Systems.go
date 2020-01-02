package system

import (
	"github.com/SirMetathyst/atom"
)

// NewConvSystemsWith ...
func NewConvSystemsWith(e *atom.EntityManager) []atom.S {
	return []atom.S{
		NewConvCreateWindowSystemWith(e),
	}
}

// NewConvSystems ...
func NewConvSystems() []atom.S {
	return NewConvSystemsWith(atom.Default())
}
