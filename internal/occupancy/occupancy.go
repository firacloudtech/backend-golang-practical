package occupancy

const highLimit = 70.0
const mediumLimit = 20.0

func level(occupancyRate float32) string {
	switch {
	case occupancyRate > highLimit:
		return "High"
	case occupancyRate > mediumLimit:
		return "Medium"
	default:
		return "Low"
	}
}

func rate(roomsOccupied int, totalRooms int) float32 {

	return (float32(roomsOccupied) / float32(totalRooms)) * 100
}
