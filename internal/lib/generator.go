package lib

import "math/rand"

type Generator struct {
	r *rand.Rand
}

func NewGenerator(r *rand.Rand) *Generator {
	return &Generator{r: r}
}