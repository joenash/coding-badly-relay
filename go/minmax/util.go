package minmax

import "github.com/BattlesnakeOfficial/rules"

func cartersianProductOfStringArrays(arrays ...[]string) [][]string {
	var result [][]string
	if len(arrays) == 1 {
		for _, a := range arrays[0] {
			result = append(result, []string{a})
		}
		return result
	}
	first := arrays[0]
	cartOfTail := cartersianProductOfStringArrays(arrays[1:]...)
	for _, head := range first {
		for _, tail := range cartOfTail {
			result = append(result, append([]string{head}, tail...))
		}
	}

	return result
}

func findYou(board *rules.BoardState, youID string) *rules.Snake {
	for _, snake := range board.Snakes {
		if snake.ID == youID {
			return &snake
		}
	}
	panic("you not found")
}
