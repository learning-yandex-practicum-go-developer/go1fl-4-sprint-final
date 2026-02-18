package spentcalories

import (
	"time"
	"strings"
	"errors"
	"strconv"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, parts[1], 0, errors.New("invalid training data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, parts[1], 0, err
	}
	if steps <= 0 {
		return 0, parts[1], 0, errors.New("steps must be positive")
	}

	dur, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, parts[1], 0, err
	}

	return steps, parts[1], dur, nil
}

func distance(steps int, height float64) float64 {
	lenStep := height * stepLengthCoefficient
	distance := (float64(steps) * lenStep) / mInKm

	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

// RunningSpentCalories calculates calories burned during running.
// 
// It takes the following parameters:
//   steps int       — number of steps taken
//   weight float64  — user's weight in kilograms
//   height float64  — user's height in meters
//   duration time.Duration — running duration
//
// It returns two values:
//   float64 — calories burned during the run
//   error   — non-nil if input parameters are invalid (e.g., non-positive steps, weight, height, or duration) 
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
	}

	if weight <= 0 {
		return 0, errors.New("weight must be positive")
	}

	if height <= 0 {
		return 0, errors.New("height must be positive")
	}

	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	return (weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

// WalkingSpentCalories calculates the number of calories burned during walking.
//
// Parameters:
//   steps int       — number of steps taken
//   weight float64  — user's weight in kilograms
//   height float64  — user's height in meters
//   duration time.Duration — duration of the walk
//
// Returns:
//   float64 — calories burned during the walk
//   error   — non-nil if input parameters are invalid (e.g., non-positive steps, weight, height, or duration)
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
	}

	if weight <= 0 {
		return 0, errors.New("weight must be positive")
	}

	if height <= 0 {
		return 0, errors.New("height must be positive")
	}

	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}


	return ((weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH) * walkingCaloriesCoefficient, nil
}
