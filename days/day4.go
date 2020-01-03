package days

import "strconv"

func Day4Part1() int {
	low, high := day4input()
	return countMatches(low, high, increasingOrder, hasDoubleDigits)
}

func Day4Part2() int {
	low, high := day4input()
	return countMatches(low, high, increasingOrder, hasAdjacentDigits)
}

func countMatches(low, high int, matchFuncs ...func(string) bool) int {
	count := 0

outer:
	for i := low; i <= high; i++ {
		for _, match := range matchFuncs {
			if !match(strconv.Itoa(i)) {
				continue outer
			}
		}
		count++
	}

	return count
}

func increasingOrder(a string) bool {
	a1, _ := strconv.Atoi(string(a[0]))
	a2, _ := strconv.Atoi(string(a[1]))
	a3, _ := strconv.Atoi(string(a[2]))
	a4, _ := strconv.Atoi(string(a[3]))
	a5, _ := strconv.Atoi(string(a[4]))
	a6, _ := strconv.Atoi(string(a[5]))
	return a1 <= a2 && a2 <= a3 &&
		a3 <= a4 && a4 <= a5 && a5 <= a6
}

func hasDoubleDigits(a string) bool {
	return a[0] == a[1] || a[1] == a[2] || a[2] == a[3] || a[3] == a[4] || a[4] == a[5]
}

func hasAdjacentDigits(a string) bool {
	return (a[0] == a[1] && a[1] != a[2]) || // 00xxxx
		(a[0] != a[1] && a[1] == a[2] && a[2] != a[3]) || // x00xxx
		(a[1] != a[2] && a[2] == a[3] && a[3] != a[4]) || // xx00xx
		(a[2] != a[3] && a[3] == a[4] && a[4] != a[5]) || // xxx00x
		(a[3] != a[4] && a[4] == a[5]) // xxxx00
}

func day4input() (low int, high int) {
	return 367479, 893698
}
