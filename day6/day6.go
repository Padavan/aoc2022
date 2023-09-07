package day6

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
)

type instruction struct {
	move int
	from int
	to int
}

func part1(line string) {
	var index = GetMarkerIndex(line, 4)
	fmt.Println("\tPart 1: ", index);
}
// 1034

func part2(line string) {
	var index = GetMarkerIndex(line, 14);
	fmt.Println("\tPart 2: ", index)
}
// 2472

func Run() {
	fmt.Println("--- Day 6 ---")
	absPath, _ := filepath.Abs("./input/day6.txt")
	var multiLineInput = utils.ReadFile(absPath)
	part1(multiLineInput[0])
	part2(multiLineInput[0])
}

func AreCharDifferent(testString string, window int)(bool) {
	var result = true;
	for i:=0; i < window-1; i++ {
		var pickedLetter = testString[i:i+1];
		var newString = testString[i+1:];

		if (strings.Contains(newString, pickedLetter)) {
			result = false;
		}
	}

	return result;
}

func GetMarkerIndex(line string, window int)(int) {
	var inputLength = len(line);
	var uniqueIndex = []int{};

	for i:=0; i <= inputLength-window; i++ {
		var substring = line[i:i+window];
		if (AreCharDifferent(substring, window)) {
			uniqueIndex = append(uniqueIndex, i);
		};
	}

	return uniqueIndex[0] + window;
}