package matrix

import "fmt"

// Matrix -
type Matrix [][]int

func CreateMatrix(rowsCount int, columnsCount int) Matrix {

	var rows Matrix = make([][]int, rowsCount)

	for i := 0; i < rowsCount; i++ {
		column := make([]int, columnsCount)
		rows[i] = column
	}

	return rows
}

func (m Matrix) GetRowsCount() int {
	if !m.IsValid() {
		return 0
	}

	return len(m)
}

func (m Matrix) GetColumnsCount() int {
	if !m.IsValid() {
		return 0
	}

	if len(m) > 0 {
		firstRow := m[0]

		return len(firstRow)
	}

	return 0
}

func (m Matrix) GetRow(n int) []int {
	if !m.IsValid() {
		return []int{}
	}

	if m.GetRowsCount() < n || n < 1 {
		return []int{}
	}

	return m[n-1]
}

func (m Matrix) SetRow(n int, row []int) Matrix {
	if !m.IsValid() {
		return Matrix{}
	}

	if m.GetRowsCount() < n || n < 1 {
		return Matrix{}
	}

	if m.GetColumnsCount() != len(row) {
		return Matrix{}
	}

	m[n-1] = row

	return m
}

func (m Matrix) GetColumn(n int) []int {
	if !m.IsValid() {
		return []int{}
	}

	if n < 1 || m.GetColumnsCount() < n {
		return []int{}
	}

	var column []int
	for i := 0; i < m.GetRowsCount(); i++ {
		row := m[i]

		column = append(column, row[n-1])
	}

	return column
}

func (m Matrix) SetColumn(n int, column []int) Matrix {
	if !m.IsValid() {
		return Matrix{}
	}

	if n < 1 || m.GetColumnsCount() < n {
		return Matrix{}
	}

	if m.GetRowsCount() != len(column) {
		return Matrix{}
	}

	for rowIndex := 0; rowIndex < m.GetRowsCount(); rowIndex++ {
		row := m[rowIndex]
		row[n-1] = column[rowIndex]
	}

	return m
}

func (m Matrix) GetDiaginal() []int {
	if !m.IsDiagonal() {
		return []int{}
	}

	var diagonal []int
	for diagonalIndex := 0; diagonalIndex < m.GetRowsCount(); diagonalIndex++ {
		row := m[diagonalIndex]
		element := row[diagonalIndex]

		diagonal = append(diagonal, element)
	}

	return diagonal
}

func (m Matrix) IsValid() bool {
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

func (m Matrix) IsSquare() bool {
	if !m.IsValid() {
		return false
	}

	return m.GetRowsCount() == m.GetColumnsCount()
}

func (m Matrix) IsNull() bool {
	if !m.IsValid() {
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

func (m Matrix) IsDiagonal() bool {
	if !m.IsSquare() {
		return false
	}

	if m.GetRowsCount() < 2 {
		return false
	}

	for diagonalIndex := 0; diagonalIndex < m.GetRowsCount(); diagonalIndex++ {
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

func (m Matrix) IsUnit() bool {
	if !m.IsDiagonal() {
		return false
	}

	diagonal := m.GetDiaginal()

	for _, number := range diagonal {
		if number != 1 {
			return false
		}
	}

	return true
}

func (m Matrix) IsScalar() bool {
	if !m.IsDiagonal() {
		return false
	}

	diagonal := m.GetDiaginal()

	for _, currentNumber := range diagonal {
		if currentNumber != diagonal[0] {
			return false
		}
	}

	return true
}

func (m Matrix) IsUpperTriangular() bool {
	if !m.IsSquare() || m.GetRowsCount() < 2 {
		return false
	}

	for rowIndex := 0; rowIndex < m.GetRowsCount(); rowIndex++ {
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

func (m Matrix) IsLowerTriangular() bool {
	if !m.IsSquare() || m.GetRowsCount() < 2 {
		return false
	}

	for rowIndex := 0; rowIndex < m.GetRowsCount(); rowIndex++ {
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

func (m Matrix) IsSameSizeWith(b Matrix) bool {
	if !m.IsValid() || !b.IsValid() {
		return false
	}

	if m.GetRowsCount() != b.GetRowsCount() || m.GetColumnsCount() != b.GetColumnsCount() {
		return false
	}

	return true
}

func (m Matrix) IsMatchedWith(b Matrix) bool {
	if !m.IsValid() || !b.IsValid() {
		return false
	}

	if m.GetColumnsCount() != b.GetRowsCount() {
		return false
	}

	return true
}

func (m Matrix) Add(b Matrix) Matrix {
	if !m.IsValid() || !b.IsValid() {
		return Matrix{}
	}

	if !m.IsSameSizeWith(b) {
		return Matrix{}
	}

	for i := 0; i < m.GetRowsCount(); i++ {
		for j := 0; j < m.GetColumnsCount(); j++ {
			m[i][j] = m[i][j] + b[i][j]
		}
	}

	return m
}

func (m Matrix) MultipyBy(b Matrix) Matrix {
	if !m.IsValid() || !b.IsValid() {
		return Matrix{}
	}

	if !m.IsMatchedWith(b) {
		return Matrix{}
	}

	var newMatrix [][]int
	for rowIndex := 1; rowIndex <= m.GetRowsCount(); rowIndex++ {
		row := m.GetRow(rowIndex)

		var newRow []int
		for columnIndex := 1; columnIndex <= b.GetColumnsCount(); columnIndex++ {
			column := b.GetColumn(columnIndex)

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

func (m Matrix) Transpose() Matrix {
	if !m.IsValid() {
		return Matrix{}
	}

	c := CreateMatrix(m.GetColumnsCount(), m.GetRowsCount())

	for i := 1; i <= m.GetRowsCount(); i++ {
		row := m.GetRow(i)
		c.SetColumn(i, row)
	}

	return c
}

func (m Matrix) String() string {
	var str string

	if m.GetRowsCount() == 0 {
		str = "\n[ ]\n"
	}

	for i := 1; i <= m.GetRowsCount(); i++ {
		row := m.GetRow(i)

		if i == 1 {
			str += fmt.Sprintf("\n%v\n", row)
		} else {
			str += fmt.Sprintf("%v\n", row)
		}
	}

	return str
}
