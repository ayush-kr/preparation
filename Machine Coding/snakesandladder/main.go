package main

import (
	"snakesandladder/snakesandladder"
)

func main() {
	readInputFromFile("C:\\Users\\Ayush\\go\\src\\snakesandladder\\data\\input.txt")
	controller := snakesandladder.GetController()
	controller.Start()
}
