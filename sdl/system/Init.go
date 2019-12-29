package system

import (
	"github.com/SirMetathyst/atom"
	"github.com/veandco/go-sdl2/sdl"
)

// SDLInitSystem ...
type SDLInitSystem struct {
	entityManager *atom.EntityManager
}

// NewSDLInitSystem ...
func NewSDLInitSystem() *SDLInitSystem {
	return &SDLInitSystem{
		entityManager: atom.Default(),
	}
}

// NewSDLInitSystemWith ...
func NewSDLInitSystemWith(e *atom.EntityManager) *SDLInitSystem {
	return &SDLInitSystem{
		entityManager: e,
	}
}

// Initialize ...
func (s SDLInitSystem) Initialize() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}
