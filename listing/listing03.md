Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
вывод программы:
nil
false

err не nil т.к. тип содержит, а само значение nil, печатается значение, а сравниваеся сам интерфейс.
С точки зрения внутреннего устройства интерфейса: поля data структры интерфейса cодержит тип, который nil, но само это
поле не nil.

```
