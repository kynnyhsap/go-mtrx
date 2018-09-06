package matrix

import (
	"reflect"
	"testing"
)

func TestIsValidMatrix(t *testing.T) {
	tests := []struct {
		matrix Matrix
		want   bool
	}{
		{
			want: true,
			matrix: Matrix{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{0, 0, 0},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{0, 0, 0},
				{0, 1},
				{0},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{0},
			},
		},
		{
			want:   true,
			matrix: Matrix{},
		},
	}

	for _, tt := range tests {
		if got := tt.matrix.isValid(); got != tt.want {
			t.Errorf("Matrix.isValid() = %v, want %v", got, tt.want)
		}
	}
}

func TestIsNullMatrix(t *testing.T) {
	tests := []struct {
		matrix Matrix
		want   bool
	}{
		{
			want: true,
			matrix: Matrix{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{0, 0, 0},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		if got := tt.matrix.isNull(); got != tt.want {
			t.Errorf("Matrix.isNull() = %v, want %v", got, tt.want)
		}
	}
}

func TestGetRowsCount(t *testing.T) {
	tests := []struct {
		matrix Matrix
		want   int
	}{
		{
			want: 3,
			matrix: Matrix{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			want: 1,
			matrix: Matrix{
				{0, 0, 0},
			},
		},
		{
			want:   0,
			matrix: Matrix{},
		},
		{
			want: 6,
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

	for _, tt := range tests {
		if got := tt.matrix.getRowsCount(); got != tt.want {
			t.Errorf("Matrix.getRowsCount() = %v, want %v", got, tt.want)
		}
	}
}

func TestGetColumnsCount(t *testing.T) {
	tests := []struct {
		matrix Matrix
		want   int
	}{
		{
			want: 3,
			matrix: Matrix{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			want: 2,
			matrix: Matrix{
				{0, 0},
			},
		},
		{
			want:   0,
			matrix: Matrix{},
		},
		{
			want: 6,
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

	for _, tt := range tests {
		if got := tt.matrix.getColumnsCount(); got != tt.want {
			t.Errorf("Matrix.getColumnsCount() = %v, want %v", got, tt.want)
		}
	}
}

func TestIsDiagonalMatrix(t *testing.T) {
	tests := []struct {
		matrix Matrix
		want   bool
	}{
		{
			want: true,
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
				{0, 0, 3},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{0, 0},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{0},
			},
		},
		{
			want:   false,
			matrix: Matrix{},
		},
		{
			want: true,
			matrix: Matrix{
				{53, 0, 0, 0},
				{0, -234, 0, 0},
				{0, 0, 234, 0},
				{0, 0, 0, 745},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{4, 0, 0},
				{0, 5, 0},
				{0, 0, 0},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{4, 0, 0},
				{0, 0, 0},
				{0, 0, 6},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{4, 0},
				{0, 1},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{0, 1},
				{5, 0},
			},
		},
	}

	for _, tt := range tests {
		if got := tt.matrix.isDiagonal(); got != tt.want {
			t.Errorf("Matrix.isDiagonal() = %v, want %v", got, tt.want)
		}
	}
}

func TestGetRow(t *testing.T) {
	tests := []struct {
		matrix Matrix
		index  int
		want   []int
	}{
		{
			index: 2,
			want:  []int{0, 2, 0},
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
			},
		},
		{
			index: 1,
			want:  []int{4},
			matrix: Matrix{
				{4},
				{1},
				{0},
			},
		},
		{
			index: 3,
			want:  []int{0, 59, 23},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			index: 4,
			want:  []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			index: 30,
			want:  []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			index: -30,
			want:  []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			index: 1,
			want:  []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6},
				{0},
			},
		},
	}

	for _, tt := range tests {
		if got := tt.matrix.getRow(tt.index); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Matrix.isDiagonal(%v) = %v, want %v", tt.index, got, tt.want)
		}
	}
}

func TestGetColumn(t *testing.T) {
	tests := []struct {
		matrix Matrix
		index  int
		want   []int
	}{
		{
			index: 2,
			want:  []int{0, 2},
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
			},
		},
		{
			index: 1,
			want:  []int{4, 1, 0},
			matrix: Matrix{
				{4},
				{1},
				{0},
			},
		},
		{
			index: 3,
			want:  []int{5, 1, 23},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 9, 23},
			},
		},
		{
			index: 0,
			want:  []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			index: 30,
			want:  []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			index: -30,
			want:  []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6, 1},
				{0, 59, 23},
			},
		},
		{
			index: 1,
			want:  []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6},
				{0},
			},
		},
	}

	for _, tt := range tests {
		if got := tt.matrix.getColumn(tt.index); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Matrix.getColumn(%v) = %v, want %v", tt.index, got, tt.want)
		}
	}
}

func TestGetDiagonal(t *testing.T) {
	tests := []struct {
		matrix Matrix
		want   []int
	}{
		{
			want: []int{},
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
			},
		},
		{
			want: []int{},
			matrix: Matrix{
				{4},
				{1},
				{0},
			},
		},
		{
			want:   []int{},
			matrix: Matrix{},
		},
		{
			want: []int{4, 6, 23},
			matrix: Matrix{
				{4, 0, 0},
				{0, 6, 0},
				{0, 0, 23},
			},
		},
		{
			want: []int{},
			matrix: Matrix{
				{0, 0, 2},
				{0, 6, 0},
				{2, 0, 0},
			},
		},
		{
			want: []int{},
			matrix: Matrix{
				{4, 3, 5},
				{1, 6},
				{0},
			},
		},
		{
			want: []int{43, 61},
			matrix: Matrix{
				{43, 0},
				{0, 61},
			},
		},
		{
			want: []int{},
			matrix: Matrix{
				{43},
			},
		},
	}

	for _, tt := range tests {

		if got := tt.matrix.getDiaginal(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Matrix.getDiaginal() = %v, want %v", got, tt.want)
		}
	}
}

func TestIsUnitMatrix(t *testing.T) {
	tests := []struct {
		matrix Matrix
		want   bool
	}{
		{
			want: false,
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{4},
				{1},
				{0},
			},
		},
		{
			want:   false,
			matrix: Matrix{},
		},
		{
			want: false,
			matrix: Matrix{
				{4, 0, 0},
				{0, 6, 0},
				{0, 0, 23},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{4, 3, 5},
				{1, 6},
				{0},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{1, 0},
				{0, 1},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{1},
			},
		},
	}

	for _, tt := range tests {
		if got := tt.matrix.isUnit(); got != tt.want {
			t.Errorf("Matrix.isUnit() = %v, want %v", got, tt.want)
		}
	}
}

func TestIsUpperTriangularMatrix(t *testing.T) {
	tests := []struct {
		matrix Matrix
		want   bool
	}{
		{
			want: true,
			matrix: Matrix{
				{1, 2, 4},
				{0, 2, 3},
				{0, 0, 1},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{1, 2, 4, 6},
				{0, 2, 3, 2},
				{0, 0, 1, 6},
				{0, 0, 0, 7},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{1, 2},
				{0, 2},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{1, 2},
				{0},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{1, 2, 0},
				{0, 10, 20},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{0},
			},
		},
		{
			want:   false,
			matrix: Matrix{},
		},
	}

	for _, tt := range tests {
		if got := tt.matrix.isUpperTriangular(); got != tt.want {
			t.Errorf("Matrix.isUpperTriangular() = %v, want %v", got, tt.want)
		}
	}
}

func TestIsLowerTriangularMatrix(t *testing.T) {
	tests := []struct {
		matrix Matrix
		want   bool
	}{
		{
			want: true,
			matrix: Matrix{
				{1, 0, 0},
				{6, 2, 0},
				{7, 3, 1},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{1, 0, 0, 0},
				{5, 2, 0, 0},
				{3, 5, 1, 0},
				{5, 3, 7, 7},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{1, 0},
				{4, 2},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{1, 0},
				{1},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{1, 0, 0},
				{4, 10, 0},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{0},
			},
		},
		{
			want:   false,
			matrix: Matrix{},
		},
	}

	for _, tt := range tests {
		if got := tt.matrix.isLowerTriangular(); got != tt.want {
			t.Errorf("Matrix.isLowerTriangular() = %v, want %v", got, tt.want)
		}
	}
}

func TestIsScalarMatrix(t *testing.T) {
	tests := []struct {
		matrix Matrix
		want   bool
	}{
		{
			want: true,
			matrix: Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{1, 0, 0},
				{0, 2, 0},
				{0, 0, 1},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{2, 0, 0, 0},
				{0, 2, 0, 0},
				{0, 0, 2, 0},
				{0, 0, 0, 1},
			},
		},
		{
			want: true,
			matrix: Matrix{
				{2, 0},
				{0, 2},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{2, 0, 0},
				{0, 2, 0},
			},
		},
		{
			want: false,
			matrix: Matrix{
				{0},
			},
		},
		{
			want:   false,
			matrix: Matrix{},
		},
	}

	for _, tt := range tests {
		if got := tt.matrix.isScalar(); got != tt.want {
			t.Errorf("Matrix.isScalar() = %v, want %v", got, tt.want)
		}
	}
}

func TestIsSameSizeWith(t *testing.T) {
	tests := []struct {
		matrixA Matrix
		matrixB Matrix
		want    bool
	}{
		{
			want: true,
			matrixA: Matrix{
				{13456354, 34564356, 2345234},
				{2345230, 12345234, 345635},
				{34634, 3456345, 34563451},
			},
			matrixB: Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			want: true,
			matrixA: Matrix{
				{56, 32, 5},
			},
			matrixB: Matrix{
				{1, 0, 0},
			},
		},
		{
			want: true,
			matrixA: Matrix{
				{56},
				{32},
				{5},
			},
			matrixB: Matrix{
				{1},
				{0},
				{0},
			},
		},
		{
			want: true,
			matrixA: Matrix{
				{56},
			},
			matrixB: Matrix{
				{1},
			},
		},
		{
			want:    true,
			matrixA: Matrix{},
			matrixB: Matrix{},
		},
		{
			want: false,
			matrixA: Matrix{
				{13456354, 34564356, 2345234},
				{2345230, 12345234, 345635},
			},
			matrixB: Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			want: false,
			matrixA: Matrix{
				{13456354, 34564356, 2345234},
				{2345230, 12345234, 345635},
			},
			matrixB: Matrix{
				{1, 0},
				{0, 1},
				{0, 0},
			},
		},
		{
			want:    false,
			matrixA: Matrix{},
			matrixB: Matrix{
				{1},
			},
		},
	}

	for _, tt := range tests {
		if got := tt.matrixA.isSameSizeWith(tt.matrixB); got != tt.want {
			t.Errorf("Matrix.isSameSizeWith() = %v, want %v", got, tt.want)
		}
	}
}

func TestIsMatchedWith(t *testing.T) {
	tests := []struct {
		matrixA Matrix
		matrixB Matrix
		want    bool
	}{
		{
			want: true,
			matrixA: Matrix{
				{13456354, 34564356, 2345234},
				{2345230, 12345234, 345635},
				{34634, 3456345, 34563451},
			},
			matrixB: Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
		{
			want: true,
			matrixA: Matrix{
				{56, 32, 5},
			},
			matrixB: Matrix{
				{56},
				{32},
				{5},
			},
		},
		{
			want: true,
			matrixA: Matrix{
				{56},
				{32},
				{5},
			},
			matrixB: Matrix{
				{56, 32, 5},
			},
		},
		{
			want: true,
			matrixA: Matrix{
				{56},
			},
			matrixB: Matrix{
				{56},
			},
		},
		{
			want: true,
			matrixA: Matrix{
				{56, 40},
				{32, 45},
				{5, 22},
				{5, 22},
				{5, 22},
				{5, 22},
			},
			matrixB: Matrix{
				{56},
				{56},
			},
		},
		{
			want: true,
			matrixA: Matrix{
				{56, 40},
				{32, 45},
				{5, 22},
				{5, 22},
				{5, 22},
				{5, 22},
			},
			matrixB: Matrix{
				{56},
				{56},
			},
		},
		{
			want: false,
			matrixA: Matrix{
				{56, 40, 4, 3345, 3, 45, 34},
				{32, 45},
			},
			matrixB: Matrix{
				{56},
				{56},
				{56},
			},
		},
		{
			want: false,
			matrixA: Matrix{
				{56, 40},
				{32, 45},
			},
			matrixB: Matrix{
				{56},
				{56},
				{56},
			},
		},
	}

	for _, tt := range tests {
		if got := tt.matrixA.isMatchedWith(tt.matrixB); got != tt.want {
			t.Errorf("Matrix.isMatchedWith(%v) = %v, want %v", tt.matrixB, got, tt.want)
		}
	}
}
