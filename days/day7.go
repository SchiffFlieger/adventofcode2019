package days

import (
	"adventofcode/intcode"
	"math"
	"sync"

	"github.com/gitchander/permutation"
)

func Day7Part1() int {
	input := day7input()
	phaseSettings := []int{0, 1, 2, 3, 4}
	perm := permutation.New(permutation.IntSlice(phaseSettings))

	max := math.MinInt64
	for perm.Next() {
		wg := &sync.WaitGroup{}

		channel := make(chan int, 1) // buffered
		c1 := intcode.NewComputer("AMP_A", input, intcode.IntProviderChan(wg, phaseSettings[0], 0), channel)
		c1.RunUntilDone()

		c2 := intcode.NewComputer("AMP_B", input, intcode.IntProviderChan(wg, phaseSettings[1], <-channel), channel)
		c2.RunUntilDone()

		c3 := intcode.NewComputer("AMP_C", input, intcode.IntProviderChan(wg, phaseSettings[2], <-channel), channel)
		c3.RunUntilDone()

		c4 := intcode.NewComputer("AMP_D", input, intcode.IntProviderChan(wg, phaseSettings[3], <-channel), channel)
		c4.RunUntilDone()

		c5 := intcode.NewComputer("AMP_E", input, intcode.IntProviderChan(wg, phaseSettings[4], <-channel), channel)
		c5.RunUntilDone()

		wg.Wait()
		max = maxInt(max, <-channel)
	}

	return max
}

func Day7Part2() int {
	type pair struct {
		computer *intcode.Computer
		out      chan int
	}

	input := day7input()
	phaseSettings := []int{5, 6, 7, 8, 9}
	perm := permutation.New(permutation.IntSlice(phaseSettings))

	max := math.MinInt64
	for perm.Next() {
		wg := &sync.WaitGroup{}

		abWriteChan := make(chan int)
		abReadChan := intcode.ChanConcatenate(wg, intcode.IntProviderChan(wg, phaseSettings[1]), abWriteChan)
		bcWriteChan := make(chan int)
		bcReadChan := intcode.ChanConcatenate(wg, intcode.IntProviderChan(wg, phaseSettings[2]), bcWriteChan)
		cdWriteChan := make(chan int)
		cdReadChan := intcode.ChanConcatenate(wg, intcode.IntProviderChan(wg, phaseSettings[3]), cdWriteChan)
		deWriteChan := make(chan int)
		deReadChan := intcode.ChanConcatenate(wg, intcode.IntProviderChan(wg, phaseSettings[4]), deWriteChan)

		initChan := intcode.IntProviderChan(wg, phaseSettings[0], 0)
		eaWriteChan := make(chan int)
		resultChan, eaReadChan := intcode.ChanSplit(wg, eaWriteChan)
		eaReadChan = intcode.ChanConcatenate(wg, initChan, eaReadChan)

		computers := []pair{
			{computer: intcode.NewComputer("AMP_A", input, eaReadChan, abWriteChan), out: abWriteChan},
			{computer: intcode.NewComputer("AMP_B", input, abReadChan, bcWriteChan), out: bcWriteChan},
			{computer: intcode.NewComputer("AMP_C", input, bcReadChan, cdWriteChan), out: cdWriteChan},
			{computer: intcode.NewComputer("AMP_D", input, cdReadChan, deWriteChan), out: deWriteChan},
			{computer: intcode.NewComputer("AMP_E", input, deReadChan, eaWriteChan), out: eaWriteChan},
		}

		for _, v := range computers {
			wg.Add(1)
			go func() {
				v.computer.RunUntilDone()
				close(v.out)
				wg.Done()
			}()
		}

		go func() {
			for val := range resultChan {
				max = maxInt(max, val)
			}
		}()
		wg.Wait()
	}

	return max
}

func day7input() []int {
	return []int{3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 38, 63, 72, 81, 106, 187, 268, 349, 430, 99999, 3, 9, 101, 5,
		9, 9, 1002, 9, 3, 9, 101, 3, 9, 9, 4, 9, 99, 3, 9, 102, 3, 9, 9, 101, 4, 9, 9, 1002, 9, 2, 9, 1001, 9, 2, 9,
		1002, 9, 4, 9, 4, 9, 99, 3, 9, 1001, 9, 3, 9, 4, 9, 99, 3, 9, 102, 5, 9, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 1001,
		9, 2, 9, 1002, 9, 5, 9, 1001, 9, 2, 9, 102, 3, 9, 9, 4, 9, 99, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9,
		4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9,
		3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3,
		9, 1001, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9,
		1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101,
		1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002,
		9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9,
		9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2,
		9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9,
		4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9,
		3, 9, 1002, 9, 2, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3,
		9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9,
		1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99}
}
