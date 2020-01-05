package system

import (
	"github.com/SirMetathyst/atom"
	"github.com/SirMetathyst/atomkit"
	sdlatom "github.com/SirMetathyst/sdl2kit/sdl"
)

// ConvCreateWindowSystem ...
type ConvCreateWindowSystem struct {
	collector     atom.C
	entityManager *atom.EntityManager
}

// NewConvCreateWindowSystem ...
func NewConvCreateWindowSystem() *ConvCreateWindowSystem {
	return &ConvCreateWindowSystem{
		entityManager: atom.Default(),
		collector:     atom.Default().CreateCollector(atom.Added(atomkit.CreateWindowKey)),
	}
}

// NewConvCreateWindowSystemWith ...
func NewConvCreateWindowSystemWith(e *atom.EntityManager) *ConvCreateWindowSystem {
	return &ConvCreateWindowSystem{
		entityManager: e,
		collector:     e.CreateCollector(atom.AddedX(e, atomkit.CreateWindowKey)),
	}
}

// Update ...
func (s ConvCreateWindowSystem) Update(dt float64) {
	for _, id := range s.collector.Entities() {
		cw := atomkit.CreateWindowX(s.entityManager, id)
		nid := s.entityManager.CreateEntity()
		sdlatom.SetSDLCreateWindowX(s.entityManager, nid, sdlatom.SDLCreateWindowData{Title: cw.Title})
		s.entityManager.DeleteEntity(id)
	}
	s.collector.ClearCollectedEntities()
}
