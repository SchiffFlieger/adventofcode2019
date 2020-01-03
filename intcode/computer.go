package intcode

import (
	"fmt"
	"math"
	"os"
)

type (
	Computer struct {
		name string
		data []int
		ptr  int
		in   <-chan int
		out  chan<- int
		done bool
	}

	InstructionPointerChange struct {
		Absolute bool
		Value    int
	}
)

func NewComputer(name string, data []int, in <-chan int, out chan<- int) *Computer {
	dataCopy := append([]int{}, data...)
	dataCopy = extendInputCapacity(dataCopy)
	return &Computer{
		name: name,
		data: dataCopy,
		in:   in,
		out:  out,
	}
}

func (c *Computer) RunUntilDone() {
	for !c.Done() {
		cmd := c.NextCommand()
		c.ApplyCommand(cmd)
	}
}

func (c *Computer) NextCommand() Command {
	if c.done {
		return nil
	}

	op := c.data[c.ptr] % 100
	modes := c.data[c.ptr] / 100
	switch op {
	case 1:
		params := c.getParameters(3, modes)
		return &AddCommand{
			Summand1:       params[0],
			Summand2:       params[1],
			ResultPosition: params[2],
		}
	case 2:
		params := c.getParameters(3, modes)
		return &MultiplyCommand{
			Factor1:        params[0],
			Factor2:        params[1],
			ResultPosition: params[2],
		}
	case 3:
		params := c.getParameters(1, modes)
		return &InputCommand{
			Input:          c.in,
			ResultPosition: params[0],
		}
	case 4:
		params := c.getParameters(1, modes)
		return &OutputCommand{
			Output:      c.out,
			OutputValue: params[0],
		}
	case 5:
		params := c.getParameters(2, modes)
		return &JumpIfTrueCommand{
			Compare: params[0],
			JumpTo:  params[1],
		}
	case 6:
		params := c.getParameters(2, modes)
		return &JumpIfFalseCommand{
			Compare: params[0],
			JumpTo:  params[1],
		}
	case 7:
		params := c.getParameters(3, modes)
		return &LessThanCommand{
			Param1:         params[0],
			Param2:         params[1],
			ResultPosition: params[2],
		}
	case 8:
		params := c.getParameters(3, modes)
		return &EqualsCommand{
			Param1:         params[0],
			Param2:         params[1],
			ResultPosition: params[2],
		}
	case 99:
		return &HaltCommand{}
	}

	panic(fmt.Sprintf("opcode %d unknown", op))
}

func (c *Computer) ApplyCommand(cmd Command) {
	if c.done {
		return
	}

	if os.Getenv("DEBUG") == "1" {
		fmt.Printf("%s: execute %T in line %d\n", c.name, cmd, c.ptr)
	}

	change, err := cmd.Apply(c.data)
	if err != nil {
		c.done = true
		if os.Getenv("DEBUG") == "1" {
			fmt.Printf("%s: %v\n", c.name, err)
		}
		return
	}

	if change.Absolute {
		c.ptr = change.Value
	} else {
		c.ptr += change.Value
	}
}

func (c *Computer) Done() bool {
	return c.done
}

func (c *Computer) GetValue(pos int) int {
	return c.data[pos]
}

func (c *Computer) getParameters(n int, modeMask int) []Parameter {
	params := make([]Parameter, 0, n)
	for i := 0; i < n; i++ {
		mode := (modeMask / int(math.Pow10(i))) % 10
		raw := c.data[c.ptr+1+i]
		params = append(params, Parameter{
			Mode:     ParameterMode(mode),
			rawValue: raw,
		})
	}

	return params
}

func extendInputCapacity(input []int) []int {
	max := 0
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	max = max - len(input) + 1
	if max < 0 {
		return input
	}

	add := make([]int, max, max)
	input = append(input, add...)
	return input
}
