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
	motion *RotMotion
}

func NewMovableEntity(pos Point, poly Polygon) *Entity {
	movable := NewEntity(pos)
	behavior := new(movableBehavior)
	behavior.transform = &movable.Transform
	movable.Behavior = behavior
		
	visual := NewShapeVisual(poly)
	visual.SetColor(0, 1, 0, 0.5)
	movable.Visual = visual
	
	var motion RotMotion
	motion.Speed = 0.5
	motion.Orientation = Rad(180)
	motion.Position = Point{3, 3}
	behavior.motion = &motion
	
	movable.Collider.Init(pos, DirectionWithAngle(motion.Orientation), poly)

	return movable
}


func (mb *movableBehavior) Update(t, dt float64) {
	mb.motion.Update(t, dt)
	mb.transform.Input(mb.motion.Output())
	mb.transform.Update(t, dt)
}
