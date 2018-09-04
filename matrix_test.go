package matrix

import (
	"reflect"
	"testing"
)

func TestIsValidMatrix(t *testing.T) {
	validMatrix := Matrix{
		{1, 3, 4},
		{12, -9, 9},
		{6, 1, -8},
	}

	if !validMatrix.isValid() {
		t.Errorf("Err")
	}

	invalidMatrix := Matrix{
		{1, 3, 4, 7},
		{1},
		{6, 1, -8},
	}

	if invalidMatrix.isValid() {
		t.Errorf("Err")
	}
}

func TestIsNullMatrix(t *testing.T) {
	tables := []struct {
		matrix Matrix
		isNull bool
	}{
		{
			isNull: true,
			matrix: Matrix{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			isNull: true,
			matrix: Matrix{
				{0, 0, 0},
			},
		},
		{
			isNull: false,
			matrix: Matrix{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
	}

	for _, table := range tables {
		isNullMatrix := table.matrix.isNullMatrix()
		if isNullMatrix != table.isNull {
			t.Errorf("method isNullMatrix \nGot: %t\nWant: %t ", isNullMatrix, table.isNull)
		}
	}
}

func TestGetRowsCount(t *testing.T) {
	tables := []struct {
		matrix Matrix
		rows   int
	}{
		{
			rows: 3,
			matrix: Matrix{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			rows: 1,
			matrix: Matrix{
				{0, 0, 0},
			},
		},
		{
			rows:   0,
			matrix: Matrix{},
		},
		{
			rows: 6,
			matrix: Matrix{
				{0},
				{0},
				{0},
				{0},
				{0},
				{0},
			},
		},
	}

	for _, table := range tables {
		countRows := table.matrix.getRowsCount()
		if countRows != table.rows {
			t.Errorf("method countRows \nGot: %d\nWant: %d ", countRows, table.rows)
		}
	}
}

func TestGetColumnsCount(t *testing.T) {
	tables := []struct {
		matrix  Matrix
		columns int
	}{
		{
			columns: 3,
			matrix: Matrix{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			columns: 2,
			matrix: Matrix{
				{0, 0},
			},
		},
		{
			columns: 0,
			matrix:  Matrix{},
		},
		{
			columns: 6,
			matrix: Matrix{
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
			},
		},
	}

	for _, table := range tables {
		countColumns := table.matrix.getColumnsCount()
		if countColumns != table.columns {
			t.Errorf("method countColumns \nGot: %d\nWant: %d ", countColumns, table.columns)
		}
	}
}

func TestIsDiagonalMatrix(t *testing.T) {
	tables := []struct {
		matrix     Matrix
		isDiagonal bool
	}{
		{
			isDiagonal: true,
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
				{0, 0, 3},
			},
		},
		{
			isDiagonal: false,
			matrix: Matrix{
				{0, 0},
			},
		},
		{
			isDiagonal: false,
			matrix: Matrix{
				{0},
			},
		},
		{
			isDiagonal: false,
			matrix:     Matrix{},
		},
		{
			isDiagonal: true,
			matrix: Matrix{
				{53, 0, 0, 0},
				{0, -234, 0, 0},
				{0, 0, 234, 0},
				{0, 0, 0, 745},
			},
		},
		{
			isDiagonal: false,
			matrix: Matrix{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			isDiagonal: false,
			matrix: Matrix{
				{4, 0, 0},
				{0, 5, 0},
				{0, 0, 0},
			},
		},
		{
			isDiagonal: false,
			matrix: Matrix{
				{4, 0, 0},
				{0, 0, 0},
				{0, 0, 6},
			},
		},
		{
			isDiagonal: true,
			matrix: Matrix{
				{4, 0},
				{0, 1},
			},
		},
		{
			isDiagonal: false,
			matrix: Matrix{
				{0, 1},
				{5, 0},
			},
		},
	}

	for _, table := range tables {
		isDiagonalMatrix := table.matrix.isDiagonalMatrix()
		if isDiagonalMatrix != table.isDiagonal {
			t.Errorf("method isDiagonalMatrix \nGot: %t\nWant: %t ", isDiagonalMatrix, table.isDiagonal)
		}
	}
}

func TestGetRow(t *testing.T) {
	tables := []struct {
		matrix   Matrix
		rowIndex int
		row      []int
	}{
		{
			rowIndex: 2,
			row:      []int{0, 2, 0},
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
			},
		},
		{
			rowIndex: 1,
			row:      []int{4},
			matrix: Matrix{
				{4},
				{1},
				{0},
			},
		},
		{
			rowIndex: 3,
			row:      []int{0, 59, 23},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			rowIndex: 4,
			row:      []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			rowIndex: 30,
			row:      []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			rowIndex: -30,
			row:      []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			rowIndex: 1,
			row:      []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6},
				{0},
			},
		},
	}

	for _, table := range tables {
		row := table.matrix.getRow(table.rowIndex)
		isSameRows := reflect.DeepEqual(row, table.row)

		if !isSameRows {
			t.Errorf("method getRow \nGot: %v\nWant: %v ", row, table.row)
		}
	}
}

func TestGetColumn(t *testing.T) {
	tables := []struct {
		matrix      Matrix
		columnIndex int
		column      []int
	}{
		{
			columnIndex: 2,
			column:      []int{0, 2},
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
			},
		},
		{
			columnIndex: 1,
			column:      []int{4, 1, 0},
			matrix: Matrix{
				{4},
				{1},
				{0},
			},
		},
		{
			columnIndex: 3,
			column:      []int{5, 1, 23},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 9, 23},
			},
		},
		{
			columnIndex: 0,
			column:      []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			columnIndex: 30,
			column:      []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			columnIndex: -30,
			column:      []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			columnIndex: 1,
			column:      []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6},
				{0},
			},
		},
	}

	for _, table := range tables {
		column := table.matrix.getColumn(table.columnIndex)
		isSameComlumn := reflect.DeepEqual(column, table.column)

		if !isSameComlumn {
			t.Errorf("method getColumn \nGot: %v\nWant: %v ", column, table.column)
		}
	}
}
