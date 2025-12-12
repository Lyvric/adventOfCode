# Advent of Code 2025 Go Solutions

Small Go programs solving Advent of Code style puzzles. Each day lives in its own folder with separate solutions for star 1 and star 2. Inputs are read from a sibling `input.txt` to keep runs reproducible.

## Repository layout

- Day 1: [aoc1/aoc1.1](aoc1/aoc1.1) and [aoc1/aoc1.2](aoc1/aoc1.2)
- Template for new days: [template/star1](template/star1) and [template/star2](template/star2) with a shared [template/input.txt](template/input.txt)

## Running a solution

1. Place your puzzle input in the `input.txt` sitting next to the solution directories (for Day 1 this is [aoc1/input.txt](aoc1/input.txt)).
2. Change into the solution folder, then run:
   - `go run .`
3. The programs print the final position and computed password for the current input.

## Day 1 notes

- Star 1 ([aoc1/aoc1.1/aoc1.go](aoc1/aoc1.1/aoc1.go)): Tracks a cursor on a 0-99 ring, moving left/right by the signed distance on each line and counting how often it lands on position 0.
- Star 2 ([aoc1/aoc1.2/aoc1.2.go](aoc1/aoc1.2/aoc1.2.go)): Extends star 1 by adding wrap counting and handling moves larger than the ring size while preserving direction.

## Using the template for new days

1. Copy `template` to a new day folder, e.g. `cp -r template aoc2`.
2. Fill in the puzzle input at `aoc2/input.txt`.
3. Implement solutions in `aoc2/star1/main.go` and `aoc2/star2/main.go`.

## Tooling

- Go modules inside each star folder declare `module main`; use any modern Go version compatible with your toolchain.
- Dependencies are only from the standard library.
