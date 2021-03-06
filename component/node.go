// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package component

import (
	. "geom2d"
)

// Source node in graph of components. Has only output.
type SourceNode struct {
	Dirty bool
	Comp SourceComp
}

func (Source *SourceNode) Update(t, dt float64) {
	// if Source.Dirty {
		Source.Comp.Update(t, dt)
		Source.Dirty = false
	// }
}

func (Source *SourceNode) Output() Placement {
	return Source.Comp.Output()
}

// Sink node in graph of components. Has no output only input.
type SinkNode struct {
	Comp SinkComp
	Source SourceComp	
}

func (sink *SinkNode) Update(t, dt float64) {
	sink.Source.Update(t, dt)
	pl := sink.Source.Output()
	sink.Comp.Input(pl)
	sink.Comp.Update(t, dt)
}

// Filter node in graph of components. Has both input and output.
type FilterNode struct {
	Dirty bool
	Comp FilterComp
	Source SourceComp
}

func (filter *FilterNode) Update(t, dt float64) {
	// if filter.Dirty {
		filter.Source.Update(t, dt)
		pl := filter.Source.Output()
		filter.Comp.Input(pl)
		filter.Comp.Update(t, dt)
		filter.Dirty = false
	// }
}

func (filter *FilterNode) Output() Placement {
	return filter.Comp.Output()
}