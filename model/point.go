package model

import "fmt"

type Point struct {
	X, Y int
}

type Size struct {
	Width, Height int
}

func (point Point) String() string {
	return fmt.Sprintf("(%d,%d)", point.X, point.Y)
}

func (size Size) String() string {
	return fmt.Sprintf("[%d,%d]", size.Width, size.Height)
}

func (point Point) Up(bounds *Size) *Point {
	if bounds == nil || point.Y > 0 {
		return &Point{point.X, point.Y - 1}
	} else {
		return nil
	}
}

func (point Point) Down(bounds *Size) *Point {
	y := point.Y + 1
	if bounds == nil || y < bounds.Height {
		return &Point{point.X, y}
	} else {
		return nil
	}
}

func (point Point) Left(bounds *Size) *Point {
	if bounds == nil || point.X > 0 {
		return &Point{point.X - 1, point.Y}
	} else {
		return nil
	}
}

func (point Point) Right(bounds *Size) *Point {
	x := point.X + 1
	if bounds == nil || x < bounds.Width {
		return &Point{x, point.Y}
	} else {
		return nil
	}
}

func (p1 Point) Add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func (point Point) In(size Size) bool {
	return point.X >= 0 && point.Y >= 0 && point.X < size.Width && point.Y < size.Height
}
