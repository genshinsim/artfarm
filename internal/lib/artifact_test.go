package lib

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestGenerator_RandSubs(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := NewGenerator(r)
	res := g.RandSubs(ATK, 20)
	PrintSubs(res)
}

func TestGenerator_FarmArtifact(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := NewGenerator(r)
	var main [5]StatType
	main[Feather] = ATK
	main[Flower] = HP
	main[Sand] = ATKP
	main[Goblet] = PyroP
	main[Circlet] = CR
	var desired [EndStatType]float64
	desired[CR] = 0.2
	count, err := g.FarmArtifact(main, desired)
	if err != nil {
		t.Error(err)
	}
	log.Println(count)
}

var result [][]float64

func BenchmarkGenerator_RandSubs(b *testing.B) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := NewGenerator(r)
	var res [][]float64
	for i := 0; i < b.N; i++ {
		res = g.RandSubs(ATK, 20)
	}
	result = res
}
