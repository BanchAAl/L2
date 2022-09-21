package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	data := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	a := anagrams(data)
	fmt.Println(a)
}

func anagrams(data []string) (anagrams map[string][]string) {
	anagrams = make(map[string][]string)
	for _, datum := range data {
		datum = strings.ToLower(datum)
		runes := []rune(datum)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		ok, key := findAnagrams(runes, anagrams)
		if ok {
			anagrams[key] = append(anagrams[key], datum)
			break
		}
		anagrams[datum] = []string{datum}
	}

	for k, v := range anagrams {
		if len(v) < 2 {
			delete(anagrams, k)
		}
	}

	return
}

func findAnagrams(checked []rune, anagrams map[string][]string) (ok bool, key string) {
	for k, _ := range anagrams {
		runes := []rune(k)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		if reflect.DeepEqual(runes, checked) == true {
			ok = true
			key = k
			break
		}
	}

	return
}
