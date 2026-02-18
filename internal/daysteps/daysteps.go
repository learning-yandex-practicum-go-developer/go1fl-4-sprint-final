package daysteps

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"errors"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid training data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errors.New("steps must be positive")
	}

	dur, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}

	return steps, dur, nil
}

// DayActionInfo parses input data, calculates distance in kilometers and burned calories, and returns formatted result string.
func DayActionInfo(data string, weight, height float64) string {
	step, timeDuration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if step <= 0 {
		return ""
	}

	durationStep := float64(step) * stepLength
	distance := durationStep / mInKm
	calories := internal.WalkingSpentCalories() // TODO reaization function

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", step, durationStep, calories)
}
