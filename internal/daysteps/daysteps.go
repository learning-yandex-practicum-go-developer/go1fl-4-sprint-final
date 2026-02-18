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

// DayActionInfo parses input data about steps and duration, calculates the distance in kilometers
// and the calories burned, and returns a formatted result string.
//
// Parameters:
//   data string    — input string containing the number of steps and walking duration (e.g., "10000,0h50m")
//   weight float64 — user's weight in kilograms
//   height float64 — user's height in meters
//
// Returns:
//   string — formatted string with step count, distance in km, and calories burned, e.g.:
//            "Количество шагов: 792.\nДистанция составила 0.51 км.\nВы сожгли 221.33 ккал."
func DayActionInfo(data string, weight, height float64) string {
	steps, timeDuration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}

	durationStep := float64(steps) * stepLength
	distance := durationStep / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, timeDuration)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)
}
