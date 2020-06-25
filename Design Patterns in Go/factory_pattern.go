package main

import (
	"errors"
	"log"
)

//Shaper defines shape interface
type Shaper interface {
	Draw()
}

type rectangle struct {
	width  int
	length int
}

type circle struct {
	radius int
}

func (r *rectangle) Draw() {
	log.Println("length:", r.length, "width:", r.width)
}

func (c *circle) Draw() {
	log.Println("radius:", c.radius)
}

//ShapeFactory is a factory for shapes
type ShapeFactory struct {
}

//GetShape returns instance of shape
func (s *ShapeFactory) GetShape(name string) (Shaper, error) {
	if name == "RECTANGLE" {
		return &rectangle{
			length: 5,
			width:  10,
		}, nil
	} else if name == "CIRCLE" {
		return &circle{
			radius: 5,
		}, nil
	}
	return nil, errors.New("Invalid Shape")
}

func main() {
	shapeFactory := &ShapeFactory{}
	shapeRectangle, err := shapeFactory.GetShape("RECTANGLE")
	if err != nil {
		log.Println("Error in getting RECTANGLE shape:", err.Error())
	} else {
		shapeRectangle.Draw()
	}

	shapeCircle, err := shapeFactory.GetShape("CIRCLE")
	if err != nil {
		log.Println("Error in getting CIRCLE shape:", err.Error())
	} else {
		shapeCircle.Draw()
	}
}
