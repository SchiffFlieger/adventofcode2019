package intcode

type intslice struct {
	internal []int
}

func (s *intslice) Get(pos int) int {
	extend(s, pos)
	return s.internal[pos]
}

func (s *intslice) Set(pos, val int) {
	extend(s, pos)
	s.internal[pos] = val
}

func extend(s *intslice, pos int) {
	if pos >= len(s.internal) {
		add := make([]int, pos-len(s.internal)+1)
		s.internal = append(s.internal, add...)
	}
}
