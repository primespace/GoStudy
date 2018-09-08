package main

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	println(area)
	area = rect.area2()
	println(area)
}
