package playerControl

import (
	"net"

	"github.com/hajimehoshi/ebiten"
	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/gamestate"
	"github.com/jtheiss19/project-undying/pkg/networking/connection"
)

//Replicator is the component that handles all
//replication of an element onto a server.
type Shooter struct {
	container *elements.Element
	Type      string
	HasFired  bool
	DestX     float64
	DestY     float64
}

func init() {
	var shoot = new(Shooter)
	gamestate.MRPMAP["Shooter"] = shoot
}

//NewReplicator creates a Replicator which is
//the component that handles all replication
//of an element onto a server.
func NewShooter(container *elements.Element) *Shooter {
	return &Shooter{
		container: container,
		Type:      "Shooter",
	}
}

func (shoot *Shooter) MRP(finalElem *elements.Element, conn net.Conn) {
	myComp := NewShooter(finalElem)
	finalElem.AddComponent(myComp)
}

//OnDraw is used to qualify SpriteRenderer as a component
func (shoot *Shooter) OnDraw(screen *ebiten.Image, xOffset float64, yOffset float64) error {
	if shoot.container.ID != connection.GetID() && shoot.HasFired == true {
		return nil
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		shoot.HasFired = true
		w, h := ebiten.CursorPosition()
		shoot.DestX = float64(w) - xOffset
		shoot.DestY = float64(h) - yOffset
	} else {
		shoot.HasFired = false
	}
	return nil
}

//OnUpdate sends the state of the current element to the
//connection if it exists. On servers to not init elements
//with a connection. On clients init the objects with a
//connection.
func (shoot *Shooter) OnUpdate() error {
	return nil
}

func (shoot *Shooter) OnCheck(elemC *elements.Element) error {
	return nil
}

func (shoot *Shooter) OnMerge(compM elements.Component) error {
	return nil
}

func (shoot *Shooter) OnUpdateServer() error {
	if shoot.DestX != 0 {
		//fmt.Println("Fire Shot toward X:", shoot.DestX, "Y:", shoot.DestY)
	}

	return nil
}
