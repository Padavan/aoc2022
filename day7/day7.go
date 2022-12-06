package day7

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	// "strings"
)

func part1(line []string) {
	var index = Part1Function(line)
	fmt.Println("\tPart 1: ", index);
}
// 

func part2(line []string) {
	var index = Part2Function(line);
	fmt.Println("\tPart 2: ", index)
}
// 

func Run() {
	fmt.Println("--- Day 7 ---")
	absPath, _ := filepath.Abs("./day7/day7.txt")
	var multiLineInput = utils.ReadFile(absPath)
	part1(multiLineInput)
	part2(multiLineInput)
}





// PART 1
func Part1Function(input []string) (int) {
	return len(input);
}

// PART 2
func Part2Function(input []string) (int) {
	return len(input);
}
