package main

import (
	"fmt"
	"math"
	_ "math"
	"os"
	"slices"
	"strings"
)

func main() {
	//P1()
	P2()
}

func P1() {
	data, _ := os.ReadFile("test.txt")
	lines := strings.Split(string(data), "\n")
	lines = lines[0 : len(lines)-1]
	PrintMap(lines)
	fmt.Println()
	fmt.Println(getTrailScores(lines))
}

func P2() {
	data, _ := os.ReadFile("test.txt")
	lines := strings.Split(string(data), "\n")
	lines = lines[0 : len(lines)-1]
	PrintMap(lines)
	fmt.Println()
	fmt.Println(getTrailScores2(lines))
}

func getTrailScores(_map []string) int {
	scores := 0
	for y := range _map {
		for x := range _map[y] {
			if _map[y][x] != '0' {
				continue
			}
			scores += calculateTrailScore([2]int{y, x}, _map)
		}
	}
	return scores
}

func calculateTrailScore(head [2]int, _map []string) int {
	paths := getPaths(head, _map)
	nines := make([][2]int, 0)
	for {
		if len(paths) == 0 {
			break
		}
		current := paths[len(paths)-1]
		paths = paths[0 : len(paths)-1]
		if _map[current[0]][current[1]] != '9' {
			paths = append(paths, getPaths(current, _map)...)
		} else {
			if slices.Contains(nines, current) == false {
				nines = append(nines, current)
			}
		}
	}
	return len(nines)
}

func getTrailScores2(_map []string) int {
	scores := 0
	for y := range _map {
		for x := range _map[y] {
			if _map[y][x] != '0' {
				continue
			}
			score := calculateTrailScore2([2]int{y, x}, _map)
			fmt.Printf("head:%v,score:%d\n", [2]int{y, x}, score)
			scores += score
		}
	}
	return scores
}

func calculateTrailScore2(head [2]int, _map []string) int {
	paths := getPaths(head, _map)
	score := len(paths)
	prev := score
	for {
		if len(paths) == 0 {
			break
		}
		current := paths[len(paths)-1]
		paths = paths[0 : len(paths)-1]
		paths = append(paths, getPaths(current, _map)...)
		if prev != len(paths) {
			if len(paths) > prev {
				diff := int(math.Abs(float64(len(paths) - prev)))
				score += diff
			}
			prev = len(paths)
		}
	}
	return score
}

func getPaths(head [2]int, _map []string) [][2]int {
	paths := make([][2]int, 0)
	y := head[0]
	x := head[1]
	c := _map[head[0]][head[1]]
	if c == '.' {
		return [][2]int{}
	}

	if y > 0 {
		next := _map[y-1][x]
		if next-1 == c {
			paths = append(paths, [2]int{y - 1, x})
		}
	}

	if y < len(_map)-1 {
		next := _map[y+1][x]
		if next-1 == c {
			paths = append(paths, [2]int{y + 1, x})
		}
	}

	if x > 0 {
		next := _map[y][x-1]
		if next-1 == c {
			paths = append(paths, [2]int{y, x - 1})
		}
	}

	if x < len(_map[y])-1 {
		next := _map[y][x+1]
		if next-1 == c {
			paths = append(paths, [2]int{y, x + 1})
		}
	}
	return paths
}

func validPathsCount(paths [][2]int, _map []string) int {
	count := 0
	for i := range paths {
		if isPathValid(paths[i], _map) {
			count++
		}
	}
	return count
}

func isPathValid(head [2]int, _map []string) bool {
	y := head[0]
	x := head[1]
	c := _map[head[0]][head[1]]
	if c == '.' {
		return false
	}
	if c == '9' {
		return true
	}

	if y > 0 {
		next := _map[y-1][x]

		if next-1 == c {
			return isPathValid([2]int{y - 1, x}, _map)
		}
	}

	if y < len(_map)-1 {
		next := _map[y+1][x]
		if next-1 == c {
			return isPathValid([2]int{y + 1, x}, _map)
		}
	}

	if x > 0 {
		next := _map[y][x-1]
		if next-1 == c {
			return isPathValid([2]int{y, x - 1}, _map)
		}
	}

	if x < len(_map[y])-1 {
		next := _map[y][x+1]
		if next-1 == c {
			return isPathValid([2]int{y, x + 1}, _map)
		}
	}
	return false
}

func PrintMap(_map []string) {
	for _, line := range _map {
		fmt.Println(line)
	}
}
