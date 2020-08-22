package coffeemachine

import (
	"errors"
	"log"
	"sync"
	"time"
)

var coffeeMachine *CoffeeMachine

const prepareTime = 5

//BeverageMaker interface declares methods for an item that can make beverages
type BeverageMaker interface {
	OrderBeverage(beverage string)
	ServeBeverage(beverage string)
	AddBeverage(name string, ingredients []string, quantities []int)
	AddIngredient(name string, initialQuantity int, threshold int)
	ReFillIngredient(name string, quantity int)
	ConsumeIngredient(name string, quantity int, isConsume bool) error
	Run()
}

//CoffeeMachine implements BeverageMaker interface
type CoffeeMachine struct {
	ingredients   map[string]Ingredient
	beverages     map[string]Beverage
	outlets       int
	timeToPrepare int

	orderChannel           chan string
	prepareBeverageChannel chan string

	currentOperations int
	mutex             sync.Mutex
}

//GetCoffeeMachine Returns instance of coffee machine
func GetCoffeeMachine() BeverageMaker {
	return coffeeMachine
}

//NewMachine returns new instance of a coffee machine
func NewMachine(outlets int) {
	coffeeMachine = &CoffeeMachine{
		beverages:              make(map[string]Beverage),
		ingredients:            make(map[string]Ingredient),
		outlets:                outlets,
		timeToPrepare:          prepareTime,
		prepareBeverageChannel: make(chan string),
		orderChannel:           make(chan string),
	}
}

//ServeBeverage prepares the beverage and prints error message if preparation is not possible
func (cm *CoffeeMachine) ServeBeverage(beverage string) {
	if val, ok := cm.beverages[beverage]; ok {
		err := val.Prepare()
		if err != nil {
			log.Println("Cannot prepare", beverage, err.Error())
			return
		}
		log.Println("Preparing", beverage, "please wait...")
		time.Sleep((time.Duration)(cm.timeToPrepare) * time.Second)
		log.Println(beverage, "is Ready.", "Enjoy!")
		return
	}
	log.Println("Sorry, Cannot serve", beverage)
}

//AddBeverage adds a new type of beverage making capability in the Coffee Machine, Takes name of beverage, required ingredients and
//their respective quantity as parameters
func (cm *CoffeeMachine) AddBeverage(name string, ingredients []string, quantity []int) {

	ingredientMap := make(map[string]int)
	for i := 0; i < len(ingredients); i++ {
		ingredientMap[ingredients[i]] = quantity[i]
	}

	bevarage := &BeverageItem{
		name:        name,
		ingredients: ingredientMap,
		machine:     cm,
	}

	cm.beverages[name] = bevarage
}

//AddBeverage adds a new type of ingredient in the Coffee Machine, takes name of ingredient, its initial amount and max capacity as parameters
func (cm *CoffeeMachine) AddIngredient(name string, initialQuantity, maxQuantity int) {
	ingredient := &IngredientItem{
		name:      name,
		threshold: maxQuantity,
		quantity:  initialQuantity,
	}

	cm.ingredients[name] = ingredient
}

//ReFillIngredient refills a particular ingredient and increases its amount
func (cm *CoffeeMachine) ReFillIngredient(name string, quantity int) {
	if val, ok := cm.ingredients[name]; ok {
		val.Refill(quantity)
		return
	}
	log.Println("Cannot Refill.", name, "does not exist")
}

//ConsumeIngredient consumes an ingredient by given amount if isConsume flag is true
func (cm *CoffeeMachine) ConsumeIngredient(name string, quantity int, isConsume bool) error {
	if val, ok := cm.ingredients[name]; ok {
		return val.Consume(quantity, isConsume)
	}
	return errors.New(name + " Not Present")
}

//OrderBeverage takes an order and puts that order in queue
func (cm *CoffeeMachine) OrderBeverage(name string) {
	cm.orderChannel <- name
}

//worker thread repeatedly reads from prepareBeverageChannel and performs the task
func (cm *CoffeeMachine) worker() {
	for task := range cm.prepareBeverageChannel {
		cm.ServeBeverage(task)
		cm.mutex.Lock()
		cm.currentOperations--
		cm.mutex.Unlock()
	}
}

//Run spawns the worker threads and reads from orderChannel queue. If current parallel tasks is more than the number of outlets, it doesn't accept the order
//otherwise gives that order to the worker by writing on the prepareBeverageChannel
func (cm *CoffeeMachine) Run() {
	for i := 0; i < cm.outlets; i++ {
		go cm.worker()
	}

	for order := range cm.orderChannel {
		cm.mutex.Lock()
		if cm.currentOperations < cm.outlets {
			cm.currentOperations++
			cm.prepareBeverageChannel <- order
		} else {
			log.Println("Cannot accept the order, Machine is busy")
		}
		cm.mutex.Unlock()
	}
}
