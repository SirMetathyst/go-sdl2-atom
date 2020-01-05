package system

import (
	"github.com/SirMetathyst/zinc"
	"github.com/SirMetathyst/zinckit"
	sdlzinc "github.com/SirMetathyst/sdl2kit/sdl"
)

// ConvCreateWindowSystem ...
type ConvCreateWindowSystem struct {
	collector     zinc.C
	entityManager *zinc.EntityManager
}

// NewConvCreateWindowSystem ...
func NewConvCreateWindowSystem() *ConvCreateWindowSystem {
	return &ConvCreateWindowSystem{
		entityManager: zinc.Default(),
		collector:     zinc.Default().CreateCollector(zinc.Added(zinckit.CreateWindowKey)),
	}
}

// NewConvCreateWindowSystemWith ...
func NewConvCreateWindowSystemWith(e *zinc.EntityManager) *ConvCreateWindowSystem {
	return &ConvCreateWindowSystem{
		entityManager: e,
		collector:     e.CreateCollector(zinc.Added(zinckit.CreateWindowKey)),
	}
}

// Update ...
func (s ConvCreateWindowSystem) Update(dt float64) {
	for _, id := range s.collector.Entities() {
		cw := zinckit.CreateWindowX(s.entityManager, id)
		nid := s.entityManager.CreateEntity()
		sdlzinc.SetSDLCreateWindowX(s.entityManager, nid, sdlzinc.SDLCreateWindowData{Title: cw.Title})
		s.entityManager.DeleteEntity(id)
	}
	s.collector.ClearCollectedEntities()
}
