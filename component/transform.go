// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package component

import . "geom2d"

// Position and orientation of a game object.
// Transforms form a scene graph, where each transform defines
// a local coordinate system using a matrix.
type Transform struct {
	parent, Local, World Matrix3x3
}

// Create a new transform
func NewTransform(pos Point) Transform {
	var trans Transform
	trans.parent = Identity()
	trans.Local = Translate(Vector2D(pos))
	trans.World = trans.Local
	return trans
}

// Set the world matrix of the parent transform
// Make sure this is called before Output() and Update() is called.
func (trans *Transform) SetParent(parent Matrix3x3) {
	trans.parent = parent
}

// Input a placement in local coordinate system
func (trans *Transform) Input(pl Placement) {
	trans.Local.SetPos(pl.Pos)
	trans.Local.SetDir(pl.Dir)
}

//  Output is given in world coordinate system
func (trans *Transform) Output() Placement {
	return Placement{trans.World.Pos(), trans.World.Dir()}
}

// Update world matrix based on local matrix and world matrix of parent.
// Before calling this you must call trans.SetParent()
func (trans *Transform) Update(t, dt float64) {
	trans.World = trans.parent.Mul(trans.Local)
}