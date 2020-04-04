package gamestate

import (
	"encoding/json"
	"net"

	"github.com/jtheiss19/project-undying/pkg/elements"
	"github.com/jtheiss19/project-undying/pkg/networking/mrp"
)

var connectionList = make(map[int]net.Conn)

func SendElem(conn net.Conn, elem *elements.Element) {
	bytes, _ := json.Marshal(&elem)

	myMRP := mrp.NewMRP([]byte("ELEM"), bytes, []byte(""))
	conn.Write(myMRP.MRPToByte())
}

func ForceUpdate(conn net.Conn) {
	myMRP := mrp.NewMRP([]byte("END"), []byte(""), []byte(""))
	conn.Write(myMRP.MRPToByte())
}

func NewConnection(conn net.Conn, ID int) {
	connectionList[ID] = conn
}

func UpdateElemToAll(elem *elements.Element) {
	for _, client := range connectionList {
		SendElem(client, elem)
	}
}
