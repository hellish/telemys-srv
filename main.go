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

// Command represents comments send by the app
type Command struct {
	Type    byte
	Address *net.UDPAddr
	X       int
	Y       int
	Raw     []byte
}

func readF32(data []byte) (ret float32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func startServer(commands chan Command) {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
	checkError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	checkError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	fmt.Printf("server started...\n")

	for {
		num, address, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("error: ", err)
			continue
		}

		if num == 0 {
			fmt.Println("read bytes num 0")
			continue
		}

		cmd := Command{Address: address, Raw: buf}

		switch code := buf[0]; code {
		case MessageActionMove:
			// case MessageActionTap:
			// case MessageActionDblTab:
			fmt.Println("mouse movement action found")
			cmd.Type = code
			cmd.X = int(readF32(buf[1:5]))
			cmd.Y = int(readF32(buf[5:9]))
			break
		case MessageActionTap:
			fmt.Println("tap action found")
			cmd.Type = code
			break
		case MessageActionDblTab:
			fmt.Println("double tap action found")
			cmd.Type = code
			break
		default:
			fmt.Printf("unsupported message type %v\n", code)
			continue
		}

		commands <- cmd
	}
}

// MUST WATCH https://www.youtube.com/watch?v=vqvJn4G3NMM
func main() {
	x, y := robotgo.GetMousePos()
	fmt.Printf("current mouse position %d,%d\n", x, y)
	// robotgo.MoveMouse(99, 99)

	w, h := robotgo.GetScreenSize()
	fmt.Printf("screen dimensions w=%d h=%d\n", w, h)

	commands := make(chan Command)

	go startServer(commands)

	for command := range commands {
		fmt.Printf("handling command %v - %d : %d\n", command.Type, command.X, command.Y)
		switch command.Type {
		case MessageActionMove:
			x, y = robotgo.GetMousePos()
			x += command.X
			y += command.Y
			robotgo.MoveMouse(x, y)
			break
		case MessageActionTap:
			robotgo.Click("left", false)
			break
		case MessageActionDblTab:
			robotgo.Click("left", true)
		}
	}
}
