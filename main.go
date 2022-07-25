package main

import (
	"fmt"

	"github.com/yogenp/practical-go-lessons/internal/room"
)

func init() {
	fmt.Println("launch init")
}
func main() {
	room.PrintRoomDetails(112, 3, 2)
}
