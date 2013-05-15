// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package event

import (
	. "geom2d"
)

// Two events might have the same data structure but be of different type.
type EventType int

const (
	Undefined EventType = iota
	Collision
	MouseButtonStateChange
	MouseDrag
	MouseHitObject
	NoEvents
)

// Indicates which button has been pressed or released.
type MouseButton int

const (
	ButtonLeft MouseButton = iota
	ButtonRight
)

// Indicates whether button is being pressed or released
type ButtonState int

const (
	ButtonDown ButtonState = iota
	ButtonUp
)

// Events are sent when something happens which one want to notify other parts of the system.
// Events need to contain their sender, because they are often stored in queues.
type Event struct {
	Type EventType
	Sender, Data interface{}
}

// Data structure for a mouse related event.
type MouseEvent struct {
	Pos Point
	Button MouseButton
	State  ButtonState
}
