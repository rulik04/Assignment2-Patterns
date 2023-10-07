package main

import "fmt"

type Shape interface {
	Draw() string
}

type Circle struct{}

func (c *Circle) Draw() string {
	return "Рисуем круг"
}

type Square struct{}

func (s *Square) Draw() string {
	return "Рисуем квадрат"
}

type Triangle struct{}

func (t *Triangle) Draw() string {
	return "Рисуем треугольник"
}

type ShapeFactory interface {
	CreateShape() Shape
}

type CircleFactory struct{}

func (cf *CircleFactory) CreateShape() Shape {
	return &Circle{}
}

type SquareFactory struct{}

func (sf *SquareFactory) CreateShape() Shape {
	return &Square{}
}

type TriangleFactory struct{}

func (tf *TriangleFactory) CreateShape() Shape {
	return &Triangle{}
}

func DrawShape(factory ShapeFactory) {
	shape := factory.CreateShape()
	fmt.Println(shape.Draw())
}

func main() {
	DrawShape(&CircleFactory{})
	DrawShape(&SquareFactory{})
	DrawShape(&TriangleFactory{})
}
