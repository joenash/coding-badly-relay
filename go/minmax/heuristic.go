package minmax

import "github.com/BattlesnakeOfficial/rules"

func heuristic(board *rules.BoardState, youID string) float32 {
	// find the snake that is you
	you := findYou(board, youID)

	// if we're dead, worst possible score
	if you.EliminatedCause != "" {
		return -1.0
	}

	// if you are the longest snake return 0.8
	if len(board.Snakes) == 1 {
		return 0.8
	}
	for _, snake := range board.Snakes {
		if snake.ID != youID {
			if len(snake.Body) > len(you.Body) {
				return 0.0
			}
		}
	}
	return 0.8

	// TODO: more heuristics here
}
