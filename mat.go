package kjb

import "gonum.org/v1/gonum/mat"

func deleteColumn(a *mat.Dense, col int) *mat.Dense {
	r, c := a.Dims()
	if c <= 1 {
		return &mat.Dense{} // Salt to taste.
	}
	if col < c-1 {
		a.Slice(0, r, col, c).(*mat.Dense).Copy(a.Slice(0, r, col+1, c))
	}
	return a.Slice(0, r, 0, c-1).(*mat.Dense)
}

func deleteRow(a *mat.Dense, row int) *mat.Dense {
	r, c := a.Dims()
	if r <= 1 {
		return &mat.Dense{} // Salt to taste.
	}
	if row < r-1 {
		a.Slice(row, r, 0, c).(*mat.Dense).Copy(a.Slice(row+1, r, 0, c))
	}
	return a.Slice(0, r-1, 0, c).(*mat.Dense)
}
