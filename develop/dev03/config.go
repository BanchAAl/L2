package main

const (
	helpTopic = "ОПИСАНИЕ\n\n\tСортировка строк из файла. Отображение в стандартный вывод\n\nИСПОЛЬЗОВАНИЕ\n\n\twbsort [ОПЦИИ]... [ИМЯ_ФАЙЛА]\n\nОПЦИИ\n\n" +
		"\t-k — указание колонки для сортировки. default - 0\n\t-n — сортировать по числовому значению\n\t-r — сортировать в обратном порядке\n" +
		"\t-u — не выводить повторяющиеся строки\n\t-M — сортировать по названию месяца\n\t-b — игнорировать хвостовые пробелы\n\t-c — проверять отсортированы ли данные\n" +
		"\t-h — сортировать по числовому значению с учётом суффиксов\n"
)

var options struct {
	numSortColumn     int
	sortByDigit       bool
	descSort          bool
	delDoubleStr      bool
	sortByMonth       bool
	endSpaceIgnore    bool
	checkSort         bool
	sortByDigitSuffix bool
}
