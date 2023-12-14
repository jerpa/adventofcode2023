package main

import (
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day14/part1"
	"github.com/jerpa/adventofcode2023/day14/part2"
	"time"
)

func main() {
	start := time.Now()

	c.Print("Part1: ", part1.Part1(c.GetData(Data)))
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	c.Print("Part2: ", part2.Part2(c.GetData(Data)))
	c.Print("Took: ", time.Since(start).String())
}

var Data = `
#...#...O....OO.#O.O........#.OO.#.O..O.O.#.#..O#..O...#.#......#..O.....O.###O..OO.#.....O.O.......
O.#..#..O.O...O..#OO#...#.#.OO....O.#.....O.#OOO.#O.#O#O.......O......O.OO.#.#.........O..#.....#...
....O..O..O#...#O#......O...##O#.O.....#O..##..O...#O..O.O..#.O..O..#.....OO....O.O.O...O......#.O.O
##O....O....#.OOO...O.#....O#.O#OO....O....O#...O..O#O.#..O.OO...O..O..#.#O..OO.O...O..#.#OO#.......
.O.O#OO..#...#..........O#.O..##.O..OO.O...#........O##.OOOO..#.#....O..O#O.#O.##..OO..O...#..O##...
OO...O##.....#...O.#O.....#.OOO.#.##...##...##.O.....##....O..O..#..#..OO...#..........OO..#.O#...O.
.OO#..#.O.O##...O....O...O....#.O.O..#.OO........O..O..O..###..O...O#..#..#......OO.O..#O....O##.#..
#O.O..O#...O..O.O.O...#.....O..O......O....OO.OO..O#.O#..#.O......#..O...#.O.##.#....O.OO.#...O...O#
.#.O.....O...O...O...O.#.........OO.....#................O....O..#..........##.O#.OO..O###..#.......
#O...#O..O#.......O...#..O...OO..#...O........#....O#..OO...O.O..#.OOO.....O..O.#...O.OO#.OO#.OO.O#.
......#.O...O...O.O...#...#O.O#....#.O..###..O....#.#.O.OO#....O#.....O...#.O......##......#O..#OO..
#OO.##...#..O....#...#...OO...........OO.O.OOOO...#...#O..#...#O#.O.O......O....O..#.O..O...#O#.O.O.
#...O.OO...#.#.OO...O#O....OO....#O....O.O#..#....O...##..O..O.#..#..#...........O...OO#.#..O.....O#
..#..O.O#.....O....O#..#........OO#.#...#O.O....O#....O#...O..#....OO#.O.#........#O....O.O...##O.#O
....#O.....O#...O.......O.......OO..#.....#...O##.O.......#O#....O...O..O.O.......#.....#..#.O.O.O..
.....O.....O.O.O#O.O##.....OOO#...#O.O#O.....#.O#..........O#O.....O.OOO..##.#O....O.O.....#.O.O#O.#
O#O....O.#.......O.O.OO.#..#...#O#..O...O.......O.#........#O#....O....O.#...#.O.#O...O.#.........O.
........O.O...O.#O....#..O....#O.....OO#..O.O..####..#...O..O...OOO.#OO..#....#...OO..O.....#O..O..#
O#...##...#..#O#..#O#..O#....OO..O#O#.....OOOOOO#.#.O.O..O.OO.O........O...#.#......OO#...#..#.#...#
..OO.O..OO.OO....#........#O.#...#O....#...O.OO...O#..O......O.###....#O......#.....####...#..#.#OO.
..O.O..O#....#.O.#.#.#.O#..O.......OO.........##.O.....#OOO..O#..#OO#.....O...OO.O.O...#...OO#.O.#..
O...#........###..##..O#.#O.O..O.#OO.OO...#....#O.O#O..OO...##....O......OO..#...#.......#...#......
..O..#.O#...OO#O.O#....#.....O#..O..O#..O.O..O.#O.O.#O.#OO..#........#.#O..#O....O..#...#OO#...#.O.O
..#O..O.....#....O.O...O.#O.#.O#.O......#........OO...O....#.O........O#.OO#..........O.O..OOOOO#.#.
O..##.##...........O##.#...#.O.#.OO..O.O...OO#.#..#...O..O....#..#..#...O###......O#.#O#..O.#....O.#
#.O.O.O.###.##......#.OO...OO.O..OOO.O.OOO.....##..O#...O......O.#O..#...#...OO..O#......O....##.O.#
#...#.O.O....O..#...#........##O#O#O..##O.#......O...OO#O...OOOO#O..OO##..#......O.#O.#.O#.OO...O..#
#..#........#O.#O.#...O...O....#OO...O....#...O..#......#..O..O...##........O..#.O..#.#OO..OO.O##.O.
..#...#...OO.#.O#O.O.##..#..##..O#...O#....#..O....OO.OOO..O#.#....O.......OO.#.O#.....O.#O.O.#.....
.......O...O..O...#.O..OO...O.....O.O#..#OO#.O...O..O#O..O..#.OO.......O..O....OOOO..O.#.O#O#..O.##.
.####...O.....#...#......#.O...........O......O.#O.....O.#......#.#...O...#......#..O#.....O.O..O..#
O....O...#OO##.O.O...O....O..O......O.O..O..O..O#O..O.O#O.#O.#..OO.....#.O.#O#..##.#.O.O...OO....#..
....#.O..O..#...#....O............O.O.#O.#.O...O....O.#.OO....O..O##...O##......O....O...OO........#
.#O...O..O.O.#O.......O###OO.O.O.O.....##O..OO...#.O..#OO..#...#...#.O.O###.....#O#..##.#O.....O....
..#.###.#..##......#..#.#OO..O......O..#OOO.#O...#....OO.##O...O..#.#O....#OO....O.#.......#....O...
#OO...#......#..#........#.OO.##....#...OOOO..O#O..O.#..#O........#O.OO#..........OO..OO....O..#O.#.
..##.##.O....#.##O..#...###..O##.#.#......O#O..###..##......O...#.....##.#...........O.O.O....O.#...
...O.#..#......#.OO..#..O##...O..O...OO##...OO.O#....O.O.#.O.....O..#..O..#...O..#O.....O.O..O.O#.##
O......###OOO#.#.....OO...O.....O.##.....OO..........OO#O#.O.#.#O....O..#O.#..........##O...#O....O.
..#...#...OO.OO#.....OO...O.O.O.##.##O##..O.#.O#..O.O.....OOOO....#....O.O.#..##OO.O#..#O.O.#...O...
.#.OO.##....O..O...O..#.##.O#OO..#.....OO.........O.OOO.#..O....O....OOO.#.OO..O..#...O#OO.OO##.#.O.
.#.OO.O#.O.O.....#O..O.O..O......O.#O...O..O...##.O#..#.O.....O...O....OO.OOO..OO.O...OO..#.#..OOO#.
.#..O.#.O..#.OOO....#O#O#............##.#.O...O....O.OO..#.O....O#O#.O.#...#....#..OO.O.O#.#.....O.O
O....O.O#....O..#.......O..#O.##.#...O..O.......O#.#.....#..#..O......O..O.....#............##OO....
O#.O...O.##O..OO..O.O.O.#O.O..##..O#.#..#O..O#O...#...OO..#.O.O....OO..O..#..O.......O.O#....O..#.O#
#.O..##......#..O#O..........##O#.#O..O#.#..#.O..##O.O.#.OO....O#..........O..#.##..#..O.#...#O.OO..
.O..#....OO...O.......#...#...O.##...O..O.O.........OO#O.OO.O..O...O.##.....##.....##O..#......O....
#..O.O..OO....O.O#.#O#........O..OO#.OO#O.O..O..#......O#...##.......#..##..O.#..#.O.O...#....#.O...
#..O.#...OO#.OO#.O.OO.O##.#....OO....#........#.O.O....O.#..O#..#....#.O#..O#..O.....OO..OOOO.....O.
.O..OO#.O#OO..#O..O.#......O..O..O.#O.#..#.#.O..O.O..O.....##.OO..#O.#............O#...#......O.OO.#
O......##OO....O...O......#...O..O.#...OO#O..O.OOO..O.....O...#.##..OOO.##...O.O....O#O....##...#...
..OO..#.O.O..#O.O...O...#...O.O#.##O....##...###O..O...O.O##..........###...O...O..#...#.#.....O....
#..OO....O..#O..OO...OO.............O.O..OOO.O.#.....O#O....#O.....#.O..O....O..O....O..O....O.#....
OO.O##..O...#O........#O...#OOOO.O..O....O...O..O..O#.O.O..##.##..O.#O...O#...#..O...O...OO.......O.
#..O#..O...O.#......#OO.OO.O.#...O...O#...#...O..##O#....#......#.....#..#O...#.O.O..............O#.
##..OO#..........#.O..OOOOO..#...#....#.#..O#O...O....#.O.#O#OO.#......#...O........O...O#....#....O
.....O...O#....O..O..OO.O..#...O.#..#..#...O#..#O....O.#.....#O....O...##....#..#......###.O.....O#.
#...O#..O..O#..OO....#.#...#O#..#.....#..#..................OO....O......OO.........O###O#....OO.O..
O#...O.O.O..##.....##....#....#......#...O.#...O..##...O..O...##..#....#..O..O#...##..O.O..OOO..#...
..#.......#.O..O..#O....#...#.....OOO#...O#...O#.#.#O.....O#O.#.O.#.O......O.#..O.O..O#......O.O..O#
...O.#.#..O.#..###..#...#.....O..O.....OO#..O..#.O..#....OOO...........#.......#..OO..OO#.#O....#O..
.O.O...#....#.#O..#..OO#..#....OO....#O..O#.#..#..O.#..O.#O#..OO........#O.#...OO....OO......#.#O...
OO.O.O#..#.#.#.O.OOO...#....O..O##...O.OO#.#.....O.....#.....O....#.........O...O..........OOO.O.OOO
......O...O...###....##.O.O..OOO..O.O......O.#.......#..O..........O#..#..#O.#.##.....###.#O#O#O..#.
O..#..#O.OO.......#.OO...#OO..OO.O#O.....#.#.O..OOO.OO.##O.O#...#.O......#..#O#O...O#..........O.O#.
OO#...O..#..O.#..O........O.....O.....O........O..#....###..O..O..#...O...#.#.##.......O..#.O#O#O...
.......O..O...O...O.OO#.O.#..OO............O.O.O.#..OO#....##O.....O#.....O##O.....O.#...###O.....#.
.....O....OO..O..#O.....O#..O..O.#.#O#.#O#....O.#O..O#.O#.OO......#.O#..O.O...O...#...#..O....#...#O
#...#....#..O..#.O..#.##O..#O..#..OOO.#..##.....OO.O.....OOO.#..O.#.##O#..O.#..#.#.#....#O#...##....
.OO.O..###....#O....O..#.O.#.......#....O###..O.#......O#....#........O...#O..#..#.........O......##
O#......O#.....#.O....#..#.O.#O.......OO#...#..#.#O.#.O.##......#.#.#..O#O.#...O......O..OOO..OO.###
#O.#.O...O.....O....#OO##...OO..#...OO.#O.OOO.O..#O....O............#....#....#...........O.OOO#.OO.
.O.O.#OO#O.O.OO#..O..#....O.#..O...O....O.#OO#O...##.......O...O.....#.O...#...O.##O...OOO.O........
.O#..O.......#..##..OO#.....O.#.OO#...#O.O..O......O..OO.#...#.#............O#O.OO.......O#........#
.OO#..O##...........O.OO.....#.##....O#.#.#O..O...O..#..#..O.......#......#.#...........###O##......
...#...#.....O........OO..#.#...O.#.....#.OO.....#.O.#OO......#OO..#.....#..O##..O..#...O.O.......O#
O...O...#O..#O#O.......#..#O....#OO....#.#O.....O....O##.#...O.......#.....#........#O.O...O...O....
..#.....OO.#.OO....O#O....O#O...O.O...#O.##....##O...#..##O#.OO#O..OOO.#.#......O.OO.O..O.O.OO.#O.#.
.......O.....##...O#.#O.#....O.....O........#...#.OO...#...OO........#OO#...#..OO..O...#....O#.#OO.O
..#..O.OO....O...#.......OO.O.##...#...O...#.O#.OO#....O...O..O###.O.###O..#....##.O.##....O..##...#
.......OOO.....OO..#..O..#O.OOO#O.....O.O#.....O....###..O#..OO..O....###O.OO...#..#.O....#.O.O..#OO
.#...#.........#.###..#...#.#...#O.#..O#O#OOOO#.#.O.....O.#....#..O......#OO....O..O.O.#.O......O..O
O.O....O....O..O......O...#.....O##....OOO.....#.OO.#..O....O..O.O#.O#.OO......#....#......OO......O
..#O.#O.....O...#O...##.O#..O...#..##.O.O.O....#.....#.O#O.....#............O....#.#...#....###...#.
.O..O...OO.....#...........##...O#.O##O........OO..O.###O.#..#..#.O..O.......O....#O.#.........#...#
..##.O....O...O.O.O#O...#O..#..#.#.O...#O.O.O.....#....#...........O...O....O....OO#O.#.O....#..O..O
.#O#....O.....O#O..O##O#..O..O.O.....O........#.#......O#..O..OO.OO#.....OO........O......O#...O.#..
....#.....O..#......O...##.O.O.#..#.....O.O..#.O#.#O...##....O.......#O..O.....#..#.....#.....#.....
.OO....OOO.OO....O...#.......O......#O.#O..O.O#O..............O.#......#.#.......#....#.O....#.#O.#.
OO...OO.#.O#OO#........##.O.....#..#O.#....O..O..O.O#..........O..O.O....#..O.......O....O....O.OO.O
O.#O.......O#.#...#........O...O.O....O#OO....O.O.O...#.#.O#...#O#O.....##..O#O...#O....O....O#OOO.#
.OO#..#..O..........##.O.#OO#..O.O.#...#....O..O#.O..#.#.....O.O..#..OO..##O#...O.O.O#........O.....
.O............#.O...#...#....###.O#.#.#....O.....OO...#O.#.O..O##...#O.#O..OO#..#.O#.#.....O#.......
#..O#..#.OOO#O...O.O..O..#OOO....#...........O##...OOO##O....O.O...O..O#O.........O.OO.....O....O...
.#....#OO..##O..#O##O.#O....#......O.OO..........##.OO#O...#.....O.O.....O#..OO.#........O..O.O..O..
....O.##...#.O.........O.OO....#....##O..#O.......#..O.##O#..#O..#.....O.O#O....#...O..O.OO.........
..###.O.....##.....OO.O..OO.O..#...O.....#....O.O...............O.##.#.O#O#.#...O..#.OOOOO....#.#.#.
O...O....O.#.#.O#.O#..#......O#.....OO.##....OO.....O.......O....O.O.....#...O..OO#..OO..O....#OOO.#
#.#.#.O..#.O...#....#..O....O......O..#..O.OO#O......O.......O.O.........O...O.......OO.##.......#.#
O#..........O...#O.###....#O..O.#......O....#.O.O...O...O.....O#..#.O#..O.#.O..OOO#.O...#..O.#O.....
`
