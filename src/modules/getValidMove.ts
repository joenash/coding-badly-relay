import { GameState } from "../interfaces/GameState";
import { Move } from "../interfaces/Move";

export const getValidMove = (gameState: GameState, snakeId: string): Move => {
  const targetSnake = gameState.board.snakes.find(
    (snake) => snake.id === snakeId
  );
  if (!targetSnake) {
    console.error(`Invalid ID passed: ${snakeId}`);
    return "right";
  }

  const snakePos = targetSnake.head;
  const top = 0;
  const left = 0;
  const bottom = gameState.board.height - 1;
  const right = gameState.board.width - 1;

  const possibleMoves: Move[] = ["up", "down", "left", "right"];

  if (snakePos.y === top) {
    possibleMoves.splice(possibleMoves.indexOf("up"), 1);
  }
  if (snakePos.y === bottom) {
    possibleMoves.splice(possibleMoves.indexOf("down"), 1);
  }
  if (snakePos.x === left) {
    possibleMoves.splice(possibleMoves.indexOf("left"), 1);
  }
  if (snakePos.x === right) {
    possibleMoves.splice(possibleMoves.indexOf("right"), 1);
  }

  return possibleMoves[Math.floor(Math.random() * possibleMoves.length)];
};
