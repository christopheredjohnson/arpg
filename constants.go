package main

const (
	screenWidth  = 640
	screenHeight = 480
)

const (
	dirUp = iota
	dirDown
	dirLeft
	dirRight
)

const (
	stateIdle = iota
	stateWalk
	stateAttack
)

type ItemType int

const (
	ItemWeapon ItemType = iota
	ItemArmor
	ItemConsumable
)
