# Bensin 2D Game Engine
*Bensin* is Norwegian for gasoline. Bensin is a game engine written in Go. At the moment it is not usable since it is at such an early stage.

Bensin is not being made to create the greatest 2D game engine every. The main motivation for the project is to explore both *Object Oriented Design*, *Data Oriented Design* and *Component Based* game engine architecture. A second goal is to explore how suitable the Go programming languge is for game development. The project will aim to use Go both for creating the engine itself and for scripting. Given the quick compile times of Go, duck typing and simple syntax I think Go might also be suitable for scripting.

# Overview
I have factored out the most generic code such as the `geom2d` package into separate git repositories since it can be easily used in other projects not related to *Bensin*. The packages `component` and `graphics` are quite specific to the game engine.

## Entity
Every game object you see on the screen is an entity. E.g. in a space invader game, both the alien invasion and the spaceship. Even the bullets and bombs are entities. Entities have a *position* and *orientation* on screen. They need not nessesarily be visible on screen, but they will typically be visible and move.

## Component
An enity is really just an emtpy shell. It does not do much by itself. It gets its abilities from components. Each component gives a different ability. A *Collider* component allows the component to register collision or intersection with other entities. A *Motion* component gives the entity velocity and lets it move across the screen. A *Visual* component will make it draw itself on the screen. Without a *Visual* component it would be invisible.