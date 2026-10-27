# Asteroids

A fully playable Asteroids clone built with Python and Pygame. Ships shoot, asteroids split, and the game ends when one catches you.

## How to play

| Key | Action |
|---|---|
| `W` | Thrust forward |
| `S` | Thrust backward |
| `A` / `D` | Rotate left / right |
| `Space` | Shoot |

Hit an asteroid to split it into two smaller, faster ones. Asteroids below minimum size are destroyed outright. Get hit by any asteroid and it's game over.

## How it's built

The project uses pygame's sprite system as its backbone. Every object in the game — player, asteroids, and shots — extends a shared `CircleShape` base class that handles position, velocity, and circle-based collision detection.

**Key pieces:**

- `circle_shape.py` — base class with `collides_with()` using distance vs. combined radius
- `player.py` — reads keyboard input each frame, handles rotation, movement, shooting, and a shoot cooldown timer
- `asteroid.py` — moves each frame and splits into two smaller asteroids on hit, each inheriting velocity rotated ±20–50° and scaled up by 1.2x
- `asteroid_field.py` — spawns asteroids at a set rate from random edges of the screen
- `shot.py` — projectile fired from the player's position in the direction they're facing
- `logger.py` — writes game state snapshots and events (splits, shots, hits) to `.jsonl` files for the first 16 seconds of each run

## Setup

```bash
uv sync
uv run main.py
```

## Tech

- Python 3.12
- [Pygame](https://www.pygame.org/) 2.6.1
