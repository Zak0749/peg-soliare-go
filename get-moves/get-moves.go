package moves

func moveimpossible(idx int) bool {
	impossible := [...]int{0, 2, 9, 17, 27, 35, 42, 44}
	for _, num := range impossible {
		if num == idx {
			return true
		}
	}
	return false
}

func horizonalpossible(idx int) bool {
	impossible := [...]int{3, 4, 6, 8, 18, 26, 36, 38, 39, 41}
	for _, num := range impossible {
		if num == idx {
			return false
		}
	}
	return true
}

func verticalpossible(idx int) bool {
	impossible := [...]int{1, 10, 11, 15, 16, 28, 29, 33, 34, 43}
	for _, num := range impossible {
		if num == idx {
			return false
		}
	}
	return true
}

func getneigbors(idx int) (int, int, int, int) {
	l := idx - 1
	r := idx + 1
	var up int
	var down int

	if idx > 17 && idx < 27 {
		up = idx - 9
		down = idx + 9
	} else if idx < 6 || idx > 38 {
		up = idx - 3
		down = idx + 3
	} else if idx == 6 || idx == 7 || idx == 8 {
		up = idx - 3
		down = idx + 6
	} else if idx == 12 || idx == 13 || idx == 14 {
		up = idx - 6
		down = idx + 9
	} else if idx == 30 || idx == 31 || idx == 32 {
		up = idx - 9
		down = idx + 6
	} else {
		up = idx - 6
		down = idx - 3
	}

	return l, r, up, down
}

func Get(board [45]bool) [][45]bool {
	moves := [][45]bool{}
	for idx, val := range board {
		if !val || moveimpossible(idx) {
			continue
		}

		ridx, lidx, uidx, didx := getneigbors(idx)

		if horizonalpossible(idx) && board[ridx] != board[lidx] {
			copy := board
			copy[idx] = false
			copy[ridx] = !board[ridx]
			copy[lidx] = !board[lidx]
			moves = append(moves, copy)
		}

		if verticalpossible(idx) && board[uidx] != board[didx] {
			copy := board
			copy[idx] = false
			copy[uidx] = !board[uidx]
			copy[didx] = !board[didx]
			moves = append(moves, copy)
		}
	}
	return moves
}
