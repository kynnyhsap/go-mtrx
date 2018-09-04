package main

// Matrix -
type Matrix [][]int

func (m Matrix) isValid() bool {
	if len(m) == 0 {
		return true
	}

	firstRow := m[0]
	for _, row := range m {
		if len(firstRow) != len(row) {
			return false
		}
	}

	return true
}

func (m Matrix) isSquareMatrix() bool {
	if !m.isValid() {
		return false
	}

	return len(m) == len(m[0])
}

func (m Matrix) countRows() int {
	if !m.isValid() {
		return 0
	}

	return len(m)
}

func (m Matrix) countColumns() int {
	if !m.isValid() {
		return 0
	}

	if len(m) > 0 {
		firstRow := m[0]

		return len(firstRow)
	}

	return 0
}

func (m Matrix) isNullMatrix() bool {
	if !m.isValid() {
		return false
	}

	for _, row := range m {
		for _, number := range row {
			if number != 0 {
				return false
			}
		}
	}

	return true
}

func main() {

}
