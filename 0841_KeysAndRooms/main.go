package main

func main() {
	canVisitAllRooms([][]int{{1}, {2}, {3}, {}})
}

func canVisitAllRooms(rooms [][]int) bool {
	visitedRooms := make(map[int]struct{}, len(rooms))
	currKeys := make([]int, 0, len(rooms)) // Stack
	currKeys = append(currKeys, rooms[0]...)
	visitedRooms[0] = struct{}{}

	for len(currKeys) > 0 {
		lastKeyIdx := len(currKeys) - 1
		key := currKeys[lastKeyIdx]
		if _, visited := visitedRooms[key]; visited {
			currKeys = currKeys[:lastKeyIdx] // Already visited this room
			continue
		}

		currKeys = currKeys[:lastKeyIdx] // Already visited this room
		currKeys = append(currKeys, rooms[key]...)
		visitedRooms[key] = struct{}{}

		if len(visitedRooms) == len(rooms) {
			return true
		}
	}

	return false
}
