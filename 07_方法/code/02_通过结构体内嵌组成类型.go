package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X float64
	Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p ColoredPoint) Distance(q Point) {

}

func main() {
	var cp ColoredPoint
	cp.X = 1
	cp.Point.Y = 2
	fmt.Println(cp)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	p.ScaleBy(2)
	p.Distance(q.Point)
	//p.Distance(q) //编译错误
	fmt.Println(p)
}
