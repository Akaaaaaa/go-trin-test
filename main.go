package main

// импортируйте нужные пакеты
import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	K1 = 0.035
	K2 = 0.029
)

var (
	Format     = "20060102 15:04:05" // формат даты и времени
	StepLength = 0.65                // длина шага в метрах
	Weight     = 75.0                // вес кг
	Height     = 1.75                // рост м
	Speed      = 1.39                // скорость м/с
)

func parsePackage(data string) (t time.Time, steps int, ok bool) {
	// 1. Разделите строку на две части по запятой в слайс ds
	ds := strings.Split(data, ",")
	if len(ds) != 2 {
		return
	}

	var err error

	t, err = time.Parse(Format, strings.TrimSpace(ds[0]))
	if err != nil {
		return
	}

	steps, err = strconv.Atoi(strings.TrimSpace(ds[1]))
	if err != nil || steps < 0 {
		return
	}
	ok = true
	return
	// 2. Проверьте, чтобы ds состоял из двух элементов
	// 3. Используйте strings.TrimSpace() для очистки пробелов
	// 4. Парсите время с помощью time.Parse()
	// 5. Парсите шаги с помощью strconv.Atoi()
	// 6. Проверьте, что шаги >= 0
	// ...
}

func stepsDay(storage []string) int {
	total := 0
	for _, pkg := range storage {
		_, steps, ok := parsePackage(pkg)
		if ok {
			total += steps
		}
	}
	return total
	// тема оптимизации не затрагивается, поэтому можно
	// использовать parsePackage для каждого элемента списка
	// Не забудьте проверить ok перед добавлением шагов!
	// ...
}

func calories(distance float64) float64 {
	timeMinute := distance / Speed / 60
	burn := (K1*Weight + (Speed*Speed/Height)*K2*Weight) * timeMinute
	// Расчёт времени в минутах: distance / Speed / 60
	// Расчёт калорий по формуле: s
	// ...
	return burn
}

func achievement(distance float64) string {
	if distance >= 6.5 {
		return "Отличный результат! Цель достигнута."
	} else if distance >= 3.9 {
		return "Неплохо! День был продуктивный."
	} else if distance >= 2.0 {
		return "Завтра наверстаем!"
	}
	return "Лежать тоже полезно. Главное — участие, а не победа!"

	//
	// < 2.0 км:
	// ...
}

func showMessage(s string) {
	fmt.Println(s)
	// Выведите сообщение и пустую строку
	fmt.Println()
}

func AcceptPackage(data string, storage []string) []string {
	// 1. Используйте parsePackage для разбора пакета
	//    t, steps, ok := parsePackage(data)
	//    выведите сообщение в случае ошибки
	//    также проверьте количество шагов на равенство нулю
	// ...
	t, steps, ok := parsePackage(data)
	if !ok {
		showMessage("ошибочный формат")
		return storage
	}
	if steps == 0 {
		return storage
	}

	// 2. Получите текущее UTC-время и сравните дни
	//    выведите сообщение, если день в пакете t.Day() не совпадает
	//    с текущим днём
	// ...
	now := time.Now().UTC()

	if t.Day() != now.Day() {
		showMessage("неверный день")
		return storage
	}

	// 3. Проверьте, что время пакета не больше текущего
	// ...
	if t.After(now) {
		showMessage("некорректное значение времени")
		return storage
	}

	// 4. Если есть предыдущие пакеты, проверьте корректность времени
	//    и смену суток
	// ...
	if len(storage) > 0 {
		if data[:8] != storage[len(storage)-1][:8] {
			storage = storage[:0]
		}
	}

	// 5. Добавить пакет в storage
	storage = append(storage, data)
	// 6. Получить общее количество шагов
	totalSteps := stepsDay(storage)
	// 7. Вычислить общее расстояние (в метрах)
	distanceM := float64(totalSteps) * StepLength
	distanceKM := distanceM / 1000
	// 8. Получить потраченные кило калории
	cal := calories(distanceM)
	// 9. Получить мотивирующий текст
	ach := achievement(distanceKM)
	message := fmt.Sprintf("Время: %s\nКоличество шагов за день: %d\n Дистанция составил: %.2f км.\nВы сожгли: %.2f ккал.\n%s", t.Format("15:04:05"), totalSteps, distanceKM, cal, ach)
	showMessage(message)
	// 10. Сформировать и вывести полный текст сообщения
	// 11. Вернуть storage
	// ...

	return storage
}

func main() {
	now := time.Now().UTC()
	today := now.Format("20060102")

	input := []string{
		"01:41:03,-100",
		",3456",
		"12:40:00, 3456 ",
		"something is wrong",
		"02:11:34,678",
		"02:11:34,792",
		"17:01:30,1078",
		"03:25:59,7830",
		"04:00:46,5325",
		"04:45:21,3123",
	}

	var storage []string
	storage = AcceptPackage(today+" 00:11:33,100", storage)
	for _, v := range input {
		storage = AcceptPackage(today+" "+v, storage)
	}
}
