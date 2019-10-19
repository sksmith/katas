package cubepile

func FindNb(m int) int {
	return addBlock(m, 1, 0)
}

func addBlock(m int, b int, v int) (count int) {
	volume := (b * b * b) + v
	if volume < m {
		return addBlock(m, b+1, volume)
	}
	if volume > m {
		return -1
	}
	return b
}
