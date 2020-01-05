package system

import (
	"github.com/SirMetathyst/zinc"
	"github.com/veandco/go-sdl2/sdl"
)

// SDLInitSystem ...
type SDLInitSystem struct {
	entityManager *zinc.EntityManager
}

// NewSDLInitSystem ...
func NewSDLInitSystem() *SDLInitSystem {
	return &SDLInitSystem{
		entityManager: zinc.Default(),
	}
}

// NewSDLInitSystemWith ...
func NewSDLInitSystemWith(e *zinc.EntityManager) *SDLInitSystem {
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
