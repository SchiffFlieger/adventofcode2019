package days

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1Part1(t *testing.T) {
	assert.Equal(t, 3335787, Day1Part1())
}

func TestDay1Part2(t *testing.T) {
	assert.Equal(t, 5000812, Day1Part2())
}

func TestDay2Part1(t *testing.T) {
	assert.Equal(t, 3850704, Day2Part1())
}

func TestDay2Part2(t *testing.T) {
	assert.Equal(t, 6718, Day2Part2())
}

func TestDay3Part1(t *testing.T) {
	assert.Equal(t, 1431, Day3Part1())
}

func TestDay3Part2(t *testing.T) {
	assert.Equal(t, 48012, Day3Part2())
}

func TestDay4Part1(t *testing.T) {
	assert.Equal(t, 495, Day4Part1())
}

func TestDay4Part2(t *testing.T) {
	assert.Equal(t, 305, Day4Part2())
}

func TestDay5Part1(t *testing.T) {
	assert.Equal(t, 15259545, Day5Part1())
}

func TestDay5Part2(t *testing.T) {
	assert.Equal(t, 7616021, Day5Part2())
}

func TestDay6Part1(t *testing.T) {
	assert.Equal(t, 144909, Day6Part1())
}

func TestDay6Part2(t *testing.T) {
	assert.Equal(t, 259, Day6Part2())
}

func TestDay7Part1(t *testing.T) {
	assert.Equal(t, 206580, Day7Part1())
}

func TestDay7Part2(t *testing.T) {
	assert.Equal(t, 2299406, Day7Part2())
}

func TestDay8Part1(t *testing.T) {
	assert.Equal(t, 2760, Day8Part1())
}

func TestDay8Part2(t *testing.T) {
	// the correct answer is AGUEB
	// format the string with w=25 and h=6 and remove all 0s to see the letters
	expected := "011000110010010111101110010010100101001010000100101001010000100101110011100111101011010010100001001010010100101001010000100101001001110011001111011100"
	assert.Equal(t, expected, Day8Part2())
}

func TestDay9Part1(t *testing.T) {
	assert.Equal(t, 3429606717, Day9Part1())
}

func TestDay9Part2(t *testing.T) {
	assert.Equal(t, 33679, Day9Part2())
}

func TestDay10Part1(t *testing.T) {
	assert.Equal(t, 256, Day10Part1())
}

func TestDay10Part2(t *testing.T) {
	assert.Equal(t, 1707, Day10Part2())
}
