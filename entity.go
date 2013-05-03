package main

import (
	. "geom2d"
)

type Entity interface {
	Parent() Entity
	Update(t, dt float64)
	Local() Matrix3x3
	World() Matrix3x3
}

type LeafEntity struct {
	parent Entity
	local, world Matrix3x3
}

func (entity *LeafEntity) Parent() Entity {
	return entity.parent
}

func (entity *LeafEntity) World() Matrix3x3 {
	return entity.world
}

func (entity *LeafEntity) Local() Matrix3x3 {
	return entity.local
}

func (entity *LeafEntity) UpdateWorld() {
	entity.world = entity.parent.World().Mul(entity.local)
}

func (entity *LeafEntity) Update(t, dt float64) {
	
}

type ParentEntity struct {
	LeafEntity
	children []Entity
}

func (entity *ParentEntity) Update(t, dt float64) {
	entity.LeafEntity.UpdateWorld()
	for _, child := range entity.children {
		child.Update(t, dt)
	}
}

