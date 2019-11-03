package main

import "github.com/01-edu/z01"

const CLOSE = false
const OPEN = true

type Door struct {
	state bool
}

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(door *Door) bool {
	PrintStr("is the Door opened ?")
	return door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

func OpenDoor(door *Door) bool {
	PrintStr("Door Opening...")
	door.state = OPEN
	return true
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
