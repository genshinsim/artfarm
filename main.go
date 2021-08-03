package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/genshinsim/artfarm/internal/lib"
)

func main() {
	var main [5]lib.StatType
	main[lib.Feather] = lib.ATK
	main[lib.Flower] = lib.HP
	main[lib.Sand] = lib.ATKP
	main[lib.Goblet] = lib.PyroP
	main[lib.Circlet] = lib.CR
	var desired [lib.EndStatType]float64
	desired[lib.CR] = 0.2

	defer elapsed("artifact farm sim")()
	min, max, mean, sd := sim(10000000, 24, main, desired)
	fmt.Printf("avg: %v, min: %v, max: %v, sd: %v\n", mean, min, max, sd)
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

type result struct {
	count int
	err   error
}

func sim(n, w int, main [lib.EndSlotType]lib.StatType, desired [lib.EndStatType]float64) (min, max int, mean, sd float64) {
	var progress, ss float64
	var sum int
	var data []int
	min = math.MaxInt64
	max = -1
	count := n

	resp := make(chan result, n)
	req := make(chan struct{})
	done := make(chan struct{})
	for i := 0; i < int(w); i++ {
		m := cloneMain(main)
		d := cloneDesired(desired)
		go worker(m, d, req, resp, done)
	}

	go func() {
		var wip int
		for wip < n {
			//try sending a job to req chan while wip < cfg.NumSim
			req <- struct{}{}
			wip++
		}
	}()

	fmt.Print("\tProgress: 0%")

	for count > 0 {
		r := <-resp
		if r.err != nil {
			log.Panicln(r.err)
		}

		data = append(data, r.count)
		count--
		sum += r.count
		if r.count < min {
			min = r.count
		}
		if r.count > max {
			max = r.count
		}

		if (1 - float64(count)/float64(n)) > (progress + 0.1) {
			progress = (1 - float64(count)/float64(n))
			fmt.Printf("...%.0f%%", 100*progress)
		}
	}
	fmt.Printf("\n")
	close(done)

	mean = float64(sum) / float64(n)

	for _, v := range data {
		ss += (float64(v) - mean) * (float64(v) - mean)
	}

	sd = math.Sqrt(ss / float64(n))

	return
}

func cloneDesired(in [lib.EndStatType]float64) (r [lib.EndStatType]float64) {
	for i, v := range in {
		r[i] = v
	}
	return
}

func cloneMain(in [lib.EndSlotType]lib.StatType) (r [lib.EndSlotType]lib.StatType) {
	for i, v := range in {
		r[i] = v
	}
	return
}

func worker(main [lib.EndSlotType]lib.StatType, desired [lib.EndStatType]float64, req chan struct{}, resp chan result, done chan struct{}) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	gen := lib.NewGenerator(r)
	for {
		select {
		case <-req:
			count, err := gen.FarmArtifact(main, desired)
			resp <- result{
				count: count,
				err:   err,
			}
		case <-done:
			return
		}
	}
}
