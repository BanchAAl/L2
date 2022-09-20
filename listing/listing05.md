Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
вывод программы:
error

т.к. err интерфейсный тип и при проверке на nil, по сути, происходит проверка содержит ли err реальное значение, т.е 
"реализованный тип"

```
