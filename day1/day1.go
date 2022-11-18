package day1

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
	"strconv"
	"math"
)

type instruction struct{
  direction string
  distance int
}

func GetDistanceFromCenter(xcoord int, ycoord int) int {

	var sum = math.Abs(float64(xcoord)) + math.Abs(float64(ycoord))
	return int(sum) 
}

func part1(list []instruction) {
	var x = 0
	var y = 0
	var dir = "N";

	for _, step := range list {
		if (dir == "N") {
			if (step.direction == "L") {
				x = x - step.distance
				dir = "W"
			} else {
				x = x + step.distance
				dir = "E"
			}
		} else if (dir == "S") {
			if (step.direction == "R") {
				x = x - step.distance
				dir = "W"
			} else {
				x = x + step.distance
				dir = "E"
			}
		} else if (dir == "W") {
			if (step.direction == "L") {
				y = y - step.distance
				dir = "S"
			} else {
				y = y + step.distance
				dir = "N"
			}
		} else if (dir == "E") {
			if (step.direction == "R") {
				y = y - step.distance
				dir = "S"
			} else {
				y = y + step.distance
				dir = "N"
			}
		}
	}
	var result = GetDistanceFromCenter(x, y)
	fmt.Println("\tPart 1: ", result);
}

func Run() {
	fmt.Println("--- Day1 ---")
	absPath, _ := filepath.Abs("./day1/day1.txt")
	var input = utils.ReadFile(absPath)
	var instructionString = input[0];
	var instuctionList = getInstruction(instructionString)
	part1(instuctionList)
}

func getInstruction(line string) []instruction {
	var strList = strings.Split(line, ", ");
	var list []instruction;
	for _, element := range strList {
		var direction = string([]rune(element)[0]);
		distance, err := strconv.Atoi(string([]rune(element)[1:]))
    if err != nil {
        // ... handle error
        panic(err)
    }
    var newInstruction = instruction{direction, distance};
		list = append(list, newInstruction);
	}

	return list
}