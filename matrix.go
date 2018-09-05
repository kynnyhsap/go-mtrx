package matrix

// Matrix -
type Matrix [][]int

func (m Matrix) getRowsCount() int {
	if !m.isValid() {
		return 0
	}

	return len(m)
}

func (m Matrix) getColumnsCount() int {
	if !m.isValid() {
		return 0
	}

	if len(m) > 0 {
		firstRow := m[0]

		return len(firstRow)
	}

	return 0
}

func (m Matrix) getRow(n int) []int {
	if !m.isValid() {
		return []int{}
	}

	if m.getRowsCount() < n || n < 1 {
		return []int{}
	}

	return m[n-1]
}

func (m Matrix) getColumn(n int) []int {
	if !m.isValid() {
		return []int{}
	}

	if n < 1 || m.getColumnsCount() < n {
		return []int{}
	}

	var column []int
	for i := 0; i < m.getRowsCount(); i++ {
		row := m[i]

		column = append(column, row[n-1])
	}

	return column
}

func (m Matrix) getDiaginal() []int {
	if !m.isDiagonal() {
		return []int{}
	}

	var diagonal []int
	for diagonalIndex := 0; diagonalIndex < m.getRowsCount(); diagonalIndex++ {
		row := m[diagonalIndex]
		element := row[diagonalIndex]

		diagonal = append(diagonal, element)
	}

	return diagonal
}

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

func (m Matrix) isSquare() bool {
	if !m.isValid() {
		return false
	}

	return m.getRowsCount() == m.getColumnsCount()
}

func (m Matrix) isNull() bool {
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

func (m Matrix) isDiagonal() bool {
	if !m.isSquare() {
		return false
	}

	if m.getRowsCount() < 2 {
		return false
	}

	for diagonalIndex := 0; diagonalIndex < m.getRowsCount(); diagonalIndex++ {
		row := m[diagonalIndex]

		for rowIndex := 0; rowIndex < len(row); rowIndex++ {
			isZero := row[rowIndex] == 0

			if rowIndex == diagonalIndex && isZero {
				return false
			}

			if rowIndex != diagonalIndex && !isZero {
				return false
			}

		}
	}

	return true
}

func (m Matrix) isUnit() bool {
	if !m.isDiagonal() {
		return false
	}

	diagonal := m.getDiaginal()

	for _, number := range diagonal {
		if number != 1 {
			return false
		}
	}

	return true
}

func (m Matrix) isScalar() bool {
	if !m.isDiagonal() {
		return false
	}

	diagonal := m.getDiaginal()

	for _, currentNumber := range diagonal {
		if currentNumber != diagonal[0] {
			return false
		}
	}

	return true
}

func (m Matrix) isUpperTriangular() bool {
	if !m.isSquare() || m.getRowsCount() < 2 {
		return false
	}

	for rowIndex := 0; rowIndex < m.getRowsCount(); rowIndex++ {
		row := m[rowIndex]
		zeros := row[0:rowIndex]

		for _, number := range zeros {
			if number != 0 {
				return false
			}
		}

	}

	return true
}

func (m Matrix) isLowerTriangular() bool {
	if !m.isSquare() || m.getRowsCount() < 2 {
		return false
	}

	for rowIndex := 0; rowIndex < m.getRowsCount(); rowIndex++ {
		row := m[rowIndex]
		zeros := row[rowIndex+1:]

		for _, number := range zeros {
			if number != 0 {
				return false
			}
		}

	}

	return true
}
