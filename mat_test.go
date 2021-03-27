package kjb

import (
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestDeleteCol(t *testing.T) {
	var tests = []struct{ Mat *mat.Dense }{
		{Mat: mat.NewDense(1, 1, floatTo(1))},
		{Mat: mat.NewDense(2, 2, floatTo(2*2))},
		{Mat: mat.NewDense(5, 5, floatTo(5*5))},
	}
	for _, v := range tests {
		r, c := v.Mat.Dims()
		del := rand.Intn(c)
		res := deleteColumn(v.Mat, del)
		r1, c1 := res.Dims()
		if r != r1 && c != 1 {
			t.Error("number of rows should not change unless only one column")
		}
		if c1 != c-1 {
			t.Error("number of columns should be c-1")
		}
		// test edge column shrinkage
		if c1 != 0 {
			res = deleteColumn(res, c1-1)
			_, c2 := res.Dims()
			if c2 != c1-1 {
				t.Error("in subdivision: number of columns should be c-1")
			}
		}
	}
}

func TestDeleteRow(t *testing.T) {
	var tests = []struct{ Mat *mat.Dense }{
		{Mat: mat.NewDense(1, 1, floatTo(1))},
		{Mat: mat.NewDense(2, 2, floatTo(2*2))},
		{Mat: mat.NewDense(5, 5, floatTo(5*5))},
	}
	for _, v := range tests {
		r, c := v.Mat.Dims()
		del := rand.Intn(r)
		res := deleteRow(v.Mat, del)
		r1, c1 := res.Dims()
		if c != c1 && r != 1 {
			t.Error("number of columns should not change unless only one row")
		}
		if r1 != r-1 {
			t.Error("number of rows should be r-1")
		}
		// test edge row shrinkage
		if r1 != 0 {
			res = deleteRow(res, r1-1)
			r2, _ := res.Dims()
			if r2 != r1-1 {
				t.Errorf("in subdivision: number of rows should be r-1. got %d. expect %d", r2, r1-1)
			}
		}
	}
}

func floatTo(i int) []float64 {
	f := make([]float64, i)
	for i := range f {
		f[i] = float64(i + 1)
	}
	return f
}
