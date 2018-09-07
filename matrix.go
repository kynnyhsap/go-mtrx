package matrix

// Matrix -
type Matrix [][]int

func createMatrix(rowsCount int, columnsCount int) Matrix {

	var rows Matrix = make([][]int, rowsCount)

	for i := 0; i < rowsCount; i++ {
		column := make([]int, columnsCount)
		rows[i] = column
	}

	return rows
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

func (m Matrix) getRow(n int) []int {
	if !m.isValid() {
		return []int{}
	}

	if m.getRowsCount() < n || n < 1 {
		return []int{}
	}

	return m[n-1]
}

func (m Matrix) setRow(n int, row []int) Matrix {
	if !m.isValid() {
		return Matrix{}
	}

	if m.getRowsCount() < n || n < 1 {
		return Matrix{}
	}

	if m.getColumnsCount() != len(row) {
		return Matrix{}
	}

	m[n-1] = row

	return m
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

func (m Matrix) setColumn(n int, column []int) Matrix {
	if !m.isValid() {
		return Matrix{}
	}

	if n < 1 || m.getColumnsCount() < n {
		return Matrix{}
	}

	if m.getRowsCount() != len(column) {
		return Matrix{}
	}

	for rowIndex := 0; rowIndex < m.getRowsCount(); rowIndex++ {
		row := m[rowIndex]
		row[n-1] = column[rowIndex]
	}

	return m
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

func (m Matrix) isSameSizeWith(b Matrix) bool {
	if !m.isValid() || !b.isValid() {
		return false
	}

	if m.getRowsCount() != b.getRowsCount() || m.getColumnsCount() != b.getColumnsCount() {
		return false
	}

	return true
}

func (m Matrix) isMatchedWith(b Matrix) bool {
	if !m.isValid() || !b.isValid() {
		return false
	}

	if m.getColumnsCount() != b.getRowsCount() {
		return false
	}

	return true
}

func (m Matrix) add(b Matrix) Matrix {
	if !m.isValid() || !b.isValid() {
		return Matrix{}
	}

	if !m.isSameSizeWith(b) {
		return Matrix{}
	}

	for i := 0; i < m.getRowsCount(); i++ {
		for j := 0; j < m.getColumnsCount(); j++ {
			m[i][j] = m[i][j] + b[i][j]
		}
	}

	return m
}

func (m Matrix) multipyBy(b Matrix) Matrix {
	if !m.isValid() || !b.isValid() {
		return Matrix{}
	}

	if !m.isMatchedWith(b) {
		return Matrix{}
	}

	var newMatrix [][]int
	for rowIndex := 1; rowIndex <= m.getRowsCount(); rowIndex++ {
		row := m.getRow(rowIndex)

		var newRow []int
		for columnIndex := 1; columnIndex <= b.getColumnsCount(); columnIndex++ {
			column := b.getColumn(columnIndex)

			var number int
			for i := 0; i < len(row); i++ {
				number += row[i] * column[i]
			}

			newRow = append(newRow, number)
		}

		newMatrix = append(newMatrix, newRow)
	}

	return Matrix(newMatrix)
}

func (m Matrix) transpose() Matrix {
	if !m.isValid() {
		return Matrix{}
	}

	c := createMatrix(m.getColumnsCount(), m.getRowsCount())

	for i := 1; i <= m.getRowsCount(); i++ {
		row := m.getRow(i)
		c.setColumn(i, row)
	}

	return c
}
