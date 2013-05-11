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

type DirMotion struct {
	velocity, acceleration Vector2D
	pos Point
}

func (motion *DirMotion) Input(pl Placement) {
	motion.pos = pl.Pos
}

func (motion *DirMotion) Output() Placement {
	return Placement{motion.pos, motion.acceleration.Unit()}
}

func (motion *DirMotion) Advance(dt float64) {
	vx, vy := motion.velocity.X, motion.velocity.Y
	ax, ay := motion.acceleration.X, motion.acceleration.Y
	
	motion.pos.X += vx*dt + 0.5*ax*dt*dt
	motion.pos.Y += vy*dt + 0.5*ay*dt*dt
}

func (motion *DirMotion) Update(t, dt float64) {
	motion.Advance(dt)
}