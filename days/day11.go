package days

import (
	"adventofcode/intcode"
	"fmt"
	"math"
	"strings"
	"sync"
)

func Day11Part1() int {
	input := day11input()
	in := make(chan int)
	out := make(chan int)
	c := intcode.NewComputer("default", input, in, out)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		c.RunUntilDone()
		wg.Done()
	}()

	panels := make(map[intPoint]int)
	robot := new(hullPaintingRobot)

	for !c.Done() {
		select {
		case in <- robot.Inspect(panels):
			robot.Paint(panels, <-out)
			if <-out == 0 {
				robot.TurnLeft()
			} else {
				robot.TurnRight()
			}
			robot.Move()
		default:
			break
		}
	}

	wg.Wait()
	return len(panels)
}

func Day11Part2() (image string, w, h int) {
	input := day11input()
	in := make(chan int)
	out := make(chan int)
	c := intcode.NewComputer("default", input, in, out)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		c.RunUntilDone()
		wg.Done()
	}()

	panels := make(map[intPoint]int)
	panels[intPoint{}] = colorWhite
	robot := new(hullPaintingRobot)

	for !c.Done() {
		select {
		case in <- robot.Inspect(panels):
			robot.Paint(panels, <-out)
			if <-out == 0 {
				robot.TurnLeft()
			} else {
				robot.TurnRight()
			}
			robot.Move()
		default:
			break
		}
	}

	wg.Wait()

	maxX, minX := math.MinInt64, math.MaxInt64
	maxY, minY := math.MinInt64, math.MaxInt64

	for key := range panels {
		maxX = maxInt(maxX, key.x)
		minX = minInt(minX, key.x)
		maxY = maxInt(maxY, key.y)
		minY = minInt(minY, key.y)
	}

	width := maxX - minX + 1
	height := maxY - minY + 1
	result := make([]int, 0, width*height)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			color, _ := panels[intPoint{x: x, y: y}]
			result = append(result, color)
		}
	}

	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), ""), "[]"), width, height
}

type hullPaintingRobot struct {
	position  intPoint
	direction int
}

func (r *hullPaintingRobot) TurnLeft() {
	r.direction = (r.direction - 1 + len(directions)) % len(directions)
}
func (r *hullPaintingRobot) TurnRight() { r.direction = (r.direction + 1) % len(directions) }

func (r *hullPaintingRobot) Move() {
	r.position.x += directions[r.direction].x
	r.position.y += directions[r.direction].y
}

func (r *hullPaintingRobot) Paint(panels map[intPoint]int, color int) {
	panels[r.position] = color
}

func (r *hullPaintingRobot) Inspect(panels map[intPoint]int) int {
	if color, ok := panels[r.position]; ok {
		return color
	}

	return colorBlack
}

const (
	colorBlack = 0
	colorWhite = 1
)

type intPoint struct {
	x, y int
}

var directions = []intPoint{
	{x: 0, y: -1}, {x: 1, y: 0}, {x: 0, y: 1}, {x: -1, y: 0},
}

func day11input() []int {
	return []int{3, 8, 1005, 8, 330, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10,
		1008, 8, 0, 10, 4, 10, 102, 1, 8, 29, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 101,
		0, 8, 51, 1, 1103, 2, 10, 1006, 0, 94, 1006, 0, 11, 1, 1106, 13, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4,
		10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 87, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 0, 10, 4, 10,
		1001, 8, 0, 109, 2, 1105, 5, 10, 2, 103, 16, 10, 1, 1103, 12, 10, 2, 105, 2, 10, 3, 8, 102, -1, 8, 10, 1001, 10,
		1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 1001, 8, 0, 146, 1006, 0, 49, 2, 1, 12, 10, 2, 1006, 6, 10, 1, 1101, 4, 10,
		3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 183, 1, 6, 9, 10, 1006, 0, 32,
		3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 213, 2, 1101, 9, 10, 3, 8, 1002,
		8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 239, 1006, 0, 47, 1006, 0, 4, 2, 6, 0, 10,
		1006, 0, 58, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 102, 1, 8, 274, 2, 1005, 14,
		10, 1006, 0, 17, 1, 104, 20, 10, 1006, 0, 28, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4,
		10, 1002, 8, 1, 309, 101, 1, 9, 9, 1007, 9, 928, 10, 1005, 10, 15, 99, 109, 652, 104, 0, 104, 1, 21101, 0,
		937263411860, 1, 21102, 347, 1, 0, 1105, 1, 451, 21101, 932440724376, 0, 1, 21102, 1, 358, 0, 1105, 1, 451,
		3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0,
		104, 0, 3, 10, 104, 0, 104, 1, 21101, 0, 29015167015, 1, 21101, 0, 405, 0, 1106, 0, 451, 21102, 1, 3422723163,
		1, 21101, 0, 416, 0, 1106, 0, 451, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21101, 0, 868389376360, 1,
		21101, 0, 439, 0, 1105, 1, 451, 21102, 825544712960, 1, 1, 21102, 1, 450, 0, 1106, 0, 451, 99, 109, 2, 21201,
		-1, 0, 1, 21101, 0, 40, 2, 21102, 482, 1, 3, 21102, 1, 472, 0, 1106, 0, 515, 109, -2, 2106, 0, 0, 0, 1, 0, 0,
		1, 109, 2, 3, 10, 204, -1, 1001, 477, 478, 493, 4, 0, 1001, 477, 1, 477, 108, 4, 477, 10, 1006, 10, 509, 1101,
		0, 0, 477, 109, -2, 2106, 0, 0, 0, 109, 4, 2101, 0, -1, 514, 1207, -3, 0, 10, 1006, 10, 532, 21102, 1, 0, -3,
		22101, 0, -3, 1, 22102, 1, -2, 2, 21102, 1, 1, 3, 21101, 551, 0, 0, 1106, 0, 556, 109, -4, 2105, 1, 0, 109, 5,
		1207, -3, 1, 10, 1006, 10, 579, 2207, -4, -2, 10, 1006, 10, 579, 22102, 1, -4, -4, 1106, 0, 647, 21201, -4, 0,
		1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21102, 1, 598, 0, 1106, 0, 556, 22101, 0, 1, -4, 21101, 1, 0, -1, 2207,
		-4, -2, 10, 1006, 10, 617, 21102, 0, 1, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 639, 21201, -1, 0, 1,
		21102, 639, 1, 0, 105, 1, 514, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2105, 1, 0}
}
