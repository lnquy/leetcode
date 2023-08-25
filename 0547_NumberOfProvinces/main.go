package main

// Idea:
// 1. Sequentially iterate through the list of cities (isConnected).
// 2. If we haven't visited this city isConnected[i] before then it must belong to a new group/province.
// 3. Iterate through the connections of isConnected[i] city and
//    recursively mark all other cities of this group/province as visited.
//
// Each province/group is a graph, and the traceConnections func is a depth-first search (DFS) to find all
// the cities/nodes belong to that graph.

func findCircleNum(isConnected [][]int) (provinces int) {
	visitedCities := make([]bool, len(isConnected))

	for i := range visitedCities {
		if visitedCities[i] {
			continue
		}
		provinces++
		visitedCities = traceConnections(isConnected, i, visitedCities)
	}

	return provinces
}

func traceConnections(allConns [][]int, currCity int, visitedCities []bool) []bool {
	visitedCities[currCity] = true
	for i, conn := range allConns[currCity] {
		if conn == 1 && !visitedCities[i] {
			visitedCities = traceConnections(allConns, i, visitedCities)
		}
	}
	return visitedCities
}
