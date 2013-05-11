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

type RotMotion struct {
	Orientation Radian
	Position Point
	Speed float64
	AngVelocity Radian
	AngAccel Radian
}

func (motion *RotMotion) Input(pl Placement) {
	motion.Position = pl.Pos
	motion.Orientation = pl.Dir.Angle()
}

func (motion *RotMotion) Output() Placement {
	return Placement{motion.Position, DirectionWithAngle(motion.Orientation)}
}

func (motion *RotMotion) Advance(dt float64) {
	angVel := float64(motion.AngVelocity)
	angAccel := float64(motion.AngAccel)
	motion.Orientation += Radian(angVel*dt + 0.5*angAccel*dt*dt)
	motion.AngVelocity += Radian(angAccel*dt)
	dir := DirectionWithAngle(motion.Orientation)
	motion.Position = motion.Position.Add(dir.Mul(motion.Speed * dt))
}

func (motion *RotMotion) Update(t, dt float64) {
	motion.Advance(dt)
}