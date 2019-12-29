package system

import (
	"github.com/SirMetathyst/atom"
	sdlatom "github.com/SirMetathyst/sdl2kit/sdl"
	"github.com/veandco/go-sdl2/sdl"
)

// SDLCreateWindowSystem ...
type SDLCreateWindowSystem struct {
	collector     atom.C
	entityManager *atom.EntityManager
}

// NewSDLCreateWindowSystem ...
func NewSDLCreateWindowSystem() *SDLCreateWindowSystem {
	return &SDLCreateWindowSystem{
		entityManager: atom.Default(),
		collector:     atom.Default().CreateCollector(atom.Added(sdlatom.SDLCreateWindowKey)),
	}
}

// NewSDLCreateWindowSystemWith ...
func NewSDLCreateWindowSystemWith(e *atom.EntityManager) *SDLCreateWindowSystem {
	x := &SDLCreateWindowSystem{
		entityManager: e,
		collector:     e.CreateCollector(atom.AddedX(e, sdlatom.SDLCreateWindowKey)),
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
		cw := sdlatom.SDLCreateWindowX(s.entityManager, id)
		w := createWindow(cw.Title)
		wid := s.entityManager.CreateEntity()
		sdlatom.SetSDLWindowX(s.entityManager, wid, sdlatom.SDLWindowData{Window: w})
	}
	s.collector.ClearCollectedEntities()
}
