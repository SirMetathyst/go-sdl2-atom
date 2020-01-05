package system

import (
	"github.com/SirMetathyst/zinc"
)

// NewConvSystemsWith ...
func NewConvSystemsWith(e *zinc.EntityManager) []zinc.S {
	return []zinc.S{
		NewConvCreateWindowSystemWith(e),
	}
}

// NewConvSystems ...
func NewConvSystems() []zinc.S {
	return NewConvSystemsWith(zinc.Default())
}
