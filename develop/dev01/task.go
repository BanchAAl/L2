package main

import (
	"fmt"
	"github.com/BanchAAl/L2/develop/dev01/ntptime"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	ntp := ntptime.NewNtpTime("ntp1.stratum2.ru", 60*time.Minute)
	if ntp == nil {
		fmt.Printf("Error create ntp")
		os.Exit(1)
	}

	time, err := ntp.CurrentTime()

	if err != nil {
		fmt.Printf("Error get current time: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Current time: %s", time.String())
}
