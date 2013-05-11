// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package component

import (
	. "geom2d"
)

type Collider struct {
	world Matrix3x3
	OrigShape, transShape Polygon
	BoundingBox Rect	
}

func NewCollider(pos Point, dir Direction, poly Polygon) *Collider {
	coll := new(Collider)
	coll.Init(pos, dir, poly)
	return coll
}

func (coll *Collider) Init(pos Point, dir Direction, poly Polygon) {
	coll.world = Identity()
	coll.world.SetPos(pos)
	coll.world.SetDir(dir)
	coll.OrigShape = make(Polygon, len(poly))
	coll.transShape = make(Polygon, len(poly))
	copy(coll.OrigShape, poly)
	coll.UpdateShapePlacement()
}

func (coll *Collider) Input(pl Placement) {
	coll.world.SetPos(pl.Pos)
	coll.world.SetDir(pl.Dir)
}

func (coll *Collider) UpdateShapePlacement() {
	copy(coll.transShape, coll.OrigShape)
	coll.transShape.Transform(coll.world)
	coll.BoundingBox = coll.transShape.BoundingBox()
}

func (coll *Collider) Update(t, dt float64) {
	coll.UpdateShapePlacement()
}

func (coll *Collider) Inside(q Point) bool {
	return coll.transShape.Inside(q)
}

func (coll *Collider) Intersect(shape Shape) bool {
	return coll.transShape.Intersect(shape)
}