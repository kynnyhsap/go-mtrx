package main

import "testing"

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

func TestCountOfRows(t *testing.T) {
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
		countRows := table.matrix.countRows()
		if countRows != table.rows {
			t.Errorf("method countRows \nGot: %d\nWant: %d ", countRows, table.rows)
		}
	}
}

func TestCountOfColumns(t *testing.T) {
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
		countColumns := table.matrix.countColumns()
		if countColumns != table.columns {
			t.Errorf("method countColumns \nGot: %d\nWant: %d ", countColumns, table.columns)
		}
	}
}
