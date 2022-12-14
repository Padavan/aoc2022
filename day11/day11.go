package day11

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
	"strconv"
	"sort"
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
	fmt.Println("--- Day 11 ---")
	absPath, _ := filepath.Abs("./input/day11.txt")
	var input = utils.ReadFile(absPath)
	part1(input);
	part2(input);
}


type monkey struct {
	name int
	items []int64
	operation string
	test int64
	outcomeTrue int
	outcomeFalse int

} 

func getItems(line string) ([]int64) {
	var items = []int64{};
	var index = strings.Index(line, ":");
	var commaSeparated = line[index+1:];

	var spllittedNumbers = strings.Split(strings.Trim(commaSeparated, " "), ", ");	
	for _, string_number := range spllittedNumbers {
		var item, err = strconv.ParseInt(string_number, 10, 64);
		if (err != nil) {
			panic(err)
		};

		items = append(items, item);
	}

	return items;
}

func getName(line string)(int) {
	var nameString = strings.Fields(line);
	var converted, err = strconv.Atoi(nameString[1][0:1]);
	if (err != nil){
		panic(err);
	}

	return converted;
}

func getOperation(line string)(string) {
	var index = strings.Index(line, ":");
	var trimmed = strings.Trim(line[index+1:], " ");
	return trimmed;
}

func getTestValue(line string)(int64){
	var newline = strings.Replace(line, "Test: divisible by", "", 1);
	newline = strings.Trim(newline, " ");
	var value, err = strconv.ParseInt(newline, 10, 64);
	if (err != nil) {
		panic(err)
	}
	return value;
}

func getOutcome(line string)(int){
	var index = strings.Index(line, ":");
	var stripped = strings.Replace(line[index+1:], "throw to monkey", "", 1);
	var trimmed = strings.Trim(stripped, " ");

	var value, err = strconv.Atoi(trimmed);
	if (err != nil) {
		panic(err)
	}
	return value;
}

func DoOperation(operation string, value int64)(int64) {
	var stripped = strings.Replace(operation, "new = old ", "", 1);
	var separated = strings.Split(stripped, " ");

	var right int64;
	if (separated[1] == "old") {
		right = value;
	} else if (separated[1] == "old" && separated[0] == "*") {
		right = 1;
	} else {
		var parsedValue, err = strconv.ParseInt(separated[1], 10, 64);
		if (err != nil) {
			panic(err);
		}
		right = parsedValue;
	}

	if (separated[0] == "*") {
		return value * right;
	} else if (separated[0] == "+") {
		return value + right;
	} else {
		fmt.Printf("error occured while parsing operation. operaiation: %s, value: %d, right: %d \n", operation, value, right);
		return 0;
	}
}

func TestWorryLevel(dividers int64, level int64) (bool) {
	if (level % dividers == 0) {
		return true;
	} else {
		return false;
	}
}

func remove(arr []int, target int) ([]int) {
	var new_arr = []int{};
	for _, element := range arr {
		if (element != target) {
			new_arr = append(new_arr, element);
		}
	}

	return new_arr;
}

func PlayRound(state []monkey, throw_count []int)([]monkey, []int) {
	var current_state = state;
	var current_throw_state = throw_count;
	for idx, current_monkey := range current_state {
		for _, item := range current_monkey.items {
			var worry_level = DoOperation(current_monkey.operation, item);
			worry_level = worry_level / 3;

			var testing_result = TestWorryLevel(current_monkey.test, worry_level);

			if (testing_result) {
				current_state[current_monkey.outcomeTrue].items = append(current_state[current_monkey.outcomeTrue].items, worry_level); 
			} else {
				current_state[current_monkey.outcomeFalse].items = append(current_state[current_monkey.outcomeFalse].items, worry_level);
			}
			current_throw_state[idx] =  current_throw_state[idx] + 1;
		}		
		current_state[idx].items = []int64{};
	};

	return state, throw_count;
}

func PrintMonkeyItemState (state []monkey) {
	for _, mon := range state {
		fmt.Println("Monkey", mon.name, ": ", mon.items);
	}

	fmt.Printf("\n");
}

func PrintThrowCount(list []int) {
	for idx, element := range list {
		fmt.Println("Monkey", idx, ": ", element);
	}
}

func GetMonkeyBusiness(input []string)(int) {
	var ROUND_COUNT = 20;
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
	};

	var throw_count = make([]int, monkey_count);
	for i:= 0; i < ROUND_COUNT; i++ {
		state, throw_count = PlayRound(state, throw_count);
	}

	sort.Ints(throw_count)
	var monkey_business = throw_count[len(throw_count) - 1] * throw_count[len(throw_count) - 2];

	return monkey_business;
}



// PART 2
func PlayRound10KEdition(state []monkey, throw_count []int)([]monkey, []int) {
	var chinese_number int64 = 1;

	for _, mon := range state {
		chinese_number = chinese_number * mon.test;
	}
	// fmt.Println("chinese stuff", chinese_number);

	var current_state = state;
	var current_throw_state = throw_count;
	for idx, current_monkey := range current_state {
		for _, item := range current_monkey.items {
			var worry_level = DoOperation(current_monkey.operation, item);
			var divider = (chinese_number);

			var testing_result = TestWorryLevel(current_monkey.test, worry_level);

			if (testing_result) {
				current_state[current_monkey.outcomeTrue].items = append(current_state[current_monkey.outcomeTrue].items, worry_level % divider); 
			} else {
				current_state[current_monkey.outcomeFalse].items = append(current_state[current_monkey.outcomeFalse].items, worry_level % divider);
			}
			current_throw_state[idx] =  current_throw_state[idx] + 1;
		}		
		current_state[idx].items = []int64{};
	};

	return state, throw_count;
}

func GetMonkeyBusiness10K(input []string)(int) {
	var ROUND_COUNT = 10000;
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
	};

	var throw_count = make([]int, monkey_count);
	for i:= 0; i < ROUND_COUNT; i++ {
		state, throw_count = PlayRound10KEdition(state, throw_count);
		if (i + 1 == 1 || i + 1 == 20 || i + 1 == 1000 || i + 1 == 2000 || i + 1 == 10000) {
			// fmt.Println("--- Round ", i+1);
			// PrintThrowCount(throw_count);
			// PrintMonkeyItemState(state);
		}
	}

	sort.Ints(throw_count)
	var monkey_business = throw_count[len(throw_count) - 1] * throw_count[len(throw_count) - 2];

	return monkey_business;
}