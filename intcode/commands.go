package intcode

import "io"

type (
	Command interface {
		Apply(input *intslice) (ip CommandFeedback, err error)
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

	AdjustRelativeBaseCommand struct {
		Param1 Parameter
	}

	HaltCommand struct{}
)

func (c *AddCommand) Apply(input *intslice) (CommandFeedback, error) {
	summand1 := c.Summand1.Value(input)
	summand2 := c.Summand2.Value(input)

	input.Set(c.ResultPosition.PositionalValue(), summand1+summand2)
	return CommandFeedback{InstructionPointerChange: InstructionPointerChange{Value: 4}}, nil
}

func (c *MultiplyCommand) Apply(input *intslice) (CommandFeedback, error) {
	factor1 := c.Factor1.Value(input)
	factor2 := c.Factor2.Value(input)

	input.Set(c.ResultPosition.PositionalValue(), factor1*factor2)
	return CommandFeedback{InstructionPointerChange: InstructionPointerChange{Value: 4}}, nil
}

func (c *InputCommand) Apply(input *intslice) (CommandFeedback, error) {
	input.Set(c.ResultPosition.PositionalValue(), <-c.Input)
	return CommandFeedback{InstructionPointerChange: InstructionPointerChange{Value: 2}}, nil
}

func (c *OutputCommand) Apply(input *intslice) (CommandFeedback, error) {
	value := c.OutputValue.Value(input)
	c.Output <- value
	return CommandFeedback{InstructionPointerChange: InstructionPointerChange{Value: 2}}, nil
}

func (c *JumpIfTrueCommand) Apply(input *intslice) (CommandFeedback, error) {
	pos := c.JumpTo.Value(input)

	if c.Compare.Value(input) != 0 {
		return CommandFeedback{
			InstructionPointerChange: InstructionPointerChange{
				Absolute: true,
				Value:    pos,
			},
		}, nil
	}

	return CommandFeedback{InstructionPointerChange: InstructionPointerChange{Value: 3}}, nil
}

func (c *JumpIfFalseCommand) Apply(input *intslice) (CommandFeedback, error) {
	pos := c.JumpTo.Value(input)

	if c.Compare.Value(input) == 0 {
		return CommandFeedback{
			InstructionPointerChange: InstructionPointerChange{
				Absolute: true,
				Value:    pos,
			},
		}, nil
	}

	return CommandFeedback{InstructionPointerChange: InstructionPointerChange{Value: 3}}, nil
}

func (c *LessThanCommand) Apply(input *intslice) (CommandFeedback, error) {
	value1 := c.Param1.Value(input)
	value2 := c.Param2.Value(input)

	result := 0
	if value1 < value2 {
		result = 1
	}

	input.Set(c.ResultPosition.PositionalValue(), result)
	return CommandFeedback{InstructionPointerChange: InstructionPointerChange{Value: 4}}, nil
}

func (c *EqualsCommand) Apply(input *intslice) (CommandFeedback, error) {
	value1 := c.Param1.Value(input)
	value2 := c.Param2.Value(input)

	result := 0
	if value1 == value2 {
		result = 1
	}

	input.Set(c.ResultPosition.PositionalValue(), result)
	return CommandFeedback{InstructionPointerChange: InstructionPointerChange{Value: 4}}, nil
}

func (c *AdjustRelativeBaseCommand) Apply(input *intslice) (CommandFeedback, error) {
	base := c.Param1.Value(input)
	return CommandFeedback{
		InstructionPointerChange: InstructionPointerChange{Value: 2},
		RelativeBaseChange:       RelativeBaseChange{Value: base},
	}, nil
}

func (c *HaltCommand) Apply(_ *intslice) (CommandFeedback, error) {
	return CommandFeedback{}, io.EOF
}
