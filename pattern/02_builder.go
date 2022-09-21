package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

type ConcreteBuilder struct {
	product *Product
}

func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>"
}

func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.Content += "<article>" + str + "</article>"
}

func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>"
}

type Product struct {
	Content string
}

func (p *Product) Show() string {
	return p.Content
}
