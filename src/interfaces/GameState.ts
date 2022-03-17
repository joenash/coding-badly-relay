export interface GameState {
  game: {
    id: string;
    ruleset: {
      name: string;
      version: string;
      settings: Record<string, string | number>;
      squad: Record<string, boolean>;
    };
    source: string;
    timeout: number;
  };
  turn: number;
  board: {
    height: number;
    width: number;
    food: { x: number; y: number }[];
    hazards: { x: number; y: number }[];
    snakes: {
      id: string;
      name: string;
      health: number;
      body: { x: number; y: number }[];
      latency: string;
      head: { x: number; y: number };
      length: number;
      shout?: string;
      squad: string;
      customizations: {
        color: string;
        head: string;
        tail: string;
      };
    }[];
  };
  you: {
    id: string;
    name: string;
    health: number;
    body: { x: number; y: number }[];
    latency: string;
    head: { x: number; y: number };
    length: number;
    shout?: string;
    squad: string;
    customizations: {
      color: string;
      head: string;
      tail: string;
    };
  };
}
