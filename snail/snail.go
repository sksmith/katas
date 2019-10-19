package snail

// Snail returns the path a snail would take
func Snail(area [][]int) (path []int) {
	path = make([]int, len(area)*len(area))
	lap(area, path, 0, 0)
	return path
}

func lap(area [][]int, path []int, pathidx, startx, starty int) {
	if pathidx >= len(path) {
		return
	}
	for x := 0; x < len(area[starty]); x++ {

	}
}
