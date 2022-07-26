package die_roll

import (
	"math/rand"
	"time"
)

func DieRoll(dieSides, times int) []int {
	rand.Seed(time.Now().UnixNano())
	var (
		minNum, maxNum = 1, dieSides + 1
		list           = make([]int, dieSides)
	)

	for 0 < times {
		random := rand.Intn(maxNum-minNum) + minNum
		list[random-1] = list[random-1] + 1
		times--
	}
	return list
}
