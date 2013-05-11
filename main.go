// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
	"os"
	"time"
	//"bensin/graphics"
	. "geom2d"
	. "bensin/base"
	"bensin/graphics/screen"
)

const (
	Title  = "Asteroids"
	Width  = 640
	Height = 640
)

var player *Entity

func mouseButtonPressed(button, state int) {
	x, y := glfw.MousePos()
	p := screen.ToWorld(x, y)
	if button == glfw.Mouse1 && state == glfw.KeyPress {
		player.Visual.SetColor(1, 0, 0, 1)
		if player.Inside(p) {
			player.Visual.SetColor(1, 1, 0, 1)
		}
	}
}

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	if err := glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)

	glfw.SetMouseButtonCallback(mouseButtonPressed)

	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "gl: %s\n", err)
	}

	if err := initScene(); err != nil {
		fmt.Fprintf(os.Stderr, "init: %s\n", err)
		return
	}

	c := time.Tick(10 * time.Millisecond)
	for now := range c {
		render(now)
		glfw.SwapBuffers()
		if glfw.WindowParam(glfw.Opened) != 1 {
			break
		}
	}
	
}

func initScene() (err error) {
	screen.InitScreen(Width, Height, Rect{Point{-20, -20}, Point{20, 20}})
	player = NewPlayer(Point{0, 0})
	return nil
}

func render(t time.Time) {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	player.Update(0, 0.1)
	player.Render(0, 0.01)
}
