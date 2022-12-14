package day12

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	// "strings"
	// "strconv"
	// "sort"
	// "regexp"
)


func part1(input []string) {
	var count = GetMonkeyBusiness(input)
	fmt.Println("\tPart 1: ", count);
} 
//  117624

func part2(input []string) {
	var count = GetMonkeyBusiness10K(input)
	fmt.Println("\tPart 2: ", count);
}
// 


func Run() {
	fmt.Println("--- Day 12 ---")
	absPath, _ := filepath.Abs("./input/day11.txt")
	var input = utils.ReadFile(absPath)
	part1(input);
	part2(input);
}
