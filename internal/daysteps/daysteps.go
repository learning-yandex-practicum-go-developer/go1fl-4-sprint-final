package daysteps

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"errors"
	"github.com/Yandex-Practicum/tracker/internal"
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
		return 0, 0, errors.New("len parts parsed from data less 2")
	}

	step, err := strconv.Atoi(parts[0]) // convert step in int
	if err != nil {
		return 0, 0, err
	}
	if step <= 0 { // check quantity step
		return 0, 0, errors.New("quantity step less 0")
	}

	t, err := time.ParseDuration(parts[1]) // parse time duration input in type time.Duration
	if err != nil {
		return 0, 0, err
	}

	return step, t, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	step, timeDuration, err := parsePackage(data) // get result call function parsePackage with param data
	if err != nil {
		fmt.Println(err) // output error in terminal
		return ""
	}
	if step <= 0 { // check quantity step
		return ""
	}

	durationStep := float64(step) * stepLength
	distance := durationStep / mInKm
	calories := internal.WalkingSpentCalories()

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", step, durationStep, calories)
}
