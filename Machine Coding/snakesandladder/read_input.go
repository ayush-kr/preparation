package main

import (
	"fmt"
	"log"
	"os"
	"snakesandladder/snakesandladder"
)

var (
	snakes  int
	ladders int
	start   int
	end     int
	players int
	name    string
)

var controller snakesandladder.Control

func init() {
	controller = snakesandladder.GetController()
}

func readInputFromFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Println("Error reading file", err.Error())
	}
	defer f.Close()

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = f

	if _, err := fmt.Scanf("%d\n", &snakes); err != nil {
		log.Println("error reading input")
		return
	}

	for i := 0; i < snakes; i++ {
		if _, err := fmt.Scanf("%d %d\n", &start, &end); err != nil {
			log.Println("error reading input")
			return
		}
		controller.AddSnake(start, end)
	}

	if _, err := fmt.Scanf("%d\n", &ladders); err != nil {
		log.Println("error reading input")
		return
	}

	for i := 0; i < ladders; i++ {
		if _, err := fmt.Scanf("%d %d\n", &start, &end); err != nil {
			log.Println("error reading input")
			return
		}
		controller.AddLadder(start, end)
	}

	if _, err := fmt.Scanf("%d\n", &players); err != nil {
		log.Println("error reading input")
		return
	}

	for i := 0; i < players; i++ {
		if _, err := fmt.Scanf("%s\n", &name); err != nil {
			log.Println("error reading input")
			return
		}
		controller.AddPlayer(name)
	}
}
