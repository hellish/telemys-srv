package main

import (
	"bytes"
	"encoding/binary"
)

// MessageActionNoAction does nothing
const MessageActionNoAction byte = 0x0

// MessageActionConnect used to indicate when user connects
const MessageActionConnect byte = 0x1

// MessageActionMove used to indicate when user moves
const MessageActionMove byte = 0x2

// MessageActionTap used to indicate when user taps
const MessageActionTap byte = 0x3

// MessageActionDblTab used to indicate when user double taps
const MessageActionDblTab byte = 0x4

// MessageActionSwipeLeftToRight used to indicate when user swipes left to right
const MessageActionSwipeLeftToRight byte = 0x5

// MessageActionSwipeRightToLeft used to indicate when user swipes right to left
const MessageActionSwipeRightToLeft byte = 0x6

// MessageActionSwipeUpToDown used to indicate when user swipes up to down
const MessageActionSwipeUpToDown byte = 0x7

// MessageActionSwipeDownToUp used to indicate when user swipes down to up
const MessageActionSwipeDownToUp byte = 0x8

func readF32(data []byte) (ret float32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}

func readI(data []byte) (ret int) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &ret)
	return
}
