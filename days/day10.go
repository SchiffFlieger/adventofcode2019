package days

import (
	"math"
	"sort"
)

func Day10Part1() int {
	input := day10input()
	points := parseAsteroids(input)
	visibleAsteroids := getAllVisibleAsteroidsFromBestPosition(points)

	return len(visibleAsteroids)
}

func Day10Part2() int {
	input := day10input()
	points := parseAsteroids(input)
	visibleAsteroids := getAllVisibleAsteroidsFromBestPosition(points)

	keys := make([]float64, 0, len(visibleAsteroids))
	for k, v := range visibleAsteroids {
		keys = append(keys, k)
		sort.Sort(v)
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))

	var p point
	destroyed := 0
	for i := 0; destroyed < 200; i++ {
		key := keys[i%len(keys)]
		asteroidsInLine := visibleAsteroids[key]
		if len(asteroidsInLine) > 0 {
			p = asteroidsInLine[0]
			visibleAsteroids[key] = asteroidsInLine[1:]
			destroyed++
		}
	}

	return int(p.x*100 + p.y)
}

type point struct {
	d, x, y float64
}

func (p point) withDistance(d float64) point { return point{d: d, x: p.x, y: p.y} }

type points []point

func (p points) Len() int           { return len(p) }
func (p points) Less(i, j int) bool { return p[i].d < p[j].d }
func (p points) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func getAllVisibleAsteroidsFromBestPosition(ps []point) map[float64]points {
	anglesAtPosition := make(map[float64]points)

	for _, p1 := range ps {
		angles := make(map[float64]points)
		for _, p2 := range ps {
			if p1 != p2 {
				radians := math.Atan2(p2.x-p1.x, p2.y-p1.y)
				distance := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
				angles[radians] = append(angles[radians], p2.withDistance(distance))
			}
		}
		if len(angles) > len(anglesAtPosition) {
			anglesAtPosition = angles
		}
	}

	return anglesAtPosition
}

func parseAsteroids(input []string) []point {
	points := make([]point, 0)
	for y, line := range input {
		for x, char := range line {
			if char == '#' {
				points = append(points, point{x: float64(x), y: float64(y)})
			}
		}
	}

	return points
}

func day10input() []string {
	return []string{
		".#......##.#..#.......#####...#..",
		"...#.....##......###....#.##.....",
		"..#...#....#....#............###.",
		".....#......#.##......#.#..###.#.",
		"#.#..........##.#.#...#.##.#.#.#.",
		"..#.##.#...#.......#..##.......##",
		"..#....#.....#..##.#..####.#.....",
		"#.............#..#.........#.#...",
		"........#.##..#..#..#.#.....#.#..",
		".........#...#..##......###.....#",
		"##.#.###..#..#.#.....#.........#.",
		".#.###.##..##......#####..#..##..",
		".........#.......#.#......#......",
		"..#...#...#...#.#....###.#.......",
		"#..#.#....#...#.......#..#.#.##..",
		"#.....##...#.###..#..#......#..##",
		"...........#...#......#..#....#..",
		"#.#.#......#....#..#.....##....##",
		"..###...#.#.##..#...#.....#...#.#",
		".......#..##.#..#.............##.",
		"..###........##.#................",
		"###.#..#...#......###.#........#.",
		".......#....#.#.#..#..#....#..#..",
		".#...#..#...#......#....#.#..#...",
		"#.#.........#.....#....#.#.#.....",
		".#....#......##.##....#........#.",
		"....#..#..#...#..##.#.#......#.#.",
		"..###.##.#.....#....#.#......#...",
		"#.##...#............#..#.....#..#",
		".#....##....##...#......#........",
		"...#...##...#.......#....##.#....",
		".#....#.#...#.#...##....#..##.#.#",
		".#.#....##.......#.....##.##.#.##",
	}
}
