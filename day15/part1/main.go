package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/math"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
)

const (
	initialHP   = 200
	attackPower = 3
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() *Battle {
	lines := input.Lines()
	lines = lines[:len(lines)-1]

	var goblins, elfs []math.Vector2[int]
	walls := slice.New2D[bool](len(lines), len(lines[0]))
	for r, row := range lines {
		for c, v := range row {
			switch v {
			case '#':
				walls[r][c] = true
			case 'G':
				goblins = append(goblins, math.NewVector2(c, r))
			case 'E':
				elfs = append(elfs, math.NewVector2(c, r))
			}
		}
	}

	battle := NewBattle(walls)
	for _, pos := range goblins {
		battle.AddUnit(Goblin, pos)
	}
	for _, pos := range elfs {
		battle.AddUnit(Elf, pos)
	}

	return battle
}

func process(battle *Battle) int {
	fmt.Printf("Initially:\n%v\n", battle)
	for i := 0; ; i++ {
		if completedRound := battle.PlayRound(); battle.Completed() {
			fullRounds := i
			if completedRound {
				fullRounds++
			}
			totalHP := battle.TotalHP()
			outcome := totalHP * fullRounds
			fmt.Printf("Final round:\n%v\nOutcome: %v * %v = %v\n", battle, fullRounds, totalHP, outcome)
			return outcome
		}
		fmt.Printf("After %v round:\n%v\n", i+1, battle)
	}
}

type UnitType int

const (
	Elf    UnitType = iota
	Goblin UnitType = iota
)

type Unit struct {
	hp       int
	position math.Vector2[int]
	unitType UnitType
}

type Battle struct {
	walls         [][]bool
	unitPositions map[math.Vector2[int]]*Unit
	units         []*Unit
	unitsCount    [2]int
}

func NewBattle(walls [][]bool) *Battle {
	return &Battle{
		walls:         walls,
		unitPositions: map[math.Vector2[int]]*Unit{},
	}
}

func (b *Battle) AddUnit(unitType UnitType, pos math.Vector2[int]) {
	b.unitsCount[unitType]++
	unit := Unit{hp: initialHP, unitType: unitType, position: pos}
	b.units = append(b.units, &unit)
	b.unitPositions[pos] = &unit
}

func (b *Battle) TotalHP() int {
	hp := 0
	for _, unit := range b.units {
		hp += unit.hp
	}

	return hp
}

func (b *Battle) sortUnits() {
	slices.SortFunc(b.units, func(u1, u2 *Unit) int {
		if u1.position.Y != u2.position.Y {
			return u1.position.Y - u2.position.Y
		}

		return u1.position.X - u2.position.X
	})
}

func (b *Battle) PlayRound() (completedRound bool) {
	b.sortUnits()
	diedUnits := make(map[*Unit]bool, len(b.units))
	defer func() {
		if len(diedUnits) > 0 {
			b.units = slices.DeleteFunc(b.units, func(unit *Unit) bool { return diedUnits[unit] })
		}
	}()

	for _, unit := range b.units {
		if diedUnits[unit] {
			continue
		}
		if b.Completed() {
			return false
		}
		if b.attack(unit, diedUnits) {
			continue
		}
		if b.move(unit) {
			b.attack(unit, diedUnits)
		}
	}

	return true
}

func (b *Battle) Completed() bool {
	return b.unitsCount[Elf] == 0 || b.unitsCount[Goblin] == 0
}

func (b *Battle) attack(unit *Unit, diedUnits map[*Unit]bool) bool {
	target := b.findTarget(unit)
	if target == nil {
		return false
	}

	enemy := b.unitPositions[*target]
	enemy.hp -= attackPower
	if enemy.hp <= 0 {
		delete(b.unitPositions, *target)
		diedUnits[enemy] = true
		b.unitsCount[enemy.unitType]--
	}

	return true
}

func (b *Battle) move(unit *Unit) bool {
	moveDirs := b.findMoveDirs(unit)
	if len(moveDirs) == 0 {
		return false
	}

	dir := choosePosition(moveDirs)
	pos := unit.position.Add(dir)
	delete(b.unitPositions, unit.position)
	unit.position = pos
	b.unitPositions[pos] = unit

	return true
}

func (b *Battle) findMoveDirs(unit *Unit) []math.Vector2[int] {
	cells := []math.Vector2[int]{unit.position}
	paths := map[math.Vector2[int]]int{unit.position: 1}
	for len(cells) > 0 {
		var targets []math.Vector2[int]
		nextCells := make([]math.Vector2[int], 0, len(cells))
		for _, cell := range cells {
			dst := paths[cell]
			for _, dir := range dirs {
				pos := cell.Add(dir)
				if paths[pos] > 0 || b.walls[pos.Y][pos.X] {
					continue
				}
				cellUnit, ok := b.unitPositions[pos]
				if ok && cellUnit.unitType == unit.unitType {
					continue
				}
				paths[pos] = dst + 1
				nextCells = append(nextCells, pos)
				if ok {
					targets = append(targets, pos)
				}
			}
		}
		if len(targets) > 0 {
			target := choosePosition(targets)
			positions := backTrace(paths, target)
			moveDirs := make([]math.Vector2[int], len(positions))
			for i, pos := range positions {
				moveDirs[i] = pos.Sub(unit.position)
			}
			return moveDirs
		}
		cells = nextCells
	}

	return nil
}

func backTrace(paths map[math.Vector2[int]]int, from math.Vector2[int]) []math.Vector2[int] {
	cells := map[math.Vector2[int]]bool{from: true}
	for dst := paths[from]; dst > 2; dst-- {
		nextCells := make(map[math.Vector2[int]]bool, len(cells))
		for cell := range cells {
			for _, dir := range dirs {
				pos := cell.Add(dir)
				if paths[pos] == dst-1 {
					nextCells[pos] = true
				}
			}
		}
		cells = nextCells
	}

	return slice.FromSeq(maps.Keys(cells))
}

func (b *Battle) findTarget(unit *Unit) (target *math.Vector2[int]) {
	for _, dir := range dirs {
		pos := unit.position.Add(dir)
		if u, ok := b.unitPositions[pos]; ok && u.unitType != unit.unitType {
			if target == nil || u.hp < b.unitPositions[*target].hp {
				target = &pos
			}
		}
	}

	return
}

func (b *Battle) String() string {
	var sb strings.Builder
	for r, row := range b.walls {
		var hps []string
		for c, w := range row {
			cell := '.'
			if w {
				cell = '#'
			} else if u, ok := b.unitPositions[math.NewVector2(c, r)]; ok {
				cell = 'G'
				if u.unitType == Elf {
					cell = 'E'
				}
				hps = append(hps, fmt.Sprintf("%v(%v)", string(cell), u.hp))
			}
			sb.WriteRune(cell)
		}
		if len(hps) > 0 {
			sb.WriteString("  " + strings.Join(hps, ", "))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func choosePosition(positions []math.Vector2[int]) math.Vector2[int] {
	pos := positions[0]
	for i := 1; i < len(positions); i++ {
		pos2 := positions[i]
		if pos2.Y < pos.Y || (pos2.Y == pos.Y && pos2.X < pos.X) {
			pos = pos2
		}
	}

	return pos
}

// The directions are in reading order
var dirs = []math.Vector2[int]{{Y: -1}, {X: -1}, {X: 1}, {Y: 1}}
