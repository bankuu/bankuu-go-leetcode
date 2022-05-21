package main

func criticalConnections(n int, connections [][]int) [][]int {

	mapper := make(map[int][]int, 0)

	// Creation Mapper
	for i := 0; i < n; i++ {
		for _, connection := range connections {
			var link []int

			if val, ok := mapper[i]; ok {
				link = val
			} else {
				link = make([]int, 0)
			}

			if i == connection[0] {
				link = append(link, connection[1])
			}
			if i == connection[1] {
				link = append(link, connection[0])
			}
			mapper[i] = link
		}
	}

	// Critical Route
	for a := 0; a < n; a++ {
		for b := a + 1; b < n; b++ {

		}
	}

	return nil
}

func main() {
	result := criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}})
	print(result)
}
