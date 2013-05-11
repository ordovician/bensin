// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package main

import (
	. "bensin/component"
	. "bensin/graphics"
	. "geom2d"
	. "bensin/base"
)

type movableBehavior struct {
	transform *Transform
	body *Body
}

func NewMovableEntity(pos Point, poly Polygon) *Entity {
	movable := NewEntity(pos)
	behavior := new(movableBehavior)
	behavior.transform = &movable.Transform
	movable.Behavior = behavior
		
	visual := NewShapeVisual(poly)
	visual.SetColor(0, 1, 0, 0.5)
	movable.Visual = visual
	
	var body Body
	body.Speed = 0.5
	body.Orientation = Rad(180)
	body.Position = Point{3, 3}
	behavior.body = &body
	
	movable.Collider.Init(pos, DirectionWithAngle(body.Orientation), poly)

	return movable
}


func (mb *movableBehavior) Update(t, dt float64) {
	mb.body.Update(t, dt)
	mb.transform.Input(mb.body.Output())
	mb.transform.Update(t, dt)
}
