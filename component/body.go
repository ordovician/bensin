// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package component

import (
	// "time"
	// "math"
	. "geom2d"
	// "fmt"
)

type Body struct {
	Orientation Radian
	Position Point
	Speed float64
	AngVelocity Radian
	AngAccel Radian
}

func (body *Body) Input(pos Point, dir Direction) {
	body.Position = pos
	body.Orientation = dir.Angle()
}

func (body *Body) Output() (pos Point, dir Direction) {
	return body.Position, DirectionWithAngle(body.Orientation)
}

func (body *Body) Advance(dt float64) {
	angVel := float64(body.AngVelocity)
	angAccel := float64(body.AngAccel)
	body.Orientation += Radian(angVel*dt + 0.5*angAccel*dt*dt)
	body.AngVelocity += Radian(angAccel*dt)
	dir := DirectionWithAngle(body.Orientation)
	body.Position = body.Position.Add(dir.Mul(body.Speed * dt))
}

func (body *Body) Update(t, dt float64) {
	body.Advance(dt)
}