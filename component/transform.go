// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package component

import . "geom2d"

// Position and orientation of a game object.
// Transforms form a scene graph, where each transform defines
// a local coordinate system using a matrix.
type Transform struct {
	Local, World Matrix3x3
	Parent int
}

func (trans *Transform) Input(pos Point, dir Direction) {
	trans.Local.SetPos(pos)
	trans.Local.SetDir(dir)
}

func (trans *Transform) Output() (pos Point, dir Direction) {
	return trans.World.Pos(), trans.World.Dir()
}

func (trans *Transform) Update(t, dt float64) {
	trans.World = trans.Local
}

// Holds the Transform components for every game object in scene
type SceneMgr struct {
	Comps []Transform
}

// Calculate the world transform for every game object.
// The world transform is a matrix indicating position and
// orientation of a game object in our world.
func (mgr *SceneMgr) update(t, dt float64) {
	var nodes []Transform = mgr.Comps
	
	for _, n := range nodes {
		if n.Parent == -1 {
			// for root nodes Local is the same as world transform
			n.World = n.Local
		} else {
			// assume parents have gotten their world transform
			// calculated first
			n.World = nodes[n.Parent].World.Mul(n.Local) 
		}
	}
}