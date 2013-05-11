package base

import (
	. "geom2d"
	"bensin/component"
)

// A Game object. Players, monsters, bullets, rocks etc are Entities.
// An entity has a location and orientation relative to its parent. This is defined by
// the Local matrix. The World matrix is its absolute position and orientation in the world.
// it is calculated by multiplying the local matrix with the world matrix of the parent.
//
// Since there is no inheritance in Go, we use the decorator and composite pattern to implement
// concrete entity objects. Both Visual and Behavior should be defined by composition of other
// components or visuals or by chaining following the decorator pattern.
type Entity struct {
	component.Transform
	Visual component.Visual
	Behavior component.Component
	Collider component.Collider
	Children []Entity
}

type noBehavior struct {}
func (comp noBehavior) Update(t, dt float64) {}
type invisibleVisual struct {}
func (vis invisibleVisual) Input(pl Placement) {}
func (vis invisibleVisual) Render(t, dt float64) {}
func (vis invisibleVisual) SetColor(red, green, blue, alpha float64) {}
func (vis invisibleVisual) Color() (red, green, blue, alpha float64) { return 0,  0, 0, 0}

var inactive noBehavior
var invisible invisibleVisual

// Create a new entity without a parent at absolute position pos
func NewEntity(pos Point) *Entity {
	entity := new(Entity)
	entity.Transform = component.NewTransform(pos)
	entity.Visual = &invisible
	entity.Behavior = &inactive
	
	return entity
}

func (entity *Entity) Update(t, dt float64) {	

	entity.Behavior.Update(t, dt)
	entity.Collider.Input(entity.Output())
	entity.Collider.UpdateShapePlacement()
	for _, child := range entity.Children {
		child.Transform.SetParent(entity.World)
		child.Update(t, dt)
	}
}

// Draws entity on screen. 
func (entity *Entity) Render(t, dt float64) {
	entity.Visual.Input(entity.Transform.Output())
	entity.Visual.Render(t, dt)
}

func (entity *Entity) Inside(q Point) bool {
	return entity.Collider.Inside(q)
}

func (entity *Entity) Intersect(shape Shape) bool {
	return entity.Collider.Intersect(shape)
}
