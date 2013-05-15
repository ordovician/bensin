// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

// OpenGL related code. Does not contain code related to the windowing system, so 
// there is a free choice in using glfw, glu or SDL. 
package graphics

import (
	. "geom2d"
	gl "github.com/chsc/gogl/gl21"
//	"github.com/jteeuwen/glfw"
)

// An OpenGL 4x4 matrix. Used to transform points in scene. 
type Matrix4x4 [16]gl.Float

// Set position and orientation an OpenGL 4x4 matrix
func (matrix *Matrix4x4) Set(pos Point, dir Direction) {
	matrix[0] = gl.Float(dir.X)
	matrix[1] = gl.Float(dir.Y)
	matrix[4] = gl.Float(-dir.Y)
	matrix[5] = gl.Float(dir.X)
	
	matrix[12] = gl.Float(pos.X)
	matrix[13] = gl.Float(pos.Y)
}
