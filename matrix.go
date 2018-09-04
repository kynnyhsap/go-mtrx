package matrix

// Matrix -
type Matrix [][]int

// Dimention - the dimention of matrix
type Dimention struct {
	rows    int
	columns int
}

func (m Matrix) getDimention() Dimention {
	return Dimention{
		rows:    m.getRowsCount(),
		columns: m.getColumnsCount(),
	}
}

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

	return m.getRowsCount() == m.getColumnsCount()
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

func (m Matrix) isDiagonalMatrix() bool {
	if !m.isSquareMatrix() {
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

	println(n)

	var column []int
	for i := 0; i < m.getRowsCount(); i++ {
		row := m[i]

		column = append(column, row[n-1])
	}

	return column
}
