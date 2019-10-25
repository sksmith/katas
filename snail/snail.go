package snail

import "fmt"

// Snail returns the path a snail would take
func Snail(area [][]int) (path []int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Snail", r)
		}
	}()
	path = make([]int, len(area)*len(area[0]))
	return lap(area, 0, path)
}

func lap(area [][]int, pathIdx int, path []int) []int {
	if pathIdx == len(path) || len(area) == 0 || len(area[0]) == 0 {
		return path
	}

	for _, v := range area[0] {
		path[pathIdx] = v
		pathIdx++
	}

	lastIdx := len(area[0]) - 1
	for i := 1; i < len(area); i++ {
		path[pathIdx] = area[i][lastIdx]
		pathIdx++
	}

	lastRow := area[len(area)-1]
	for i := len(lastRow) - 2; i >= 0; i-- {
		path[pathIdx] = lastRow[i]
		pathIdx++
	}

	for i := len(area) - 2; i >= 1; i-- {
		path[pathIdx] = area[i][0]
		pathIdx++
	}

	var newarea [][]int
	if len(area)-2 >= 0 && len(area[0])-2 >= 0 {
		newarea = area[1 : len(area)-1]
		for i, row := range newarea {
			newarea[i] = row[1 : len(row)-1]
		}
	}

	return lap(newarea, pathIdx, path)
}
