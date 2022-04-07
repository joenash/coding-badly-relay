import { GameState } from "../interfaces/GameState";
import { Move } from "../interfaces/Move";

export const getValidMove = (gameState: GameState, snakeId: string): String[] => {
  const targetSnake = gameState.board.snakes.find(
    (snake) => snake.id === snakeId
  );
  if (!targetSnake) {
    console.error(`Invalid ID passed: ${snakeId}`);
    return ["right"];
  }

  const snakePos = targetSnake.head;
  // neck is body[1]
  const neck = targetSnake.body[1];
  const top = gameState.board.height - 1;
  const left = 0;
  const bottom = 0;
  const right = gameState.board.width - 1;

  const possibleMoves: Move[] = ["up", "down", "left", "right"];

  if (neck.x - 1 === snakePos.x && neck.y === snakePos.y) {
    possibleMoves.splice(possibleMoves.indexOf("left"), 1);
  }

  if (neck.x + 1 === snakePos.x && neck.y === snakePos.y) {
    possibleMoves.splice(possibleMoves.indexOf("right"), 1);
  }

  if (neck.x === snakePos.x && neck.y + 1 === snakePos.y) {
    possibleMoves.splice(possibleMoves.indexOf("up"), 1);
  }

  if (neck.x === snakePos.x && neck.y - 1 === snakePos.y) {
    possibleMoves.splice(possibleMoves.indexOf("down"), 1);
  }



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

  return possibleMoves
};
