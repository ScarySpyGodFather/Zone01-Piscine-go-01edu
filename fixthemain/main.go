package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const (
	OPEN  = true
	CLOSE = false
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	ptrDoor.state = CLOSE
	return ptrDoor.state
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?\n")
	if Door.state == OPEN {
		return true
	}
	return false
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	if ptrDoor.state == CLOSE {
		return true
	}
	return false
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
