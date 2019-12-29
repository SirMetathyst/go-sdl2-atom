package system

import (
	"github.com/SirMetathyst/atom"
	"github.com/veandco/go-sdl2/sdl"
)

// SDLQuitSystem ...
type SDLQuitSystem struct {
	entityManager *atom.EntityManager
}

// NewSDLQuitSystem ...
func NewSDLQuitSystem() *SDLQuitSystem {
	return &SDLQuitSystem{
		entityManager: atom.Default(),
	}
}

// NewSDLQuitSystemWith ...
func NewSDLQuitSystemWith(e *atom.EntityManager) *SDLQuitSystem {
	return &SDLQuitSystem{
		entityManager: e,
	}
}

// Cleanup ...
func (s SDLQuitSystem) Cleanup() {
	sdl.Quit()
}
