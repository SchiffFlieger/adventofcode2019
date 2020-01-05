package days

import (
	"bytes"
	"fmt"
	"math/big"
)

func Day12Part1() int {
	moons := day12input()

	for step := 0; step < 1000; step++ {
		for i, m := range moons {
			param := make([]*moon, 3)
			copy(param[:i], moons[:i])
			copy(param[i:], moons[i+1:])
			m.ApplyGravity(param)
		}

		for _, moon := range moons {
			moon.ApplyVelocity()
		}
	}

	energy := 0
	for _, moon := range moons {
		energy += moon.Energy()
	}

	return energy
}

func Day12Part2() string {
	moons := day12input()

	stepsX, stepsY, stepsZ := make(map[string]int), make(map[string]int), make(map[string]int)
	px, py, pz := -1, -1, -1

	xValues := func(m *moon) (int, int) { return m.position.x, m.velocity.x }
	yValues := func(m *moon) (int, int) { return m.position.y, m.velocity.y }
	zValues := func(m *moon) (int, int) { return m.position.z, m.velocity.z }

	for step := 0; step < 1000000000; step++ {
		px = maxInt(px, addTo(stepsX, toString(moons, xValues), step))
		py = maxInt(py, addTo(stepsY, toString(moons, yValues), step))
		pz = maxInt(pz, addTo(stepsZ, toString(moons, zValues), step))

		if px > 0 && py > 0 && pz > 0 {
			break
		}

		for i, m := range moons {
			param := make([]*moon, 3)
			copy(param[:i], moons[:i])
			copy(param[i:], moons[i+1:])
			m.ApplyGravity(param)
		}

		for _, moon := range moons {
			moon.ApplyVelocity()
		}
	}

	bigPx := *big.NewInt(int64(px))
	bigPy := *big.NewInt(int64(py))
	bigPz := *big.NewInt(int64(pz))

	result := lcm(bigPx, lcm(bigPy, bigPz))
	return result.String()
}

func lcm(m, n big.Int) big.Int {
	var z big.Int
	return *z.Mul(z.Div(&m, z.GCD(nil, nil, &m, &n)), &n)
}

type vector3d struct {
	x, y, z int
}

type moon struct {
	position, velocity vector3d
}

func (m *moon) ApplyGravity(others []*moon) {
	for _, other := range others {
		m.velocity.x += sgnInt(other.position.x - m.position.x)
		m.velocity.y += sgnInt(other.position.y - m.position.y)
		m.velocity.z += sgnInt(other.position.z - m.position.z)
	}
}

func (m *moon) ApplyVelocity() {
	m.position.x += m.velocity.x
	m.position.y += m.velocity.y
	m.position.z += m.velocity.z
}

func (m *moon) Energy() int {
	potential := absInt(m.position.x) + absInt(m.position.y) + absInt(m.position.z)
	kinetic := absInt(m.velocity.x) + absInt(m.velocity.y) + absInt(m.velocity.z)
	return potential * kinetic
}

func toString(moons []*moon, fn func(*moon) (int, int)) string {
	b := new(bytes.Buffer)
	for _, m := range moons {
		pos, vel := fn(m)
		fmt.Fprintf(b, "(%d, %d)|", pos, vel)
	}
	return b.String()
}

func addTo(steps map[string]int, key string, step int) int {
	if val, ok := steps[key]; ok {
		return step - val
	}
	steps[key] = step
	return -1
}

func day12input() []*moon {
	return []*moon{
		{position: vector3d{x: -14, y: -4, z: -11}},
		{position: vector3d{x: -9, y: 6, z: -7}},
		{position: vector3d{x: 4, y: 1, z: 4}},
		{position: vector3d{x: 2, y: -14, z: -9}},
	}
}
