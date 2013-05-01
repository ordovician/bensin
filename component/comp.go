// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

// An Entity (game object) is a container for Components
// each component gives the Entity an ability. This package
// contains a set of standard components.
package component

import (
	. "geom2d"
)

// Minimum interface to a component. 
type Component interface {
	Update(t, dt float64)	
}

// Produce position an direction. Physics components e.g.
type SourceComp interface {
	Component
	Output() (pos Point, dir Direction)
}

// Consume position and direction. Colliders and Visuals e.g.
type SinkComp interface {
	Component
	Input(pos Point, dir Direction)
}

// Transforms input to an output
// E.g. a Transform component will take in local coordinates 
// and produce world coordinates.
type FilterComp interface {
	Component
	Input(pos Point, dir Direction)
	Output() (pos Point, dir Direction)
}
