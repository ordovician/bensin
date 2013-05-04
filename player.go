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

type playerBehavior struct {
	transform *Transform
	body *Body
}

func NewPlayer(pos Point) *Entity {
	player := NewEntity(pos)
	behavior := new(playerBehavior)
	behavior.transform = &player.Transform
	player.Behavior = behavior
	
	poly := Polygon{Point{-2, -1},
					Point{2, 0},
					Point{-2, 1}}
	visual := NewShapeVisual(poly)
	visual.Color = Color{0, 1, 0, 0.5}
	player.Visual = visual
	
	var body Body
	body.Speed = 1
	body.Orientation = Rad(180)
	body.Position = Point{3, 3}
	behavior.body = &body
	
	return player
}


func (pb *playerBehavior) Update(t, dt float64) {
	pb.body.Update(t, dt)
	pb.transform.Input(pb.body.Output())
	pb.transform.Update(t, dt)
}
