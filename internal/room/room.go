package room

import "fmt"

func PrintRoomDetails(roomNumber, size, nights int) {

	text := "night"
	if nights > 1 {
		text = "nights"
	}

	fmt.Println(roomNumber, ":", size, "people /", nights, text)
}
