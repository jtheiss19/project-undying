package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/gameloop"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	screenScale  = 1
)

func main() {

	//gamestate.Dial("localhost:8080")
	//gamestate.NewWorld()

	if err := ebiten.Run(gameloop.Update, screenWidth/screenScale, screenHeight/screenScale, screenScale, "test"); err != nil {
		log.Fatal(err)
	}
}
