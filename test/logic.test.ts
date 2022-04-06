const wasm = require("../src/wasm");

import { GameState } from "../src/interfaces/GameState";

import { info, move } from "../src/logic";

function createGameState(myBattlesnake: GameState["you"]): GameState {
  return {
    game: {
      id: "",
      ruleset: { name: "", version: "", settings: {}, squad: {} },
      source: "",
      timeout: 0,
    },
    turn: 0,
    board: {
      height: 0,
      width: 0,
      food: [],
      snakes: [myBattlesnake],
      hazards: [],
    },
    you: myBattlesnake,
  };
}

function createBattlesnake(
  id: string,
  bodyCoords: { x: number; y: number }[]
): GameState["you"] {
  return {
    id: id,
    name: id,
    health: 0,
    body: bodyCoords,
    latency: "",
    head: bodyCoords[0],
    length: bodyCoords.length,
    shout: "",
    squad: "",
    customizations: {
      color: "#F22F46",
      head: "do-sammy",
      tail: "do-sammy",
    },
  };
}

beforeAll(async () => {
  await wasm.run();
});

describe("Battlesnake API Version", () => {
  test("should be api version 1", () => {
    const result = info();
    expect(result.apiversion).toBe("1");
  });
});

describe("Battlesnake Moves", () => {
  test("should never move into its own neck", () => {
    // Arrange
    const me = createBattlesnake("me", [
      { x: 2, y: 0 },
      { x: 1, y: 0 },
      { x: 0, y: 0 },
    ]);
    const gameState = createGameState(me);

    // Act 1,000x (this isn't a great way to test, but it's okay for starting out)
    for (let i = 0; i < 1000; i++) {
      const moveResponse = move(gameState);
      // In this state, we should NEVER move left.
      const allowedMoves = ["up", "down", "right"];
      expect(allowedMoves).toContain(moveResponse.move);
    }
  });
});
