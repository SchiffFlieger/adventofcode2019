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

	panic("value: unknown parameter mode")
}

func (p *Parameter) PositionalValue() int {
	switch p.Mode {
	case ParameterModePosition:
		return p.rawValue
	case ParameterModeRelative:
		return p.rawValue + p.relativeBase
	}

	panic("positional value: unknown parameter mode")
}
