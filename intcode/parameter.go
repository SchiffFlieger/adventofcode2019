package intcode

type (
	Parameter struct {
		Mode         ParameterMode
		rawValue     int
		relativeBase int
	}

	ParameterMode int
)

const (
	ParameterModePosition  = 0
	ParameterModeImmediate = 1
	ParameterModeRelative  = 2
)

func (p *Parameter) Value(input *intslice) int {
	switch p.Mode {
	case ParameterModePosition:
		return input.Get(p.rawValue)
	case ParameterModeImmediate:
		return p.rawValue
	case ParameterModeRelative:
		return input.Get(p.rawValue + p.relativeBase)
	}

	panic("unknown parameter mode")
}

func (p *Parameter) PositionalValue() int {
	return p.rawValue
}
