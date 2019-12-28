package logic

import (
	"github.com/SirMetathyst/atom"
	"github.com/SirMetathyst/atomcommon"
)

// CreateSDLWindowSystem ...
type CreateSDLWindowSystem struct {
	collector         atom.C
	entityManager *atom.EntityManager
}

// NewCreateSDLWindowSystem ...
func NewCreateSDLWindowSystem() *CreateSDLWindowSystem {
	return &CreateSDLWindowSystem{
		entityManager: atom.Default(),
		collector:         atom.Default().CreateCollector(atom.Added(atomcommon.WindowKey)),
	}
}

// NewCreateSDLWindowSystemWith ...
func NewCreateSDLWindowSystemWith(e *atom.EntityManager) *CreateSDLWindowSystem {
	return &CreateSDLWindowSystem{
		entityManager: e,
		collector:         e.CreateCollector(atom.Added(atomcommon.WindowKey)),
	}
}

// Update ...
func (s CreateSDLWindowSystem) Update(dt float32) {
	for _, id := range s.collector.Entities() {
		w := atomcommon.Window(id)
		fmt.Println(w.Title)
	}
	s.collector.ClearCollectedEntities()
}
