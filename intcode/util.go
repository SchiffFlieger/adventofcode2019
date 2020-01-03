package intcode

import (
	"fmt"
	"io"
	"sync"
)

func IntProviderChan(wg *sync.WaitGroup, ints ...int) <-chan int {
	ch := make(chan int)
	wg.Add(1)

	go func() {
		for _, v := range ints {
			ch <- v
		}
		close(ch)
		wg.Done()
	}()

	return ch
}

func OutputChan(wg *sync.WaitGroup, out io.Writer) chan<- int {
	ch := make(chan int)
	wg.Add(1)

	go func() {
		for val := range ch {
			fmt.Fprintf(out, "out: %d\n", val)
		}
		wg.Done()
	}()

	return ch
}

func ChanSplit(wg *sync.WaitGroup, in <-chan int) (<-chan int, <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)
	wg.Add(1)

	go func() {
		for val := range in {
			out1 <- val
			out2 <- val
		}
		close(out1)
		close(out2)
		wg.Done()
	}()

	return out1, out2
}

func ChanConcatenate(wg *sync.WaitGroup, ins ...<-chan int) <-chan int {
	out := make(chan int)
	wg.Add(1)

	go func() {
		for _, in := range ins {
			for val := range in {
				out <- val
			}
		}
		close(out)
		wg.Done()
	}()

	return out
}

type OutputSlice struct {
	sync.Mutex
	Channel chan int
	done    bool
	slice   []int
}

func (out *OutputSlice) Value() []int {
	out.Lock()
	defer out.Unlock()

	if !out.done {
		panic("called value of intcode.OutputSlice before channel was closed")
	}

	sliceCopy := append([]int{}, out.slice...)
	return sliceCopy
}

func OutputSliceChan(wg *sync.WaitGroup) *OutputSlice {
	out := &OutputSlice{
		Channel: make(chan int),
	}
	wg.Add(1)

	go func() {
		for val := range out.Channel {
			out.slice = append(out.slice, val)
		}

		out.Lock()
		out.done = true
		out.Unlock()

		wg.Done()
	}()

	return out
}
