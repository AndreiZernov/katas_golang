package sum_of_time_durations

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var WrongTimeErrorMessage = errors.New("given wrong time")

func SumOfTimeDurations(times []string) (string, error) {
	var allSeconds int
	sumTime := make([]string, 3)

	for _, item := range times {
		arrayOfTimes := strings.Split(item, ":")
		if len(arrayOfTimes) == 2 {
			minutes, errMinutes := strconv.Atoi(arrayOfTimes[0])
			seconds, errSeconds := strconv.Atoi(arrayOfTimes[1])
			if errMinutes != nil || errSeconds != nil {
				return "0:0:0", WrongTimeErrorMessage
			}
			allSeconds += minutes*60 + seconds
		}
	}

	sumTime[0] = fmt.Sprint(allSeconds / 3600)
	sumTime[1] = fmt.Sprint(allSeconds / 60 % 60)
	sumTime[2] = fmt.Sprint(allSeconds % 60)

	return strings.Join(sumTime[:], ":"), nil
}
