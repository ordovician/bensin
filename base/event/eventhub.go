// Copyright 2012 Erik Engheim. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

/*
A listener can register interest in an event, and be notified by event hub when a sender
post this event. Listeners can decide whether they care about who sends the event.

	func NewMyObject() *MyObject {
		obj := new(MyObject)
		event.Register(event.MouseButton, nil, obj)
	}

To be able to receive events your type needs to implement the Receive() method:

	func (obj *MyObject) Receive(e Event) {
		switch data := e.Data(type) {
		case MouseEvent:
			if data.State == event.ButtonDown {
				fmt.Printf("Button pressed at (%f, %f)", data.Pos.X, data.Pos.Y)
			}
		}
	}

Objects should unregister from event hub before they are destroyed or when the notifications
are no longer needed.

	func (obj *MyObject) Cleanup() {
		event.Unregister(even.MouseButton, nil, obj)
	}
*/
package event

import (

)

func init() {
	for i,_ := range receivers {
		ignoreSenders	:= make(ListenerSet)
		senders 		:= make(map[interface{}]ListenerSet)
		receivers[i].ignoreSenders = ignoreSenders
		receivers[i].senders = senders
	}
}

type Listener interface {
	Receive(event Event)
}

type ListenerSet map[Listener]bool

type receiver struct {
	ignoreSenders ListenerSet
	senders map[interface{}]ListenerSet
}

var receivers [NoEvents]receiver

func Register(etype EventType, sender interface{}, receiver Listener) {
	r := &receivers[etype]
	if sender == nil {
		r.ignoreSenders[receiver] = true
	} else {
		r.senders[sender][receiver] = true
	}
}

func Unregister(etype EventType, sender interface{}, receiver Listener) {
	if sender == nil {
		delete(receivers[etype].ignoreSenders, receiver)			
	} else {
		delete(receivers[etype].senders[sender], receiver)		
	}
}

func Notify(event Event) {
	r := &receivers[event.Type]
		
	for listener,_ := range r.ignoreSenders {
		listener.Receive(event)
	}
	
	if event.Sender != nil {
		listeners := r.senders[event.Sender]
		for listener,_ := range listeners {
			listener.Receive(event)
		}
	}
}