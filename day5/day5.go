package day5

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
	"strconv"
)

type instruction struct {
	move int
	from int
	to int
}

func part1(multiLineInput []string) {
	var word = GetTopRearrangment(multiLineInput)
	fmt.Println("\tPart 1: ", word);
}
// WHTLRMZRC

func part2(multiLineInput []string) {
	var word = GetTopRearrangment2(multiLineInput);
	fmt.Println("\tPart 2: ", word)
}
// GMPMLWNMG

func Run() {
	fmt.Println("--- Day5 ---")
	absPath, _ := filepath.Abs("./input/day5.txt")
	var multiLineInput = utils.ReadFile(absPath)
	part1(multiLineInput)
	part2(multiLineInput)
}


func ParseInput(multiLineInput []string)([]string, []string) {
	var emptyLineIndex = utils.IndexOf("", multiLineInput)
	var inputLength = len(multiLineInput)

	createState := multiLineInput[0:emptyLineIndex];
	instructions := multiLineInput[emptyLineIndex+1:inputLength];
	

	return createState, instructions
}


var CRATE_WIDTH = 4;

func GetCratesStateFromInput(input []string)([][]string) {
	var lastLine = input[len(input) - 1]
	var COLUMNS = (len(lastLine) + 1) / 4;
	var height = len(input);

	var state = [][]string{};

	// fill the state with empty stacks
	var emptyStack = []string{};
	for i:=0; i < COLUMNS; i++ {
		state = append(state, emptyStack)
	}

	// read input stacks in reversed order. from bottom to top.
	for k := (height-2); k >= 0; k-- {
		currentLine := input[k];
		// get crates
		for currentStackIndex:=0; currentStackIndex < COLUMNS; currentStackIndex++ {
			var crateNameIndex = currentStackIndex * CRATE_WIDTH + 1;
			var crateName = currentLine[crateNameIndex:crateNameIndex+1]
			var currentStack = state[currentStackIndex];

			if (crateName != " ") {
				currentStack = append(currentStack, crateName);
				state[currentStackIndex] = currentStack
			}
		}
	}

	return state;
}

func GetParsedInstruction(input []string)([]instruction) {
	var result = []instruction{};

	for _, line := range input {
		var stringArray = strings.Split(line, " ");
		var move, err1 = strconv.Atoi(stringArray[1]);
		var from, err2 = strconv.Atoi(stringArray[3]);
		var to, err3 = strconv.Atoi(stringArray[5]);
		if err1 != nil || err2 != nil || err3 != nil {
		    // ... handle error
		    panic(err1)
		};

		var parsedInstruction = instruction{move, from, to};
		result = append(result, parsedInstruction);
	}

	return result
}

func GetTopCrateNames(state [][]string)(string) {
	var joinedString = "";
	for _, stack := range state {
		var stackLength = len(stack);
		var topCrate = stack[stackLength-1] 
		joinedString = joinedString + topCrate
	}

	return joinedString;
}

// PART1 
func ApplyInstruction(state [][]string, inst instruction)([][]string) {
	var amountOfCrates = inst.move;
	var currentState = state;
	var fromStack = currentState[inst.from - 1];
	var toStack = currentState[inst.to - 1];

	for i:=0; i< amountOfCrates; i++ {
		// pop stack
		var fromStackLength = len(fromStack);
		var transfferedCrate = fromStack[fromStackLength-1];
		fromStack = fromStack[:fromStackLength-1]

		// push stack
		toStack = append(toStack, transfferedCrate);

		// update state(stack list)
		currentState[inst.from - 1] = fromStack;
		currentState[inst.to - 1] = toStack;
	}

	return currentState;
}

func ApplyInstructionSet(state [][]string, instructionList []instruction)([][]string) {
	var lastState = state;
	for _, oneInstruction := range instructionList {
		lastState = ApplyInstruction(lastState, oneInstruction);
	}
	return lastState;
}

func GetTopRearrangment(multiLineInput []string)(string) {
	var cratesStateInput, instruction = ParseInput(multiLineInput)

	var cratesState = GetCratesStateFromInput(cratesStateInput);
	var parsedInstruction = GetParsedInstruction(instruction); 

	var lastState = ApplyInstructionSet(cratesState, parsedInstruction)
	var topCrateString = GetTopCrateNames(lastState);

	return topCrateString;
}

// PART 2
func ApplyInstruction9001(state [][]string, inst instruction)([][]string) {
	var amount = inst.move;
	var currentState = state;
	var fromStack = currentState[inst.from - 1];
	var toStack = currentState[inst.to - 1];

	// pop stack
	var fromStackLength = len(fromStack);
	var transfferedCrateStack = fromStack[fromStackLength-amount : fromStackLength];
	fromStack = fromStack[:fromStackLength-amount]

	// push stack
	toStack = append(toStack, transfferedCrateStack...);

	// update state
	currentState[inst.from - 1] = fromStack;
	currentState[inst.to - 1] = toStack;

	return currentState;
}

func ApplyInstructionSet2(state [][]string, instructionList []instruction)([][]string) {
	var lastState = state;
	for _, oneInstruction := range instructionList {
		lastState = ApplyInstruction9001(lastState, oneInstruction);
	}
	return lastState;
}

func GetTopRearrangment2(multiLineInput []string)(string) {
	var cratesStateInput, instruction = ParseInput(multiLineInput)

	var cratesState = GetCratesStateFromInput(cratesStateInput);
	var parsedInstruction = GetParsedInstruction(instruction); 

	var lastState = ApplyInstructionSet2(cratesState, parsedInstruction)
	var topCrateString = GetTopCrateNames(lastState);

	return topCrateString;
}