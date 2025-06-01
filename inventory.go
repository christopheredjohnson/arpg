package main

import "github.com/hajimehoshi/ebiten/v2"

type Item struct {
	ID       string
	Name     string
	Type     ItemType
	Icon     *ebiten.Image // Optional: For UI
	Quantity int           // For stackables
}

type Inventory struct {
	Items []*Item
	Size  int // Max number of slots
}

func (inv *Inventory) AddItem(item *Item) bool {
	if len(inv.Items) >= inv.Size {
		return false // inventory full
	}
	inv.Items = append(inv.Items, item)
	return true
}
