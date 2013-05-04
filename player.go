// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package main

import (
	. "geom2d"
	. "bensin/base"
)


func NewPlayer(pos Point) *Entity {
	poly := Polygon{Point{-2, -1},
					Point{2, 0},
					Point{-2, 1}}
	player := NewMovableEntity(pos, poly)
	
	return player
}

