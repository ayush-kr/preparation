package snakesandladder

var (
	controller *Controller
)

const (
	defaultBoardSize     = 100
	defaultDiceMaxOutput = 6
)

type Controller struct {
	game *Game
}

type Control interface {
	AddSnake(start, end int)
	AddLadder(start, end int)
	AddPlayer(name string)
	Start()
}

func init() {
	board := NewBoard(defaultBoardSize)
	controller = &Controller{
		game: NewGame(board),
	}
	controller.game.AddDice(NewDice(defaultDiceMaxOutput))
}

func GetController() Control {
	return controller
}

func (c *Controller) AddSnake(start, end int) {
	c.game.board.AddSnake(NewSnake(start, end))
}

func (c *Controller) AddLadder(start, end int) {
	c.game.board.AddLadder(NewLadder(start, end))
}

func (c *Controller) AddPlayer(name string) {
	c.game.AddPlayer(NewPlayer(name))
}

func (c *Controller) Start() {
	c.game.StartGame()
}
