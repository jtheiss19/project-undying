package units

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var masterTexture *ebiten.Image

type Destroyer struct {
	Unit
}

func init() {
	origEbitenImage, _, err := ebitenutil.NewImageFromFile("../../assets/sprites/destroyer.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	w, h := origEbitenImage.Size()
	masterTexture, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)

	op := &ebiten.DrawImageOptions{}

	masterTexture.DrawImage(origEbitenImage, op)
}

func NewDestroyer(unitInfo *Unit, owner string) *Destroyer {
	myDestroyer := &Destroyer{}
	if unitInfo == nil {
		myDestroyer.Unit = *NewUnit("destroyer", owner)
	} else {
		myDestroyer.Unit = *unitInfo
	}

	return myDestroyer
}

func (s *Destroyer) Draw(screen *ebiten.Image, xOffset float64, yOffset float64) {
	w, h := masterTexture.Size()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(1 * math.Pi * 270 / 360)
	op.GeoM.Translate(float64(w)/2, float64(h)/2)
	op.GeoM.Translate(s.X, s.Y)
	op.GeoM.Translate(xOffset, yOffset)

	screen.DrawImage(masterTexture, op)
}
