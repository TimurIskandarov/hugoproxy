package counter

import (
	"fmt"
	"time"
)


func GetCounterPage(tick time.Time, counter byte) string {
	return fmt.Sprintf(
		content,
		tick.Local().Format("15:04:05"),
		counter,
	)
}

const content = `
---
menu:
  before:
    name: tasks
    weight: 5
title: Обновление данных в реальном времени
---

# Задача: Обновление данных в реальном времени

Напишите воркер, который будет обновлять данные в реальном времени, на текущей странице.
Текст данной задачи менять нельзя, только время и счетчик.

Файл данной страницы: `+"`/app/static/tasks/_index.md`"+`

Должен меняться счетчик и время:

Текущее время: %v

Счетчик: %v
`
