package days

import (
	"adventofcode/intcode"
	"math"
	"runtime"
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

		initChan := intcode.IntProviderChan(wg, phaseSettings[0], 0)
		abWriteChan := make(chan int)
		abReadChan := intcode.ChanConcatenate(wg, intcode.IntProviderChan(wg, phaseSettings[1]), abWriteChan)
		bcWriteChan := make(chan int)
		bcReadChan := intcode.ChanConcatenate(wg, intcode.IntProviderChan(wg, phaseSettings[2]), bcWriteChan)
		cdWriteChan := make(chan int)
		cdReadChan := intcode.ChanConcatenate(wg, intcode.IntProviderChan(wg, phaseSettings[3]), cdWriteChan)
		deWriteChan := make(chan int)
		deReadChan := intcode.ChanConcatenate(wg, intcode.IntProviderChan(wg, phaseSettings[4]), deWriteChan)
		resultChan := make(chan int)

		computers := []*intcode.Computer{
			intcode.NewComputer("AMP_A", input, initChan, abWriteChan),
			intcode.NewComputer("AMP_B", input, abReadChan, bcWriteChan),
			intcode.NewComputer("AMP_C", input, bcReadChan, cdWriteChan),
			intcode.NewComputer("AMP_D", input, cdReadChan, deWriteChan),
			intcode.NewComputer("AMP_E", input, deReadChan, resultChan),
		}

		for _, v := range computers {
			wg.Add(1)
			go func(c *intcode.Computer) {
				c.RunUntilDone()
				wg.Done()
			}(v)
		}

		max = maxInt(max, <-resultChan)
		wg.Wait()
	}

	return max
}

func Day7Part2() int {
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

		eaWriteChan := make(chan int)
		eaReadChan := intcode.ChanConcatenate(wg, intcode.IntProviderChan(wg, phaseSettings[0], 0), eaWriteChan)

		computers := []*intcode.Computer{
			intcode.NewComputer("AMP_A", input, eaReadChan, abWriteChan),
			intcode.NewComputer("AMP_B", input, abReadChan, bcWriteChan),
			intcode.NewComputer("AMP_C", input, bcReadChan, cdWriteChan),
			intcode.NewComputer("AMP_D", input, cdReadChan, deWriteChan),
			intcode.NewComputer("AMP_E", input, deReadChan, eaWriteChan),
		}

		for _, v := range computers {
			wg.Add(1)
			go func(c *intcode.Computer) {
				c.RunUntilDone()
				wg.Done()
			}(v)
		}

		for !computers[0].Done() {
			runtime.Gosched()
		}
		max = maxInt(max, <-eaReadChan)

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
