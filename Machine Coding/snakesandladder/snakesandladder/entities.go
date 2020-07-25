package snakesandladder

import (
	"log"
	"math/rand"
)

type Player struct {
	name string
	id   int
}

func NewPlayer(name string) *Player {
	player := &Player{
		name: name,
	}
	return player
}

type Snake struct {
	startPosition int
	endPosition   int
}

func NewSnake(start, end int) *Snake {
	snake := &Snake{
		startPosition: start,
		endPosition:   end,
	}
	return snake
}

type Ladder struct {
	startPosition int
	endPosition   int
}

func NewLadder(start, end int) *Ladder {
	ladder := &Ladder{
		startPosition: start,
		endPosition:   end,
	}
	return ladder
}

type Board struct {
	snakeMap  map[int]int
	ladderMap map[int]int
	size      int
}

func NewBoard(size int) *Board {
	board := &Board{
		size:      size,
		snakeMap:  make(map[int]int),
		ladderMap: make(map[int]int),
	}

	return board
}

func (b *Board) AddSnake(snake *Snake) {
	b.snakeMap[snake.startPosition] = snake.endPosition
}

func (b *Board) AddLadder(ladder *Ladder) {
	b.ladderMap[ladder.startPosition] = ladder.endPosition
}

type Dice struct {
	maxOutput int
}

func NewDice(size int) *Dice {
	dice := &Dice{
		maxOutput: size,
	}
	return dice
}

type Game struct {
	players   []*Player
	diceList  []*Dice
	board     *Board
	positions map[*Player]int

	diceTotal int
}

func NewGame(board *Board) *Game {
	game := &Game{
		board:     board,
		positions: make(map[*Player]int),
	}
	return game
}

func (g *Game) AddDice(dice *Dice) {
	g.diceList = append(g.diceList, dice)
	g.diceTotal = g.diceTotal + dice.maxOutput
}

func (g *Game) AddPlayer(player *Player) {
	g.players = append(g.players, player)
}

func (g *Game) getDiceValue() int {
	return 1 + rand.Intn(g.diceTotal-1)
}

func (g *Game) getNextPosition(position int) int {
	if val, ok := g.board.ladderMap[position]; ok {
		return val
	}
	if val, ok := g.board.snakeMap[position]; ok {
		return val
	}
	return position
}

func (g *Game) getFinalPosition(diceVal int, player *Player) int {
	currentPosition := g.positions[player]
	finalPosition := currentPosition + diceVal

	if finalPosition > g.board.size {
		return currentPosition
	}

	next := g.getNextPosition(finalPosition)
	for next != finalPosition {
		next = g.getNextPosition(finalPosition)
		finalPosition = next
	}

	return finalPosition
}

func (g *Game) StartGame() {
	activePlayers := len(g.players)
	for {
		for i := 0; i < len(g.players); i++ {
			previousPosition := g.positions[g.players[i]]
			if previousPosition == g.board.size {
				continue
			}
			diceVal := g.getDiceValue()
			finalPosition := g.getFinalPosition(diceVal, g.players[i])
			g.positions[g.players[i]] = finalPosition
			log.Println(g.players[i].name, "rolled a", diceVal, "and moved from", previousPosition, "to", finalPosition)
			if finalPosition == g.board.size {
				log.Println("Oh Bhenchod!!!", g.players[i].name, "WON!!!")
				activePlayers--
			}
			if activePlayers == 1 {
				log.Println("GAME OVER")
				return
			}
		}
	}
}
