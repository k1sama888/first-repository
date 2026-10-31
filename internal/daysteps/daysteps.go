package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
		return 0, 0, fmt.Errorf("ожидалось 2 элемента, получено: %d", len(parts))
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования шагов: %w", err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0, получено %d", steps)
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования времени: %w", err)
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность должна быть больше 0, получено %v", duration)
	}
	return steps, duration, nil
}
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}
	distanceM := (float64(steps) * stepLength)
	distanceKm := distanceM / float64(mInKm)
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanceKm, calories)
}
