package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import (
	"log"
)

type action string

const (
	A action = "A"
	B action = "B"
	C action = "C"
)

type Creator interface {
	CreateProduct(action action) ProductFM
}

type ProductFM interface {
	Use() string
}

type ConcreteCreator struct{}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

func (p *ConcreteCreator) CreateProduct(action action) ProductFM {
	var product ProductFM

	switch action {
	case A:
		product = &ConcreteProductA{string(action)}
	case B:
		product = &ConcreteProductB{string(action)}
	case C:
		product = &ConcreteProductC{string(action)}
	default:
		log.Fatalln("Unknown Action")
	}

	return product
}

type ConcreteProductA struct {
	action string
}

func (p *ConcreteProductA) Use() string {
	return p.action
}

type ConcreteProductB struct {
	action string
}

func (p *ConcreteProductB) Use() string {
	return p.action
}

type ConcreteProductC struct {
	action string
}

func (p *ConcreteProductC) Use() string {
	return p.action
}
