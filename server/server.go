package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	"gtihub.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	termbox.Clear(termbox.ClororBlack, termbox.ClororBlack)
	adr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10234")
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.ListenUDP("udp", adr)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		handleConnection(listener)
	}

}

func handleConnection(con *net.UDPConn) {
	buf := make([]byte, 2000)
	n, err := con.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	buff := bytes.NewReader(buf[0:n])

	var data struct {
		X int32
		Y int32
	}
	err = binary.Read(buff, binary.LittleEndian, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	termbox.SetCell(data.X, data.Y, '&', termbox.ClororWhite, termbox.ClororBlack)
	termbox.Flush()
	termbox.Close()
	//fmt.Println(data)
}
