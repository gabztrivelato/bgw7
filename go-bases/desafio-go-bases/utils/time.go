package utils

import (
	"strconv"
	"strings"
)

func GetHour(time string) (int, error) {
	splitedTime := strings.Split(time, ":")

	hour, err := strconv.Atoi(splitedTime[0])
	if err != nil {
		return 0, err
	}

	return hour, nil
}
