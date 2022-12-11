package day3

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
)

func part1(multiLineInput []string) {
	var score = GetTotalItemPriority(multiLineInput)
	fmt.Println("\tPart 1: ", score);
}
// 7581

func part2(multiLineInput []string) {
	var score = GetTotalGroupPriority(multiLineInput);
	fmt.Println("\tPart 2: ", score)
}
// 2525

func Run() {
	fmt.Println("--- Day2 ---")
	absPath, _ := filepath.Abs("./input/day3.txt")
	var multiLineInput = utils.ReadFile(absPath)
	part1(multiLineInput)
	part2(multiLineInput)
}

var PRIORITY = []string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
func IndexOf(element string, data []string) (int) {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1
}

func GetScore(substr string) (int) {
	return IndexOf(substr, PRIORITY) + 1;
}

func SplitInHalf(line string) (string, string) {
	length := len(line)
	splitIndex := length / 2;
	first := line[0:splitIndex]
	second := line[splitIndex:length]

	return first, second
} 

func FindCommonItem(line1 string, line2 string) (string) {
	var result string;
	for idx, _ := range line1 {
		substr := line1[idx:idx+1];
		if (strings.Contains(line2, substr)) {
			result = substr;
		}
	}
	return result
}

func FindCommonItemInThree(line1 string, line2 string, line3 string) (string) {
	var result string;
	for idx, _ := range line1 {
		substr := line1[idx:idx+1];
		if (strings.Contains(line2, substr) && strings.Contains(line3, substr)) {
			result = substr;
		}
	}

	return result
}

// PART 1
func GetTotalItemPriority(multilineInput []string) (int) {
	var total = 0;
	for _, line := range multilineInput {

		firstHalf, secondHalf := SplitInHalf(line)
		item := FindCommonItem(firstHalf, secondHalf)
		itemScore := GetScore(item);
		total = total + itemScore
	}

	return total;
}

// PART 2
func GetTotalGroupPriority(multilineInput []string) (int) {
	var total = 0;

	elfCount := len(multilineInput);
	elfGroupCount := elfCount /3

	for i := 0; i < elfGroupCount ; i++ {
		first := multilineInput[i*3]
		second := multilineInput[i*3 + 1]
		third := multilineInput[i*3 + 2]
		badge := FindCommonItemInThree(first, second, third);
		total = total + GetScore(badge);
	}

	return total;
}