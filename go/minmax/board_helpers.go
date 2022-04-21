package minmax

import (
	"github.com/BattlesnakeOfficial/rules"
)

var (
	allMoves = []string{rules.MoveUp, rules.MoveDown, rules.MoveRight, rules.MoveLeft}
)

func buildSnakeIds(board *rules.BoardState, youID string, countOfAliveSnakes int) []string {
	snakeIds := make([]string, countOfAliveSnakes)
	writtenCount := 1
	snakeIds[0] = youID
	for _, snake := range board.Snakes {
		if snake.ID != youID {
			snakeIds[writtenCount] = snake.ID
			writtenCount++
		}
	}

	return snakeIds
}

func filterForValidMoves(rulesetName string, board *rules.BoardState, snakeId string, possibleMoves []string) []string {
	// find the snake in the snakes array
	snakeIdx := -1

	for i, s := range board.Snakes {
		if s.ID == snakeId {
			snakeIdx = i
		}
	}
	snake := &board.Snakes[snakeIdx]
	if snake == nil {
		panic("snake not found")
	}
	head := snake.Body[0]
	validMoves := []string{}

	if rulesetName != "wrapped" {
		// eliminate top down left right
		for _, move := range possibleMoves {
			invalid := false
			if move == rules.MoveUp && head.Y == board.Height-1 {
				invalid = true
			}
			if move == rules.MoveDown && head.Y == 0 {
				invalid = true
			}
			if move == rules.MoveLeft && head.X == 0 {
				invalid = true
			}
			if move == rules.MoveRight && head.X == board.Width-1 {
				invalid = true
			}
			if !invalid {
				validMoves = append(validMoves, move)
			}
		}
	} else {
		validMoves = possibleMoves
	}

	possibleMoves = validMoves
	validMoves = []string{}

	// eliminate neck moves
	neck := snake.Body[1]
	for _, move := range possibleMoves {
		invalid := false
		if move == rules.MoveUp && head.Y+1 == neck.Y {
			invalid = true
		}
		if move == rules.MoveDown && head.Y-1 == neck.Y {
			invalid = true
		}
		if move == rules.MoveLeft && head.X-1 == neck.X {
			invalid = true
		}
		if move == rules.MoveRight && head.X+1 == neck.X {
			invalid = true
		}

		if !invalid {
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}

func buildPossibleMoves(rulesetName string, countOfAliveSnakes int, board *rules.BoardState, snakeIds []string, depth int) [][]string {
	possibleMovesForSnakes := make([][]string, countOfAliveSnakes)
	for i := 0; i < countOfAliveSnakes; i++ {
		newAllMoves := make([]string, len(allMoves))
		copy(newAllMoves, allMoves)
		possibleMovesForSnakes[i] = newAllMoves
	}

	for i := 0; i < countOfAliveSnakes; i++ {
		possibleMovesForSnakes[i] = filterForValidMoves(rulesetName, board, snakeIds[i], possibleMovesForSnakes[i])
	}
	return possibleMovesForSnakes
}
