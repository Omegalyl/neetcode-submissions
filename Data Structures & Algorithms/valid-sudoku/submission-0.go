func isValidSudoku(board [][]byte) bool {
	seenRow := map[byte]struct{}{}
	seenCol := [9]map[byte]struct{}{}
	for i := 0; i < 9; i++ {
		seenCol[i] = make(map[byte]struct{})
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}

			if _, has := seenRow[v]; has {
				return false
			}

			if _, has := seenCol[j][v]; has {
				return false
			}
			seenCol[j][v] = struct{}{}
			seenRow[v] = struct{}{}
		}
		seenRow = map[byte]struct{}{}
	}

	seenMat := map[byte]struct{}{}
	for z := 0; z < 9; z++ {
		r := (z/3) * 3
		c := (z%3) * 3
		for i := 0; i < 3; i++ {
			for j:= 0; j < 3; j++ {
				v := board[r+i][c+j]
				if v == '.' {
					continue
				}

				if _, has := seenMat[v]; has {
					return false
				}
				seenMat[v] = struct{}{}
			}
		}
		seenMat = map[byte]struct{}{}
	}
	return true
}
