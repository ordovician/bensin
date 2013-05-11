// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package screen

import (
	. "geom2d"
	gl "github.com/chsc/gogl/gl21"
//	"github.com/jteeuwen/glfw"
)

// The screen is what we draw on. It has a width and height. When you 
// draw you specify positions in world coordintes. But external events
// will be relative to screen width and height.
type Screen struct {
	width, height int	// width and height of screen seen externally. 
	world Rect			// Local logical coordinates
}

var screen Screen

// Map value in range (min1, max1) to a value in range (min2, max2)
func mapValue(value, min1, max1, min2, max2 float64) float64 {
	return value * (max2 - min2) / (max1 - min1) + min2 
}

// Convert screen coordinates to world coordinates.
// This is usefull for converting mouse coordinates to world coordinates
func ToWorld(x, y int) Point {
	y = screen.height - y
	w := screen.world
	
	var pos Point
	pos.X = float64(x)
	pos.Y = float64(y)
	pos.X = mapValue(pos.X, 0, float64(screen.width) , w.Min.X, w.Max.X)
	pos.Y = mapValue(pos.Y, 0, float64(screen.height), w.Min.Y, w.Max.Y)
	
	return pos
}

func updateView() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	r := screen.world
	gl.Ortho(gl.Double(r.Min.X), 
			 gl.Double(r.Max.X), 
			 gl.Double(r.Min.Y), 
			 gl.Double(r.Max.Y), 
			 -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
}

func initGL() {
	gl.Enable(gl.TEXTURE_2D)
	gl.ShadeModel(gl.FLAT)
	
	// allows us to have partly transparent edges on textures and use
	// them to create anti aliasing
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.BLEND)
	
	gl.Disable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 0)
	
	updateView()
}

func InitScreen(width, height int, world Rect) {
	SetWorld(world)
	SetSize(width, height)
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
	
	initGL()
	
	// sprites.InitMgr(100)
	// visuals.InitMgr(100)
}

func Update(t, dt float64) {
	// sprites.Update(t, dt)
	// visuals.Update(t, dt)
}

func Render(t, dt float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	
	// sprites.Render(t, dt)
	// visuals.Render(t, dt)
}

func SetSize(width, height int) {
	screen.width = width
	screen.height = height
}

func Size() (width, height int) {
	return screen.width, screen.height
}

func SetWorld(r Rect) {
	screen.world = r
}

func World() Rect {
	return screen.world
}

func LookAt(p Point) {
	screen.world.MoveCenter(p)
	updateView()
}