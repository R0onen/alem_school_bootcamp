# Project title: Game of Life

## Description

This project is a simulation of Game of Life implemented using the Go programming language, which is a cellular automaton devised by the British mathematician John Horton Conway in 1970. It is a zero-player game, meaning that its evolution is determined by its initial state, requiring no further input.

## Rules

The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells, each of which is in one of two possible states, live(populated) or dead (unpopulated). Every cell interacts with its eight neighbors, which are the cells that are horizontally, vertically, or diagonally adjacent. At each step in time (or generation), the following transitions occur:

- **Underpopulation**: Any live cell with fewer than two live neighbors dies.
- **Survival**: Any live cell with two or three live neighbors lives on to the next generation.
- **Overpopulation**: Any live cell with more than three live neighbors dies (overpopulation).
- **Reproduction**: Any dead cell with exactly three live neighbors becomes a live cell (reproduction).

The grid evolves according to these rules, which govern the state of each cell in every generation.

## Installation

1. Clone the repository:
   git clone git@git.platform.alem.school:aashirbe/crunch03.git
2. Run the application:
   go run main.go

## Usage

The program supports several flags:

- **--help** : Show the help message and exit.
- **--verbose** : Display detailed information about the simulation, including grid size, number of ticks, speed, and map name.
- **--delay-ms=X** : Set the animation speed in milliseconds. Default is 2500 milliseconds.
- **--file=X** : Load the initial grid from a specified file.
- **--edges-portal** : Enable portal edges where cells that exit the grid appear on the opposite side.
- **--random=WxH** : Generate a random grid of the specified width (W) and height (H).
- **--fullscreen** : Adjust the grid to fit the terminal size with empty cells.
- **--footprints** : Add traces of visited cells, displayed as '∘'.
- **--colored** : Add color to live cells and traces if footprints are enabled.

To run the program with specified flag: go run main.go [option]
