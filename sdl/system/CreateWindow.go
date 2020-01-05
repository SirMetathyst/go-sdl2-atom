package system

import (
	"github.com/SirMetathyst/zinc"
	sdlzinc "github.com/SirMetathyst/sdlkit/sdl"
	"github.com/veandco/go-sdl2/sdl"
)

// SDLCreateWindowSystem ...
type SDLCreateWindowSystem struct {
	collector     zinc.C
	entityManager *zinc.EntityManager
}

// NewSDLCreateWindowSystem ...
func NewSDLCreateWindowSystem() *SDLCreateWindowSystem {
	return &SDLCreateWindowSystem{
		entityManager: zinc.Default(),
		collector:     zinc.Default().CreateCollector(zinc.Added(sdlzinc.SDLCreateWindowKey)),
	}
}

// NewSDLCreateWindowSystemWith ...
func NewSDLCreateWindowSystemWith(e *zinc.EntityManager) *SDLCreateWindowSystem {
	x := &SDLCreateWindowSystem{
		entityManager: e,
		collector:     e.CreateCollector(zinc.Added(sdlzinc.SDLCreateWindowKey)),
	}
	return x
}

func createWindow(title string) *sdl.Window {
	w, err := sdl.CreateWindow(
		title, 
		sdl.WINDOWPOS_UNDEFINED, 
		sdl.WINDOWPOS_UNDEFINED, 
		800, 600, 
		sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	return w
}

// Update ...
func (s SDLCreateWindowSystem) Update(dt float64) {
	for _, id := range s.collector.Entities() {
		cw := sdlzinc.SDLCreateWindowX(s.entityManager, id)
		w := createWindow(cw.Title)
		wid := s.entityManager.CreateEntity()
		sdlzinc.SetSDLWindowX(s.entityManager, wid, sdlzinc.SDLWindowData{Window: w})
	}
	s.collector.ClearCollectedEntities()
}
