package day11

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
	"strconv"
	// "sort"
	// "regexp"
)


func part1(input []string) {
	var count = GetMonkeyBusiness(input)
	fmt.Println("\tPart 1: ", count);
} 
// 

// func part2(input []string) {
// 	fmt.Println("\tPart 2: ");
// 	GetCRTImage(input)
// }
// 


func Run() {
	fmt.Println("--- Day 10 ---")
	absPath, _ := filepath.Abs("./input/day11.txt")
	var input = utils.ReadFile(absPath)
	part1(input)
	// part2(input)
}


type monkey struct {
	name string
	items []int
	operation string
	test int
	outcomeTrue string
	outcomeFalse string

} 

func getItems(line string) ([]int) {
	var items = []int{};
	var index = strings.Index(line, ":");
	var commaSeparated = line[index+1:];

	var spllittedNumbers = strings.Split(strings.Trim(commaSeparated, " "), ", ");	
	for _, string_number := range spllittedNumbers {
		var item, err = strconv.Atoi(string_number);
		if (err != nil) {
			panic(err)
		};

		items = append(items, item);
	}

	return items;
}

func getName(line string)(string) {
	var nameString = strings.Fields(line);
	return nameString[1];
}

func getOperation(line string)(string) {
	var index = strings.Index(line, ":");
	var trimmed = strings.Trim(line[index+1:], " ");
	return trimmed;
}

func getTestValue(line string)(int){
	var newline = strings.Replace(line, "Test: divisible by", "", 1);
	newline = strings.Trim(newline, " ");
	var value, err = strconv.Atoi(newline);
	if (err != nil) {
		panic(err)
	}
	return value;
}

func getOutcome(line string)(string){
	var index = strings.Index(line, ":");
	var stripped = strings.Replace(line[index+1:], "throw to monkey", "", 1);
	var trimmed = strings.Trim(stripped, " ");

	// var value, err = strconv.Atoi(trimmed);
	// if (err != nil) {
	// 	panic(err)
	// }
	return trimmed;
}

func GetMonkeyBusiness(input []string)(int) {


	var monkey_count = (len(input) + 1) / 7;
	var state = make([]monkey, monkey_count);

	for i:= 0; i < monkey_count; i++ {
		var monkeyName = getName(input[i*7]);
		var items = getItems(input[i*7+1]);
		var operation = getOperation(input[i*7+2]);
		var test = getTestValue(input[i*7+3]);
		var outcomeTrue = getOutcome(input[i*7+4]);
		var outcomeFalse = getOutcome(input[i*7+5]);

		var new_monkey = monkey{monkeyName, items, operation, test, outcomeTrue, outcomeFalse};
		state[i] = new_monkey;
	}
	
	for _, mon := range state {
		fmt.Println(mon);
	}

	return 0;
}


