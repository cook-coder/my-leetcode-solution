package robotReturnToOrigin

func judgeCircle(moves string) bool {

	var x, y = 0, 0
	for _, move := range moves {
		switch move {
		case 68:
			// move down
			y--
		case 76:
			// move left
			x--
		case 82:
			// move right
			x++
		default:
			// case 85:
			// move up
			y++
		}
	}
	return x == 0 && y == 0
}
