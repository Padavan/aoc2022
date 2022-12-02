package day2

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	 "math"
)

func part1(multiLineInput []string) {
	var score = GetTotalScore(multiLineInput)
	fmt.Println("\tPart 1: ", score);
}
//15632

func part2(multiLineInput []string) {
	var sum = GetTotalScore2(multiLineInput);
	fmt.Println("\tPart 2: ", sum)
}
// 14416

func Run() {
	fmt.Println("--- Day2 ---")
	absPath, _ := filepath.Abs("./day2/day2.txt")
	var multiLineInput = utils.ReadFile(absPath)
	part1(multiLineInput)
	part2(multiLineInput)
}

var SHAPE_MAPPING = []string{"A", "B", "C"};
var SHAPE_MAPPING_2 = []string{"X", "Y", "Z"};
var OUTCOME_MAPPING = []string{"X", "Y", "Z" }

func GetShapeScore(player int) (int) {
	return player + 1;
}

func GetOutcomeScore(index int) (int) {
	return index*3;
}

func IndexOf(element string, data []string) (int) {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1
}

// PART1
func GetTotalScore(listInput []string) (int) {
	var totalScore = 0;

	for _, line := range listInput {
		var opponentHand = line[0:1];
		var myHand = line[2:3]; 

		var opponentIndex = IndexOf(opponentHand, SHAPE_MAPPING);
		var playerIndex = IndexOf(myHand, SHAPE_MAPPING_2);
		var outcomeIndex = GetOutcomeIndex(opponentIndex, playerIndex);

		totalScore = totalScore + GetOutcomeScore(outcomeIndex) + GetShapeScore(playerIndex)
	}

	return totalScore;
}

func GetOutcomeIndex(op int, player int) (int) {
	var diff = player - op;
	if (diff == 0) {
		return 1;
	}

	if ((diff > 0 && math.Abs(float64(diff)) == 1) || (diff < 0 && math.Abs(float64(diff)) == 2)) {
		return 2
	}

	return 0;
}

// PART 2
func GetTotalScore2(listInput []string)(int) {
	var totalScore = 0;

	for _, line := range listInput {
		var opponentHand = line[0:1];
		var outcome = line[2:3]; 

		var opponentHandIndex = IndexOf(opponentHand, SHAPE_MAPPING);
		var outcomeIndex = IndexOf(outcome, OUTCOME_MAPPING);
		var playerHandIndex = GetPlayerHandIndex(opponentHandIndex, outcomeIndex);

		totalScore = totalScore + GetOutcomeScore(outcomeIndex) + GetShapeScore(playerHandIndex);
	}

	return totalScore;
}

func GetPlayerHandIndex(opponentHandIndex int, outcomeIndex int) (int) {
	var diff int;

	// draw
	if (outcomeIndex == 1) {
		diff = 0
	}

	// win
	if (outcomeIndex == 2) {
		if (opponentHandIndex + 1 > 2) {
			diff = -2
		} else {
			diff = 1
		}
	}

	// lose
	if (outcomeIndex == 0) {
		if (opponentHandIndex - 1 < 0) {
			diff = 2
		} else {
			diff = -1
		}
	}

	return opponentHandIndex + diff;
}