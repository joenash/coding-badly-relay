package minmax

import "github.com/BattlesnakeOfficial/rules"

func heuristic(board *rules.BoardState, youID string) float32 {
	// find the snake that is you
	you := findYou(board, youID)

	score := float32(0.0)

	// if we're dead, worst possible score
	if you.EliminatedCause != "" {
		return -1.0
	}

	// if you are the ONLY snake return 0.8
	if len(board.Snakes) > 1 {
		longest_snake_bonus := float32(0.2)
		// If you are NOT the longest snake return 0.0
		for _, snake := range board.Snakes {
			if snake.ID != youID {
				if len(snake.Body) > len(you.Body) {
					longest_snake_bonus = 0.0
				}
			}
		}
		score += longest_snake_bonus
	}

	// If you recently ate add an arbitrary value to the score
	if you.Health > 90 {
		score += 0.5
	}

	return score

	// TODO: more heuristics here
}
