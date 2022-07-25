package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mainRoom() {
	const hotelName = "Gopher Paris Inn"
	const totalRooms = 134
	const firstRoomNumber = 110

	rand.Seed(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(totalRooms)
	roomsAvailable := totalRooms - roomsOccupied

	occupancyRate := occupancyRate(roomsOccupied, totalRooms)
	occupancyLevel := occupancyLevel(occupancyRate)

	fmt.Println("Hotel:", hotelName)
	fmt.Println("Number of rooms", totalRooms)
	fmt.Println("Rooms available", roomsAvailable)
	fmt.Println("Occupancy Level:", occupancyLevel)
	fmt.Println("Occupancy Rate: %0.2f %%\n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1
			printRoomDetails(roomNumber, size, nights)
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}
}
