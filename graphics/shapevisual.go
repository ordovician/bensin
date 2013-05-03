// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package graphics

import (
	. "geom2d"
	gl "github.com/chsc/gogl/gl21"
	"unsafe"
//	"github.com/jteeuwen/glfw"
)

type Color [4]gl.Float

type ShapeVisual struct {
	bufferObject gl.Uint
	Color Color
	mode gl.Enum
	noVerticies gl.Sizei
	Direction Direction
	Position Point
}

func (vis *ShapeVisual) Input(pl Placement) {
	vis.Position = pl.Pos
	vis.Direction = pl.Dir
}

func makePolygon(poly Polygon, viz *ShapeVisual) {
	count := gl.Sizeiptr(len(poly))
	size  := gl.Sizeiptr(unsafe.Sizeof(poly[0]))
	verticies := make([]gl.Float, len(poly) * 2)
	for i, p := range poly {
		verticies[i*2] = gl.Float(p.X)
		verticies[i*2 + 1] = gl.Float(p.Y)
	}
	viz.noVerticies = gl.Sizei(len(poly))
	viz.mode = gl.POLYGON
	gl.BufferData(	gl.ARRAY_BUFFER, 
					count * size,
					gl.Pointer(&verticies[0]), 
					gl.STATIC_DRAW)
}

func NewShapeVisual(aShape Shape) *ShapeVisual {
	var viz ShapeVisual
	gl.GenBuffers(1, &viz.bufferObject)
	gl.BindBuffer(gl.ARRAY_BUFFER, viz.bufferObject)
	defer gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	switch  shape := aShape.(type) {
	case Polygon:
		makePolygon(shape, &viz)
	} 
	return &viz
}

func (vis *ShapeVisual) Render(t, dt float64) {
	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.BindBuffer(gl.ARRAY_BUFFER, vis.bufferObject)
	gl.VertexPointer(2, gl.FLOAT, 0, gl.Pointer(nil))

	var matrix [16]gl.Float
	matrix[10], matrix[15] = 1, 1
	
	gl.PushMatrix()
		gl.Color4fv(&vis.Color[0])
		
		// copy our orientation information into matrix, so
		// OpenGL will position and rotate our shape correctly
		dirX, dirY := gl.Float(vis.Direction.X), gl.Float(vis.Direction.Y)
		matrix[0] = dirX
		matrix[1] = dirY
		matrix[4] = -dirY
		matrix[5] = dirX
		
		matrix[12] = gl.Float(vis.Position.X)
		matrix[13] = gl.Float(vis.Position.Y)
		
		gl.LoadMatrixf(&matrix[0])
		gl.DrawArrays(vis.mode, 0, vis.noVerticies)
	
	gl.PopMatrix()
}

