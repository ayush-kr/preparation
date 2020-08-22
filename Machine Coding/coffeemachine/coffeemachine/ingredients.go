package coffeemachine

import (
	"errors"
	"sync"
)

//Ingredient iterface declaring methods for an ingredient
type Ingredient interface {
	Refill(quantity int)
	Consume(quantity int, isConsume bool) error
	GetName() string
}

//IngredientItem implementing Ingredient interface
type IngredientItem struct {
	name      string
	quantity  int
	threshold int
	mutex     sync.Mutex
}

//Refill adds some amount of an ingredient
func (i *IngredientItem) Refill(amount int) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.quantity += amount
	if i.quantity >= i.threshold {
		i.quantity = i.threshold
	}
}

//Consume consumes some amount of the ingredient when isConsume flag is true and returns error when unable to consume that amount
func (i *IngredientItem) Consume(amount int, isConsume bool) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.quantity < amount {
		return errors.New(i.name + "'s quantity is less than required")
	}
	if isConsume {
		i.quantity -= amount
	}

	return nil
}

//GetName gives the name of the ingredient
func (i *IngredientItem) GetName() string {
	return i.name
}
