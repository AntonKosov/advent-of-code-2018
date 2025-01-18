package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/AntonKosov/advent-of-code-2018/aoc/input"
	"github.com/AntonKosov/advent-of-code-2018/aoc/slice"
	"github.com/AntonKosov/advent-of-code-2018/aoc/transform"
)

func main() {
	answer := process(read())
	fmt.Printf("Answer: %v\n", answer)
}

func read() (immuneGroups, infectionGroups []Group) {
	lines := input.Lines()
	lines = lines[1 : len(lines)-1]
	for lines[0] != "" {
		immuneGroups = append(immuneGroups, ParseGroup(lines[0]))
		lines = lines[1:]
	}

	for _, line := range lines[2:] {
		infectionGroups = append(infectionGroups, ParseGroup(line))
	}

	return
}

func process(immuneGroups, infectionGroups []Group) int {
	for boost := 0; ; boost++ {
		immuneGroupsClone := cloneGroups(immuneGroups)
		infectionGroupsClone := cloneGroups(infectionGroups)
		boostGroups(immuneGroupsClone, boost)
		if win(immuneGroupsClone, infectionGroupsClone) {
			return slice.Sum(immuneGroupsClone, func(g Group) int { return g.units })
		}
	}
}

func boostGroups(groups []Group, boost int) {
	for i := range groups {
		groups[i].damage += boost
	}
}

func cloneGroups(groups []Group) []Group {
	clone := make([]Group, len(groups))
	copy(clone, groups)

	return clone
}

func win(immuneGroups, infectionGroups []Group) bool {
	for len(immuneGroups) > 0 && len(infectionGroups) > 0 {
		immuneTargets := selectTargets(immuneGroups, infectionGroups)
		infectionTargets := selectTargets(infectionGroups, immuneGroups)
		if attack(immuneGroups, infectionGroups, immuneTargets, infectionTargets) == 0 {
			return false
		}

		immuneGroups = removeDefeatedGroups(immuneGroups)
		infectionGroups = removeDefeatedGroups(infectionGroups)
	}

	return len(infectionGroups) == 0
}

func removeDefeatedGroups(groups []Group) []Group {
	return slices.DeleteFunc(groups, func(g Group) bool { return g.units == 0 })
}

func attack(immuneGroups, infectionGroups []Group, immuneTargets, infectionTargets []*int) int {
	type attacker struct {
		group  *Group
		target *Group
	}
	attackers := make([]attacker, 0, len(immuneGroups)+len(infectionGroups))
	addAttackers := func(attackerGroups, defenderGroups []Group, targets []*int) {
		for i := range attackerGroups {
			if idx := targets[i]; idx != nil {
				attackers = append(attackers, attacker{
					group:  &attackerGroups[i],
					target: &defenderGroups[*idx],
				})
			}
		}
	}
	addAttackers(immuneGroups, infectionGroups, immuneTargets)
	addAttackers(infectionGroups, immuneGroups, infectionTargets)

	attackingOrder := newOrderSlice(len(attackers))
	slices.SortFunc(attackingOrder, func(aIdx, bIdx int) int {
		return attackers[bIdx].group.initiative - attackers[aIdx].group.initiative
	})

	killed := 0
	for _, idx := range attackingOrder {
		a := attackers[idx]
		killed += a.group.Attack(a.target)
	}

	return killed
}

func selectTargets(attackingGroups, defendingGroups []Group) (targets []*int) {
	targets = make([]*int, len(attackingGroups))
	attackingOrder := getAttackingOrder(attackingGroups)
	selectedDefenders := make([]bool, len(defendingGroups))
	for _, attackerIdx := range attackingOrder {
		if defenderIdx := findDefendingGroup(attackingGroups[attackerIdx], defendingGroups, selectedDefenders); defenderIdx >= 0 {
			selectedDefenders[defenderIdx] = true
			targets[attackerIdx] = &defenderIdx
		}
	}

	return
}

func newOrderSlice(size int) (order []int) {
	order = make([]int, size)
	for i := range order {
		order[i] = i
	}

	return
}

func findDefendingGroup(attacker Group, defendingGroups []Group, selectedDefenders []bool) (idx int) {
	idx = -1
	for i, defender := range defendingGroups {
		if selectedDefenders[i] {
			continue
		}
		damage := defender.TakenDamage(attacker)
		if damage == 0 {
			continue
		}
		if idx < 0 {
			idx = i
			continue
		}
		bestDefender := defendingGroups[idx]
		bestDefenderTakenDamage := bestDefender.TakenDamage(attacker)
		if bestDefenderTakenDamage > damage {
			continue
		}
		if damage > bestDefenderTakenDamage {
			idx = i
			continue
		}
		if defender.EffectivePower() < bestDefender.EffectivePower() {
			continue
		}
		if defender.EffectivePower() > bestDefender.EffectivePower() {
			idx = i
			continue
		}
		if defender.initiative > bestDefender.initiative {
			idx = i
		}
	}

	return
}

func getAttackingOrder(groups []Group) (attackingOrder []int) {
	attackingOrder = newOrderSlice(len(groups))
	slices.SortFunc(attackingOrder, func(idx1, idx2 int) int {
		a, b := groups[idx1], groups[idx2]
		ap, bp := a.EffectivePower(), b.EffectivePower()
		if ap != bp {
			return bp - ap
		}

		return b.initiative - a.initiative
	})

	return
}

type DamageType string

type Group struct {
	units      int
	unitHP     int
	immuneTo   map[DamageType]bool
	weakTo     map[DamageType]bool
	damage     int
	damageType DamageType
	initiative int
}

func (g Group) EffectivePower() int {
	return g.units * g.damage
}

func (g Group) TakenDamage(attacker Group) int {
	damageType := attacker.damageType
	if g.immuneTo[damageType] {
		return 0
	}

	damage := attacker.EffectivePower()
	if g.weakTo[damageType] {
		damage *= 2
	}

	return damage
}

func (g Group) Attack(defender *Group) int {
	killedUnits := min(defender.units, defender.TakenDamage(g)/defender.unitHP)
	defender.units -= killedUnits
	return killedUnits
}

func ParseGroup(description string) Group {
	nums := transform.StrToInts(description)
	strParts := strings.Split(description, " ")
	immuneTo := map[DamageType]bool{}
	weakTo := map[DamageType]bool{}
	if strings.Contains(description, "(") {
		specs := description[strings.Index(description, "(")+1 : strings.Index(description, ")")]
		for _, spec := range strings.Split(specs, "; ") {
			m := immuneTo
			if strings.HasPrefix(spec, "weak") {
				m = weakTo
			}
			spec = strings.TrimPrefix(spec, "immune to ")
			spec = strings.TrimPrefix(spec, "weak to ")
			for _, damageType := range strings.Split(spec, ", ") {
				m[DamageType(damageType)] = true
			}
		}
	}

	return Group{
		units:      nums[0],
		unitHP:     nums[1],
		immuneTo:   immuneTo,
		weakTo:     weakTo,
		damage:     nums[2],
		damageType: DamageType(strParts[len(strParts)-5]),
		initiative: nums[3],
	}
}
