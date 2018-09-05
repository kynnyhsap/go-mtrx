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

func TestGetDiagonal(t *testing.T) {
	tables := []struct {
		matrix   Matrix
		diagonal []int
	}{
		{
			diagonal: []int{},
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
			},
		},
		{
			diagonal: []int{},
			matrix: Matrix{
				{4},
				{1},
				{0},
			},
		},
		{
			diagonal: []int{},
			matrix:   Matrix{},
		},
		{
			diagonal: []int{4, 6, 23},
			matrix: Matrix{
				{4, 0, 0},
				{0, 6, 0},
				{0, 0, 23},
			},
		},
		{
			diagonal: []int{},
			matrix: Matrix{
				{0, 0, 2},
				{0, 6, 0},
				{2, 0, 0},
			},
		},
		{
			diagonal: []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6},
				{0},
			},
		},
		{
			diagonal: []int{43, 61},
			matrix: Matrix{
				{43, 0},
				{0, 61},
			},
		},
		{
			diagonal: []int{},
			matrix: Matrix{
				{43},
			},
		},
	}

	for _, table := range tables {
		diagonal := table.matrix.getDiaginal()
		isSameDiagonal := reflect.DeepEqual(diagonal, table.diagonal)

		if !isSameDiagonal {
			t.Errorf("method getDiagonal \nGot: %v\nWant: %v ", diagonal, table.diagonal)
		}
	}
}

func TestIsUnitMatrix(t *testing.T) {
	tables := []struct {
		matrix Matrix
		isUnit bool
	}{
		{
			isUnit: false,
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
			},
		},
		{
			isUnit: false,
			matrix: Matrix{
				{4},
				{1},
				{0},
			},
		},
		{
			isUnit: false,
			matrix: Matrix{},
		},
		{
			isUnit: false,
			matrix: Matrix{
				{4, 0, 0},
				{0, 6, 0},
				{0, 0, 23},
			},
		},
		{
			isUnit: true,
			matrix: Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			isUnit: false,
			matrix: Matrix{
				{4, 3, 5},
				{1, 6},
				{0},
			},
		},
		{
			isUnit: true,
			matrix: Matrix{
				{1, 0},
				{0, 1},
			},
		},
		{
			isUnit: false,
			matrix: Matrix{
				{1},
			},
		},
	}

	for _, table := range tables {
		isUnit := table.matrix.isUnitMatrix()

		if isUnit != table.isUnit {
			t.Errorf("method isUnitMatrix \nGot: %v\nWant: %v \nMatrix: %v", isUnit, table.isUnit, table.matrix)
		}
	}
}

func TestIsUpperTriangularMatrix(t *testing.T) {
	tables := []struct {
		matrix  Matrix
		isUpper bool
	}{
		{
			isUpper: true,
			matrix: Matrix{
				{1, 2, 4},
				{0, 2, 3},
				{0, 0, 1},
			},
		},
		{
			isUpper: true,
			matrix: Matrix{
				{1, 2, 4, 6},
				{0, 2, 3, 2},
				{0, 0, 1, 6},
				{0, 0, 0, 7},
			},
		},
		{
			isUpper: true,
			matrix: Matrix{
				{1, 2},
				{0, 2},
			},
		},
		{
			isUpper: false,
			matrix: Matrix{
				{1, 2},
				{0},
			},
		},
		{
			isUpper: false,
			matrix: Matrix{
				{1, 2, 0},
				{0, 10, 20},
			},
		},
		{
			isUpper: false,
			matrix: Matrix{
				{0},
			},
		},
		{
			isUpper: false,
			matrix:  Matrix{},
		},
	}

	for _, table := range tables {
		isUpper := table.matrix.isUpperTriangularMatrix()

		if isUpper != table.isUpper {
			t.Errorf("method isUpperTriangularMatrix \nGot: %v\nWant: %v \nMatrix: %v", isUpper, table.isUpper, table.matrix)
		}
	}
}

func TestIsLowerTriangularMatrix(t *testing.T) {
	tables := []struct {
		matrix  Matrix
		isLower bool
	}{
		{
			isLower: true,
			matrix: Matrix{
				{1, 0, 0},
				{6, 2, 0},
				{7, 3, 1},
			},
		},
		{
			isLower: true,
			matrix: Matrix{
				{1, 0, 0, 0},
				{5, 2, 0, 0},
				{3, 5, 1, 0},
				{5, 3, 7, 7},
			},
		},
		{
			isLower: true,
			matrix: Matrix{
				{1, 0},
				{4, 2},
			},
		},
		{
			isLower: false,
			matrix: Matrix{
				{1, 0},
				{1},
			},
		},
		{
			isLower: false,
			matrix: Matrix{
				{1, 0, 0},
				{4, 10, 0},
			},
		},
		{
			isLower: false,
			matrix: Matrix{
				{0},
			},
		},
		{
			isLower: false,
			matrix:  Matrix{},
		},
	}

	for _, table := range tables {
		isLower := table.matrix.isLowerTriangularMatrix()

		if isLower != table.isLower {
			t.Errorf("method isLowerTriangularMatrix \nGot: %v\nWant: %v \nMatrix: %v", isLower, table.isLower, table.matrix)
		}
	}
}

func TestIsScalarMatrix(t *testing.T) {
	tables := []struct {
		matrix   Matrix
		isScalar bool
	}{
		{
			isScalar: true,
			matrix: Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			isScalar: false,
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
				{0, 0, 1},
			},
		},
		{
			isScalar: false,
			matrix: Matrix{
				{2, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 1},
			},
		},
		{
			isScalar: true,
			matrix: Matrix{
				{2, 0},
				{0, 2},
			},
		},
		{
			isScalar: false,
			matrix: Matrix{
				{2, 0, 0},
				{0, 2, 0},
			},
		},
		{
			isScalar: false,
			matrix: Matrix{
				{0},
			},
		},
		{
			isScalar: false,
			matrix:   Matrix{},
		},
	}

	for _, table := range tables {
		isScalar := table.matrix.isScalarMatrix()

		if isScalar != table.isScalar {
			t.Errorf("method isScalarMatrix \nGot: %v\nWant: %v \nMatrix: %v", isScalar, table.isScalar, table.matrix)
		}
	}
}
