package daysteps

import (
	"time"
	"strings"
	"strconv"
	"errors"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",") // parsed input string
	if len(parts) != 2 { // check len parts
		err := errors.New("[ERROR] Len parts parsed from data less 2")
		return 0, 0, err
	}

	step, err := strconv.Atoi(parts[0]) // convert step in int
	if err != nil {
		return 0, 0, err
	}

	t, err := time.ParseDuration(parts[1]) // parse time duration input in type time.Duration
	if err != nil {
		return 0, 0, err
	}

	return step, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
