package intcode

import "io"

type (
	Command interface {
		Apply(input []int) (ip InstructionPointerChange, err error)
	}

	AddCommand struct {
		Summand1       Parameter
		Summand2       Parameter
		ResultPosition Parameter
	}

	MultiplyCommand struct {
		Factor1        Parameter
		Factor2        Parameter
		ResultPosition Parameter
	}

	InputCommand struct {
		Input          <-chan int
		ResultPosition Parameter
	}

	OutputCommand struct {
		Output      chan<- int
		OutputValue Parameter
	}

	JumpIfTrueCommand struct {
		Compare Parameter
		JumpTo  Parameter
	}

	JumpIfFalseCommand struct {
		Compare Parameter
		JumpTo  Parameter
	}

	LessThanCommand struct {
		Param1         Parameter
		Param2         Parameter
		ResultPosition Parameter
	}

	EqualsCommand struct {
		Param1         Parameter
		Param2         Parameter
		ResultPosition Parameter
	}

	HaltCommand struct{}
)

func (c *AddCommand) Apply(input []int) (InstructionPointerChange, error) {
	summand1 := c.Summand1.Value(input)
	summand2 := c.Summand2.Value(input)

	input[c.ResultPosition.PositionalValue()] = summand1 + summand2
	return InstructionPointerChange{Value: 4}, nil
}

func (c *MultiplyCommand) Apply(input []int) (InstructionPointerChange, error) {
	factor1 := c.Factor1.Value(input)
	factor2 := c.Factor2.Value(input)

	input[c.ResultPosition.PositionalValue()] = factor1 * factor2
	return InstructionPointerChange{Value: 4}, nil
}

func (c *InputCommand) Apply(input []int) (InstructionPointerChange, error) {
	input[c.ResultPosition.PositionalValue()] = <-c.Input
	return InstructionPointerChange{Value: 2}, nil
}

func (c *OutputCommand) Apply(input []int) (InstructionPointerChange, error) {
	value := c.OutputValue.Value(input)
	c.Output <- value
	return InstructionPointerChange{Value: 2}, nil
}

func (c *JumpIfTrueCommand) Apply(input []int) (InstructionPointerChange, error) {
	pos := c.JumpTo.Value(input)

	if c.Compare.Value(input) != 0 {
		return InstructionPointerChange{
			Absolute: true,
			Value:    pos,
		}, nil
	}

	return InstructionPointerChange{Value: 3}, nil
}

func (c *JumpIfFalseCommand) Apply(input []int) (InstructionPointerChange, error) {
	pos := c.JumpTo.Value(input)

	if c.Compare.Value(input) == 0 {
		return InstructionPointerChange{
			Absolute: true,
			Value:    pos,
		}, nil
	}

	return InstructionPointerChange{Value: 3}, nil
}

func (c *LessThanCommand) Apply(input []int) (InstructionPointerChange, error) {
	value1 := c.Param1.Value(input)
	value2 := c.Param2.Value(input)

	result := 0
	if value1 < value2 {
		result = 1
	}

	input[c.ResultPosition.PositionalValue()] = result
	return InstructionPointerChange{Value: 4}, nil
}

func (c *EqualsCommand) Apply(input []int) (InstructionPointerChange, error) {
	value1 := c.Param1.Value(input)
	value2 := c.Param2.Value(input)

	result := 0
	if value1 == value2 {
		result = 1
	}

	input[c.ResultPosition.PositionalValue()] = result
	return InstructionPointerChange{Value: 4}, nil
}

func (c *HaltCommand) Apply(input []int) (InstructionPointerChange, error) {
	return InstructionPointerChange{}, io.EOF
}
