package behavior

import "fmt"

type shape interface {
	getType() string
	accept(visitor)
}


type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}




type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
	return "rectangle"
}

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForrectangle(*rectangle)
}


type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) visitForrectangle(s *rectangle) {
	fmt.Println("Calculating area for rectangle")
}

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *middleCoordinates) visitForrectangle(t *rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}