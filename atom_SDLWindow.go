// Package atomcommon ...
// Generated by the atom tool.  DO NOT EDIT!
// Source: atom_SDLWindow
package atomcommon

import (
	"github.com/SirMetathyst/atom"
	"github.com/veandco/go-sdl2/sdl"
)

// SDLWindowKey ...
const SDLWindowKey uint = 135545206

// SDLWindowData ...
type SDLWindowData struct {
	Window *sdl.Window
}

// SDLWindowComponent ...
type SDLWindowComponent struct {
	context atom.Context
	data    map[atom.EntityID]SDLWindowData
}

// NewSDLWindowComponent ...
func NewSDLWindowComponent() *SDLWindowComponent {
	return &SDLWindowComponent{
		data: make(map[atom.EntityID]SDLWindowData),
	}
}

// SetContext ...
func (c *SDLWindowComponent) SetContext(ctx atom.Context) {
	if c.context == nil {
		c.context = ctx
	}
}

func init() {
	x := NewSDLWindowComponent()
	context := atom.Default().RegisterComponent(SDLWindowKey, x)
	x.SetContext(context)
}

// DeleteEntity ...
func (c *SDLWindowComponent) DeleteEntity(id atom.EntityID) {
	delete(c.data, id)
}

// HasEntity ...
func (c *SDLWindowComponent) HasEntity(id atom.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// SetSDLWindow ...
func (c *SDLWindowComponent) SetSDLWindow(id atom.EntityID, sdlwindow SDLWindowData) {
	if c.context.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = sdlwindow
			c.context.ComponentUpdated(SDLWindowKey, id)
		} else {
			c.data[id] = sdlwindow
			c.context.ComponentAdded(SDLWindowKey, id)
		}
	}
}

// SDLWindow ...
func (c *SDLWindowComponent) SDLWindow(id atom.EntityID) SDLWindowData {
	return c.data[id]
}

// DeleteSDLWindow ...
func (c *SDLWindowComponent) DeleteSDLWindow(id atom.EntityID) {
	delete(c.data, id)
	c.context.ComponentDeleted(SDLWindowKey, id)
}

// SetSDLWindowX ...
func SetSDLWindowX(e *atom.EntityManager, id atom.EntityID, sdlwindow SDLWindowData) {
	v, _ := e.Component(SDLWindowKey)
	c := v.(*SDLWindowComponent)
	c.SetSDLWindow(id, sdlwindow)
}

// SetSDLWindow ...
func SetSDLWindow(id atom.EntityID, sdlwindow SDLWindowData) {
	SetSDLWindowX(atom.Default(), id, sdlwindow)
}

// SDLWindowX ...
func SDLWindowX(e *atom.EntityManager, id atom.EntityID) SDLWindowData {
	v, _ := e.Component(SDLWindowKey)
	c := v.(*SDLWindowComponent)
	return c.SDLWindow(id)
}

// SDLWindow ...
func SDLWindow(id atom.EntityID) SDLWindowData {
	return SDLWindowX(atom.Default(), id)
}

// DeleteSDLWindowX ...
func DeleteSDLWindowX(e *atom.EntityManager, id atom.EntityID) {
	v, _ := e.Component(SDLWindowKey)
	c := v.(*SDLWindowComponent)
	c.DeleteSDLWindow(id)
}

// DeleteSDLWindow ...
func DeleteSDLWindow(id atom.EntityID) {
	DeleteSDLWindowX(atom.Default(), id)
}

// HasSDLWindowX ...
func HasSDLWindowX(e *atom.EntityManager, id atom.EntityID) bool {
	v, _ := e.Component(SDLWindowKey)
	return v.HasEntity(id)
}

// HasSDLWindow ...
func HasSDLWindow(id atom.EntityID) bool {
	return HasSDLWindowX(atom.Default(), id)
}
