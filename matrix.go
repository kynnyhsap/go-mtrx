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

func (m Matrix) isSquare() bool {
	if !m.isValid() {
		return false
	}

	return len(m) == len(m[0])
}

func main() {

}
