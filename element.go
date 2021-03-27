// Package kjb for now provides a preliminary
// overview of what an FEA API would look like.
// Nothing is final.
package kjb

import "gonum.org/v1/gonum/mat"

// Node could have essential boundaries attached
// to it for special cases
type Node struct {
	X, Y, Z float64
	Fixity  FixedDof
}

// Rod 2 Node, 2 DOF/Node element (X and Y translation)
type Rod struct {
	Nodes [2]int
}

// Q4 Plane element
type Q4 struct {
	Nodes [4]int
}

// Element preliminary interface sketch
type Element interface {
	K() (mat.Matrix, error)
	DOF() []int
}
