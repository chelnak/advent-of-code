package main

import (
	"advent-of-code/utils"
	"fmt"
)

// This is a column in the grid used to count marked numbers
type Column struct {
	Count int
}

// This is a row of numbers on our grid
type Row struct {
	Cells []int
	Count int
}

// This is a player in the game
type Player struct {
	Rows          []Row
	Cols          []Column
	Bingo         *bool
	GameOver      *bool
	MarkedNumbers *[]int
}

// This represents the game
type Game struct {
	Players []Player
}

// A player can take a turn
func (p Player) TakeTurn(number int) {

	for rowId, row := range p.Rows {

		for cellId, cell := range row.Cells {

			if number == cell {
				fmt.Println(" -> found on row", rowId, "in cell", cellId)
				p.Rows[rowId].Count++
				p.Cols[cellId].Count++
				*p.MarkedNumbers = append(*p.MarkedNumbers, number)
			}
		}

		if p.Rows[rowId].Count == 5 || p.Cols[rowId].Count == 5 {
			*p.Bingo = true
			break
		}
	}
}

// Calculate our score
func (p Player) CalculateScore(lastNumber int) int {
	score := 0

	for _, row := range p.Rows {
		score += utils.Sum(row.Cells)
	}

	return (score - utils.Sum(*p.MarkedNumbers)) * lastNumber
}

// Play the game
func (g Game) Play(numbers []int) {

	var scoreBoard []map[int]int

	// for each number called
	for _, number := range numbers {

		// A player checks if they have the number
		for idx, player := range g.Players {

			playerId := idx + 1
			fmt.Println("Player", playerId, "looks for ", number, "...")
			player.TakeTurn(number)

			// and also checks if they have won
			if *player.Bingo && !*player.GameOver {

				playerScore := make(map[int]int)
				playerScore[playerId] = player.CalculateScore(number)
				fmt.Println("Player", playerId, "has won")
				scoreBoard = append(scoreBoard, playerScore)

				*player.GameOver = true
			}
		}
	}

	fmt.Println("\nFirst place ->", scoreBoard[0])
	fmt.Println("\nLast place ->", scoreBoard[len(scoreBoard)-1])
}

// Make a new grid from a list of strings
func newPlayer(X int, Y int, data []string) Player {
	player := Player{
		Rows:          make([]Row, X),
		Cols:          make([]Column, Y),
		Bingo:         new(bool),
		GameOver:      new(bool),
		MarkedNumbers: &[]int{},
	}
	for i := 0; i < len(data); i++ {
		player.Rows[i].Cells = utils.SplitStringToInt(data[i], "")
	}

	return player
}

func main() {

	var lines []string
	var players []Player

	width := 5  // X is the number of columns
	height := 5 // y is the number of rows

	utils.Read("inputs.txt", &lines)

	numbers := utils.SplitStringToInt(lines[0], ",")
	blocks := lines[1:]

	for i := 0; i <= len(blocks); i++ {
		if blocks[i] != "" {
			rows := blocks[i : i+height]
			grid := newPlayer(width, height, rows)
			players = append(players, grid)
			i += height
		}
	}

	game := Game{players}
	game.Play(numbers)

}
