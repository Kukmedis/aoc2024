package main

func isXmas(letters [][]rune, i int, j int, iinc int, jinc int) bool {
	return letters[i][j] == 'X' && isLetterRight(letters, 'M', i, j, iinc, jinc) && isLetterRight(letters, 'A', i, j, iinc*2, jinc*2) && isLetterRight(letters, 'S', i, j, iinc*3, jinc*3)
}

func isLetterRight(letters [][]rune, xmasChar rune, i int, j int, iinc int, jinc int) bool {
	ilen := len(letters)
	jlen := len(letters[i])
	return i+iinc >= 0 && i+iinc < ilen && j+jinc >= 0 && j+jinc < jlen && letters[i+iinc][j+jinc] == xmasChar
}

func calculateXmasInAnyDirection(letters [][]rune, i int, j int) int {
	sum := 0
	if isXmas(letters, i, j, 0, -1) {
		sum++
	}
	if isXmas(letters, i, j, 0, 1) {
		sum++
	}
	if isXmas(letters, i, j, -1, -1) {
		sum++
	}
	if isXmas(letters, i, j, -1, 0) {
		sum++
	}
	if isXmas(letters, i, j, -1, 1) {
		sum++
	}
	if isXmas(letters, i, j, 1, -1) {
		sum++
	}
	if isXmas(letters, i, j, 1, 0) {
		sum++
	}
	if isXmas(letters, i, j, 1, 1) {
		sum++
	}
	return sum
}

func calculateAllXmases(letters [][]rune) int {
	sum := 0
	for i := range letters {
		for j := range letters {
			sum = sum + calculateXmasInAnyDirection(letters, i, j)
		}
	}
	return sum
}

func isXmasCross(letters [][]rune, i int, j int) bool {
	if letters[i][j] == 'A' {
		if (isLetterRight(letters, 'M', i, j, 1, 1) && isLetterRight(letters, 'S', i, j, -1, -1)) || (isLetterRight(letters, 'S', i, j, 1, 1) && isLetterRight(letters, 'M', i, j, -1, -1)) {
			if (isLetterRight(letters, 'M', i, j, -1, 1) && isLetterRight(letters, 'S', i, j, 1, -1)) || (isLetterRight(letters, 'S', i, j, -1, 1) && isLetterRight(letters, 'M', i, j, 1, -1)) {
				return true
			}
		}
	}
	return false
}

func calculateXmasCrosses(letters [][]rune) int {
	sum := 0
	for i := range letters {
		for j := range letters {
			if isXmasCross(letters, i, j) {
				sum++
			}
		}
	}
	return sum
}
