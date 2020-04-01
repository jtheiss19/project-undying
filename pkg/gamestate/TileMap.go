package gamestate

import (
	"github.com/jtheiss19/project-undying/pkg/tiles"
)

var myTileMap []*tiles.Water

func GetTileMap() []*tiles.Water {
	return myTileMap
}

func NewMap() {
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			myWater := tiles.NewWater(float64(x*64), float64(y*64))
			myTileMap = append(myTileMap, myWater)
		}
	}
}
