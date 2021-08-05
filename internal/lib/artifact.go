package lib

import (
	"errors"
	"fmt"
)

func (g *Generator) RandMain(s SlotType) (StatType, error) {

	if s == Flower {
		return HP, nil
	}
	if s == Feather {
		return ATK, nil
	}

	var total, sum float64
	for _, v := range MainStatWeight[s] {
		total += v
	}

	p := g.r.Float64()

	for i, v := range MainStatWeight[s] {
		sum += v / total
		if p <= sum {
			return StatType(i), nil
		}
	}

	return 0, errors.New("error generating new stat - none found")
}

func (g *Generator) RandSubs(m StatType, lvl int) [][]float64 {
	//first element is total, rest is each roll

	var sum, total, p float64
	next := make([]float64, EndStatType)
	subIndex := make([]int, 4)
	result := make([][]float64, EndStatType)
	for i := 0; i < int(EndStatType); i++ {
		result[i] = make([]float64, 1, 6)
	}

	lines := 3
	if g.r.Float64() > .8 {
		lines = 4
		//log.Println("4 lines")
	}

	//total for initial sub
	for i, v := range SubWeights {
		if i == int(m) {
			continue
		}
		total += v
		next[i] = v
	}

	//if artifact lvl is less than 4 AND lines =3, then we only want to roll 3 substats
	n := 4
	if lvl < 4 && lines < 4 {
		n = 3
	}

	var found bool
	for i := 0; i < n; i++ {
		//log.Println(next, sum, total)
		p = g.r.Float64()
		for j, v := range next {
			sum += v / total
			if !found && p <= sum {
				result[j][0] = SubTier[j][g.r.Intn(4)]
				result[j] = append(result[j], result[j][0])
				subIndex[i] = j
				found = true
				//zero out weight for this sub
				next[j] = 0
			}
		}

		//add up total for next run
		sum = 0
		total = 0
		found = false
		for _, v := range next {
			total += v
		}
	}

	up := lvl / 4

	if lines == 3 {
		up--
	}

	for i := 0; i < up; i++ {
		pick := g.r.Intn(4)
		tier := g.r.Intn(4)
		result[subIndex[pick]][0] += SubTier[subIndex[pick]][tier]
		result[subIndex[pick]] = append(result[subIndex[pick]], SubTier[subIndex[pick]][tier])
	}

	return result
}

func (g *Generator) RandSubsNoHist(m StatType, lvl int) []float64 {
	//first element is total, rest is each roll

	var sum, total, p float64
	next := make([]float64, EndStatType)
	subIndex := make([]int, 4)
	result := make([]float64, EndStatType)

	lines := 3
	if g.r.Float64() > .8 {
		lines = 4
		//log.Println("4 lines")
	}

	//total for initial sub
	for i, v := range SubWeights {
		if i == int(m) {
			continue
		}
		total += v
		next[i] = v
	}

	//if artifact lvl is less than 4 AND lines =3, then we only want to roll 3 substats
	n := 4
	if lvl < 4 && lines < 4 {
		n = 3
	}

	var found bool
	for i := 0; i < n; i++ {
		//log.Println(next, sum, total)
		p = g.r.Float64()
		for j, v := range next {
			sum += v / total
			if !found && p <= sum {
				result[j] = SubTier[j][g.r.Intn(4)]
				subIndex[i] = j
				found = true
				//zero out weight for this sub
				next[j] = 0
			}
		}

		//add up total for next run
		sum = 0
		total = 0
		found = false
		for _, v := range next {
			total += v
		}
	}

	up := lvl / 4

	if lines == 3 {
		up--
	}

	for i := 0; i < up; i++ {
		pick := g.r.Intn(4)
		tier := g.r.Intn(4)
		result[subIndex[pick]] += SubTier[subIndex[pick]][tier]
	}

	return result
}

func PrintSubs(in [][]float64) {
	for i, v := range in {
		if v[0] == 0 {
			continue
		}
		fmt.Printf("%v: %.4f, rolls: %v\n", StatTypeString[i], v[0], v[1:])
	}
}

type Artifact struct {
	Slot SlotType
	Main StatType
	Subs []float64
	Ok   bool
}

const maxTries = 1000000000 //1 bil

//FarmArtifact return number of tries it tooks to reach the desired subs
//main is the desired main stat; if main == EndStatType then there's no requirement
func (g *Generator) FarmArtifact(main [EndSlotType]StatType, desired [EndStatType]float64) (int, error) {
	var err error
	var req, score float64
	var done bool
	count := 0
	bag := make([]Artifact, EndSlotType)

	for _, v := range desired {
		if v > 0 {
			req += 1
		}
	}

NEXT:
	for count < maxTries {
		count++
		var a Artifact
		a.Ok = true

		onSet := g.r.Intn(2) == 0
		a.Slot = SlotType(g.r.Intn(5))

		//only goblet can be offset... not quite the right solution but Ok for now?
		if !onSet && a.Slot != Goblet {
			continue NEXT
		}
		a.Main, err = g.RandMain(a.Slot)
		if err != nil {
			return -1, err
		}
		//check main stat, should be equal
		if a.Main != main[a.Slot] {
			continue NEXT
		}
		//generate subs
		a.Subs = g.RandSubsNoHist(a.Main, 20)

		//update bag
		bag, score = update(bag, a, main, desired)
		//done if score > req and all main match
		if score > req {
			done = true
			//check all slots are filled
			for _, v := range bag {
				if !v.Ok {
					done = false
				}
			}
			if done {
				return count, nil
			}
		}
	}
	return -1, errors.New("maximum tries exceeded; requirement not met")
}

func update(bag []Artifact, a Artifact, main [EndSlotType]StatType, desired [EndStatType]float64) ([]Artifact, float64) {
	//score should just be the total distance from desired stats, the lower it is
	//the better it is
	var prev, next, total float64 //score should be % of desired
	var replaced bool
	//if current slot is empty then just put it in
	if !bag[a.Slot].Ok {
		bag[a.Slot] = a
		replaced = true
	}
	for _, x := range bag {
		if x.Ok {
			for i, v := range x.Subs {
				if desired[i] > 0 {
					if desired[i] < v {
						total += 1
					} else {
						total += v / desired[i]
					}
				}
			}
		}
		prev += total

		if !replaced && x.Slot == a.Slot {
			for i, v := range a.Subs {
				if desired[i] > 0 {
					if desired[i] < v {
						next += 1
					} else {
						next += v / desired[i]
					}
				}
			}
		} else {
			next += total
		}
		total = 0
	}

	if prev >= next {
		return bag, prev
	}

	bag[a.Slot] = a

	return bag, next
}
