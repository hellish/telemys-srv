package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"

	"github.com/go-vgo/robotgo"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

type app struct {
	x  float32
	y  float32
	xv float32
	yv float32
}

func readF32(data []byte) (ret float32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func startServer(ch chan app) {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
	checkError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	checkError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	fmt.Printf("server started...\n")

	for {

		_, _, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		//fmt.Println("Received ", string(buf[0:n]), " from ", addr, len(buf), n)
		//fmt.Println("Received ", string(buf[0:4]), " from ", addr, len(buf), n)
		//fmt.Println("Received ", string(buf[4:8]), " from ", addr, len(buf), n)
		//fmt.Println("Received ", string(buf[8:12]), " from ", addr, len(buf), n)
		//fmt.Println("Received ", string(buf[12:16]), " from ", addr, len(buf), n)

		var aa app
		aa.x = readF32(buf[0:4])
		aa.y = readF32(buf[4:8])
		aa.xv = readF32(buf[8:12])
		aa.yv = readF32(buf[12:16])

		//fmt.Printf("hex for buffer %v\n", hex.EncodeToString(buf))
		//fmt.Printf("hex for buffer %v - %f \n", hex.EncodeToString(buf[0:4]), readF32(buf[0:4]))
		//fmt.Printf("hex for buffer %v\n", hex.EncodeToString(buf[4:8]))
		//fmt.Printf("hex for buffer %v\n", hex.EncodeToString(buf[8:12]))
		//fmt.Printf("hex for buffer %v\n", hex.EncodeToString(buf[12:16]))
		//fmt.Printf("hex for buffer %v\n", aa)

		ch <- aa
	}
}

// MUST WATCH https://www.youtube.com/watch?v=vqvJn4G3NMM
func main() {
	x, y := robotgo.GetMousePos()
	fmt.Printf("current mouse position %d,%d\n", x, y)
	// robotgo.MoveMouse(99, 99)

	w, h := robotgo.GetScreenSize()
	fmt.Printf("screen dimensions w=%d h=%d\n", w, h)

	cc := make(chan app)

	go startServer(cc)

	for ccc := range cc {
		x, y = robotgo.GetMousePos()

		fmt.Printf("msg arrived %v\n", ccc)

		x += int(ccc.x)
		y += int(ccc.y)

		robotgo.MoveMouse(x, y)
	}
}
