package system

import (
	"github.com/SirMetathyst/zinc"
	"github.com/veandco/go-sdl2/sdl"
)

// SDLQuitSystem ...
type SDLQuitSystem struct {
	entityManager *zinc.EntityManager
}

// NewSDLQuitSystem ...
func NewSDLQuitSystem() *SDLQuitSystem {
	return &SDLQuitSystem{
		entityManager: zinc.Default(),
	}
}

// NewSDLQuitSystemWith ...
func NewSDLQuitSystemWith(e *zinc.EntityManager) *SDLQuitSystem {
	return &SDLQuitSystem{
		entityManager: e,
	}
}

// Cleanup ...
func (s SDLQuitSystem) Cleanup() {
	sdl.Quit()
}
