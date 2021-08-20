package main

type Point2 struct {
	X float64
	Y float64
}

func (p Point2) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p *Point2) ScaleBy2(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point2{1, 2}
	//方法变量
	scaleP := p.ScaleBy
	scaleP(2)
	scaleP2 := p.ScaleBy2
	scaleP2(2)

	//方法表达式
	scale := Point2.ScaleBy
	scale(p, 2)
	scale2 := (*Point2).ScaleBy2
	scale2(&p, 2)
}
