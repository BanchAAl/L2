package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Line struct {
	data string
	key  string
}

type Lines []Line

func (l Lines) Len() int {
	return len(l)
}

func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l Lines) Less(i, j int) bool {
	return l[i].key < l[j].key
}

func (l Lines) LessDesc(i, j int) bool {
	return l[i].key > l[j].key
}

func main() {
	numSortColumnFlag := flag.Int("k", 0, "num sort column")
	helpFlag := flag.Bool("help", false, "help output")
	sortByDigitFlag := flag.Bool("n", false, "sort by digit")
	descSortFlag := flag.Bool("r", false, "desc sort")
	delDoubleStrFlag := flag.Bool("u", false, "delete double strings")
	sortByMonthFlag := flag.Bool("M", false, "sort by month")
	endSpaceIgnoreFlag := flag.Bool("b", false, "end space ignore")
	checkSortFlag := flag.Bool("c", false, "check sort")
	sortByDigitSuffixFlag := flag.Bool("h", false, "sort by digit with suffix")

	flag.Parse()

	if *helpFlag {
		fmt.Printf(helpTopic)
		return
	}

	options.numSortColumn = *numSortColumnFlag
	options.sortByDigit = *sortByDigitFlag
	options.descSort = *descSortFlag
	options.delDoubleStr = *delDoubleStrFlag
	options.sortByMonth = *sortByMonthFlag
	options.endSpaceIgnore = *endSpaceIgnoreFlag
	options.checkSort = *checkSortFlag
	options.sortByDigitSuffix = *sortByDigitSuffixFlag

	if (options.sortByDigit && options.sortByMonth) || (options.sortByDigit && options.sortByDigitSuffix) ||
		(options.sortByMonth && options.sortByDigitSuffix) {
		fmt.Printf("\nукажите сортировку только по одному признаку\n")
		os.Exit(1)
	}

	fmt.Printf("\nВведите имя файла:\n")
	var fileName string
	_, err := fmt.Scan(&fileName)
	if err != nil {
		fmt.Printf("\nошибка ввода имени файла: %s\n", err.Error())
		os.Exit(1)
	}

	data, err := createData(fileName)
	if err != nil || len(data) == 0 {
		fmt.Printf("\nошибка открытия/чтения файла %s: %s\n", fileName, err.Error())
	}

	if options.descSort {
		sort.Slice(data, data.LessDesc)
	} else {
		sort.Slice(data, data.Less)
	}

	var output string
	for _, datum := range data {
		output += datum.data + "\n"
	}

	fmt.Println(output)
}

func createData(fileName string) (data Lines, err error) {
	file, errOpen := os.Open(fileName)
	if errOpen != nil {
		return
	}
	defer file.Close()
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	numKey := options.numSortColumn

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if options.endSpaceIgnore {
			strings.TrimRight(line, " ")
		}

		var l Line

		l.data = line
		arr := strings.Split(line, " ")
		if numKey < len(arr) {
			key := arr[numKey]
			if options.sortByDigit {
				if _, errConv := strconv.Atoi(key); errConv != nil {
					key = ""
				}
			}
			l.key = key
		}

		data = append(data, l)
	}

	return
}
