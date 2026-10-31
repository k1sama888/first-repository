package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
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
		return 0, "", 0, fmt.Errorf("ожидалось 3 элемента,получено: %d", len(parts))
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования шагов %w", err)
	}
	activity := parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования времени: %w", err)
	}
	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	strideLenght := height * stepLengthCoefficient
	distanceKm := (float64(steps) * strideLenght) / float64(mInKm)
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distanceKm := distance(steps, height)
	durationInHours := duration.Hours()
	speed := distanceKm / durationInHours
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	distanceKm := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	var calories float64
	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", nil
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", activity, duration.Hours(), distanceKm, speed, calories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0,получено :%d ", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0, указанный вес: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0, указанный вес: %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCaloriesR := (weight * speed * durationInMinutes) / float64(minInH)
	return spentCaloriesR, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0,получено :%d ", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0, указанный вес: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0, указанный вес: %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}
	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCaloriesW := ((weight * speed * durationInMinutes) / float64(minInH)) * walkingCaloriesCoefficient
	return spentCaloriesW, nil
}
