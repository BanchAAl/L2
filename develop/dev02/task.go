package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpackingString(src string) (unpacking string, err error) {
	var (
		maxIndex    = len(src) - 1
		escapeExist bool
	)

	for i := 0; i < len(src); i++ {
		ch := rune(src[i])
		if unicode.IsDigit(ch) == false {
			var newSymb string
			if ch == '\\' {
				//встретили слэш - берём следующий символ, т.к. слэш экранирует и превращает любой символ в обычный
				if i < maxIndex {
					escapeExist = true
					ch = rune(src[i+1])
					newSymb, i, err = newSymbols(ch, i+1, src)
					if err != nil {
						unpacking = ""
						break
					}
					unpacking += newSymb
					continue
				}
				//слэш в конце отбрасываем
				break
			}
			newSymb, i, err = newSymbols(ch, i, src)
			if err != nil {
				unpacking = ""
				break
			}
			unpacking += newSymb
			continue
		}
		//если сюда дошли, значит две цифры подряд, а это некорректная строка
		err = errors.New("\nнекорректная строка")
		unpacking = ""
	}

	if escapeExist {
		unpacking += " (*)"
	}

	return
}

func newSymbols(ch rune, pos int, str string) (newSymb string, newPos int, err error) {
	if pos < len(str)-1 {
		nextCh := rune(str[pos+1])
		if unicode.IsDigit(nextCh) == true {
			i, errConv := strconv.Atoi(string(nextCh))
			if errConv != nil {
				err = errConv
				return
			}

			for ; i > 0; i-- {
				newSymb += string(ch)
			}
			newPos = pos + 1
		} else {
			newSymb = string(ch)
			newPos = pos
		}
		return
	}
	newSymb = string(ch)
	newPos = pos
	return
}

func main() {
	src := "qwe\\\\5"
	fmt.Printf("\nsource string: %s", src)

	unpacking, err := unpackingString(src)
	if err != nil {
		fmt.Printf("\n%s error unpacking: %s", unpacking, err.Error())
		return
	}

	fmt.Printf("\nUnpaking string: %s", unpacking)
}
