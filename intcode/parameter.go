package intcode

type (
	Parameter struct {
		Mode     ParameterMode
		rawValue int
	}

	ParameterMode int
)

const (
	ParameterModePosition  = 0
	ParameterModeImmediate = 1
)

func (p *Parameter) Value(input []int) int {
	switch p.Mode {
	case ParameterModePosition:
		return input[p.rawValue]
	case ParameterModeImmediate:
		return p.rawValue
	}

	panic("unknown parameter mode")
}

func (p *Parameter) PositionalValue() int {
	return p.rawValue
}
