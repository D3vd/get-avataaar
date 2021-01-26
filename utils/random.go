package utils

import (
	"math/rand"
	"time"
)

// GetRandom : Gets a random element from list
func GetRandom(list []string) string {
	l := len(list)

	rand.Seed(time.Now().Unix())
	n := rand.Int() % l

	return list[n]
}
