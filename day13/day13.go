package day13

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
	var count = GetSumOfIncicesInRightOrder(input)
	fmt.Println("\tPart 1: ", count);
} 
//  


// func part2(input []string) {
// 	var count = GetPathForHiking(input)
// 	fmt.Println("\tPart 2: ", count);
// }
// 


func Run() {
	fmt.Println("--- Day 12 ---")
	// absPath, _ := filepath.Abs("./input/day12.example.txt")
	absPath, _ := filepath.Abs("./input/day13.example.txt")
	var input = utils.ReadFile(absPath)
	part1(input);
	// part2(input);
}

func getNextToken(packet string, index int)(string) {
	var nextChar = packet:index
}

func IsRightOrder(left string, right string)(bool) {
	var left_length = len(left);
	var stack int;

	var leftIndex = 0;
	var rightIndex = 0;

	for leftIndex != -1 {
		var leftToken = getNextToken(left, leftIndex);
	} 


	return true;
}

func ParseArray(line string)()

// PART 1;

func GetSumOfIncicesInRightOrder(input []string)(int) {
	var pairs_count = (len(input) + 1) / 3;

	for i := 0; i < pairs_count; i++ {
		var left = input[i*3];
		var right = input[i*3 + 1];
		if (IsRightOrder(left, right)) {
			fmt.Println(i + 1);
		};
		fmt.Println(left, right);
	};


	return 0;
}