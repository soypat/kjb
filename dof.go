package kjb

//FixedDof is a bitmask to fix global degrees of freedom for a model
type FixedDof uint8

// Bitmasks to fix model degrees of freedom
const (
	Xtranslation FixedDof = 1 << iota
	Ytranslation
	Ztranslation
	Xrotation
	Yrotation
	Zrotation
	NumberDof int = iota
)

// Common fixities
const (
	XYZFree  = Xrotation | Yrotation | Zrotation
	XYFree   = XYZFree | Ztranslation
	Beam3Dof = XYFree &^ Zrotation
	AllFree  = 0
	AllFixed = 0xff
)

// AreFixed asks if all dofs are fixed by the reciever
func (d FixedDof) AreFixed(dofs FixedDof) bool { return dofs&^d == 0 }

// AreFree asks if all free dofs are also free in the reciever
func (d FixedDof) AreFree(dofs FixedDof) bool { return ^dofs&d == 0 }
