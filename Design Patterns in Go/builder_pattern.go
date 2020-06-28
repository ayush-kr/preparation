package main

import "log"

type Item interface {
	GetItemName() string
	Packing() Packer
	GetItemPrice() float64
}

type Packer interface {
	Pack() string
}

type wrapper struct {
}

type bottle struct {
}

func (w *wrapper) Pack() string {
	return "wrapper packed"
}

func (b *bottle) Pack() string {
	return "put in bottle"
}

type burger struct {
	price float64
}

func (b *burger) GetItemName() string {
	return "burger"
}

func (b *burger) Packing() Packer {
	return &wrapper{}
}

func (b *burger) GetItemPrice() float64 {
	return b.price
}

type chickenBurger struct {
	price float64
	burger
}

type vegBurger struct {
	price float64
	burger
}

type coldDrink struct {
	price float64
}

func (c *coldDrink) GetItemName() string {
	return "cold Drink"
}

func (c *coldDrink) Packing() Packer {
	return &bottle{}
}

func (c *coldDrink) GetItemPrice() float64 {
	return c.price
}

type Meal struct {
	items []Item
}

func (m *Meal) AddItem(item Item) {
	m.items = append(m.items, item)
}

func (m *Meal) GetCost() float64 {
	cost := 0.0
	for i := 0; i < len(m.items); i++ {
		cost += m.items[i].GetItemPrice()
	}
	return cost
}

func (m *Meal) ShowItems() {
	for i := 0; i < len(m.items); i++ {
		log.Println(m.items[i].GetItemName())
		log.Println(m.items[i].Packing().Pack())
		log.Println("Price:", m.items[i].GetItemPrice())
	}
}

type MealBuilder struct {
}

func (mb *MealBuilder) PrepareNonVegMeal() *Meal {
	meal := &Meal{}
	burg := burger{
		price: 25.0,
	}
	meal.AddItem(&vegBurger{
		0.0,
		burg,
	})
	meal.AddItem(&coldDrink{
		price: 15,
	})
	return meal
}

func (mb *MealBuilder) PrepareVegMeal() *Meal {
	meal := &Meal{}
	burg := burger{
		price: 18.0,
	}
	meal.AddItem(&vegBurger{
		0.0,
		burg,
	})
	meal.AddItem(&coldDrink{
		price: 15,
	})
	return meal
}

func main() {
	mealBuilder := &MealBuilder{}
	nonVegMeal := mealBuilder.PrepareNonVegMeal()
	vegMeal := mealBuilder.PrepareVegMeal()

	nonVegMeal.ShowItems()
	vegMeal.ShowItems()

	log.Println("Non Veg Meal Price: ", nonVegMeal.GetCost())
	log.Println("Veg Meal Price: ", vegMeal.GetCost())
}
