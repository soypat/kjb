package kjb

import "testing"

func TestIsFixed(t *testing.T) {
	var tests = []struct {
		ModelDof, IsFixedCheck FixedDof
		Result                 bool
	}{
		{ModelDof: AllFree, IsFixedCheck: AllFree, Result: true},
		{ModelDof: AllFixed, IsFixedCheck: AllFixed, Result: true},
		{ModelDof: Beam3Dof, IsFixedCheck: Xtranslation, Result: false},
		{ModelDof: XYFree, IsFixedCheck: Xtranslation, Result: false},
		{ModelDof: XYFree, IsFixedCheck: Ztranslation, Result: true},
		{ModelDof: XYFree, IsFixedCheck: Xtranslation | Ytranslation, Result: false},
		{ModelDof: XYFree, IsFixedCheck: XYZFree, Result: true},
	}
	for i := range tests {
		got := tests[i].ModelDof.AreFixed(tests[i].IsFixedCheck)
		expect := tests[i].Result
		if got != expect {
			t.Errorf("are fixed fail %v: got %t, expect %t. case %d", tests[i], got, expect, i)
		}
	}
}

func TestIsFree(t *testing.T) {
	var tests = []struct {
		ModelDof, IsFreeCheck FixedDof
		Result                bool
	}{
		{ModelDof: Beam3Dof, IsFreeCheck: XYFree, Result: true},
		{ModelDof: Zrotation, IsFreeCheck: XYFree, Result: true},
		{ModelDof: Xtranslation, IsFreeCheck: Xtranslation, Result: true},
		{ModelDof: AllFree, IsFreeCheck: AllFree, Result: true},
		{ModelDof: AllFixed, IsFreeCheck: AllFixed, Result: true},
		// false conditions
		{ModelDof: Xtranslation, IsFreeCheck: AllFree, Result: false},
		{ModelDof: Xtranslation | Ytranslation, IsFreeCheck: Ztranslation | Ytranslation, Result: false},
	}
	for i := range tests {
		got := tests[i].ModelDof.AreFree(tests[i].IsFreeCheck)
		expect := tests[i].Result
		if got != expect {
			t.Errorf("are free fail %v: got %t, expect %t. case %d", tests[i], got, expect, i)
		}
	}
}
