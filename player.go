// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package main

import (
	. "bensin/component"
	. "bensin/graphics"
	. "geom2d"
)

type Player struct {
	body *Body
	transform *Transform
	visual *ShapeVisual
	sinkNodes []SinkNode
}

func NewPlayer() *Player {
	var player Player
	
	poly := Polygon{Point{-2, -1},
					Point{2, 0},
					Point{-2, 1}}
	visual := NewShapeVisual(poly)
	visual.Color = Color{0, 1, 0, 0.5}
	player.visual = visual
	
	var body Body
	body.Speed = 1
	body.Orientation = Rad(180)
	body.Position = Point{3, 3}
	player.body = &body
	
	var transform Transform
	transform.Parent = -1
	transform.Local = NewMatrix(body.Position, body.Orientation)
	player.transform = &transform
	
	bodyNode := SourceNode{true, &body}
	transNode := FilterNode{true, &transform, &bodyNode}
	visualNode := SinkNode{visual, &transNode}
	player.sinkNodes = append(player.sinkNodes, visualNode)
	
	return &player
}

func (player *Player) Update(t, dt float64) {
	for _, sink := range player.sinkNodes {
		sink.Update(t, dt)
	}
	
	
	// player.body.Advance(dt)
	// pos := player.body.Position
	// dir := player.body.Orientation
	// 
	// player.transform.Local = NewMatrix(pos, dir)
	// player.visual.Direction = player.transform.Local.Dir()
	// player.visual.Position = pos
}

// func (player *Player) Render(t, dt float64) {
// 	player.visual.Render(t, dt)
// }