// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package graphics

import (
	. "geom2d"
	gl "github.com/chsc/gogl/gl21"
//	"github.com/jteeuwen/glfw"
)

var gView Rect = Rect{Point{-20, -20}, Point{20, 20}}

func updateView() {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(gl.Double(gView.Min.X), 
			 gl.Double(gView.Max.X), 
			 gl.Double(gView.Min.Y), 
			 gl.Double(gView.Max.Y), 
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

func InitSys() {
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

func SetWorldView(r Rect) {
	gView = r
}

func GetWorldView() Rect {
	return gView
}

func LookAt(p Point) {
	gView.MoveCenter(p)
	updateView()
}