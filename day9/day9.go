package day9

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
	"strconv"
	// "sort"
)

type position struct {
	x int
	y int
}

func part1(line []string) {
	var count = GetVisitedPositions(line)
	fmt.Println("\tPart 1: ", count);
} 
//  

// func part2(line []string) {
// 	var score = GetHighestScenicScore(line)
// 	fmt.Println("\tPart 2: ", score);
// }
// 


func Run() {
	fmt.Println("--- Day 9 ---")
	absPath, _ := filepath.Abs("./input/day9.txt")
	var multiLineInput = utils.ReadFile(absPath)
	part1(multiLineInput)
	// part2(multiLineInput)
}


func GetNewHeadPosition(line string, pos position) ([]position) {
	var splitted = strings.Split(line, " ");
	var direction = splitted[0];
	var diff, err = strconv.Atoi(splitted[1]);
	if (err != nil) {
		panic(err);
	}

	var xdiff = 0;
	var ydiff = 0;

	switch direction {
		case "U":
			ydiff = 1;
		case "D":
			ydiff = -1;
		case "R":
			xdiff = 1;		
		case "L":
			xdiff = -1;
	}

	var path = []position{};
	var current_position = pos;

	for i:=0; i < diff; i++ {
		var new_x = current_position.x + xdiff;
		var new_y = current_position.y + ydiff; 
		current_position = position{ new_x, new_y };
		path = append(path, current_position);
	}

	return path;
}

func IsPositionEqual(a position, b position)(bool) {
	if (a.x == b.x && a.y == b.y) {
		return true;
	}
	return false;
}

func CountVisited(list []position)(int) {
	var unique_pos_list = []position{};

	for i := 0; i < len(list); i++ {
		var checked_pos = list[i];

		var isUnique = true;
		for _, unique_pos := range unique_pos_list {
			if (IsPositionEqual(unique_pos, checked_pos)) {
				isUnique = false;
			}
		};

		if (isUnique) {
			unique_pos_list = append(unique_pos_list, checked_pos);
		};

	}

	return len(unique_pos_list);
}

// PART 1
func GetVisitedPositions(multiline []string) (int) {

	var head_position = position{0, 0};
	var head_history = []position{head_position};


	for _,line := range multiline {
		var last_position = head_history[len(head_history)-1];
		var path = GetNewHeadPosition(line, last_position);
		head_history = append(head_history, path...);
	}

	// var tail_position = position{0, 0};
	// var tail_history = []position{tail_position};

	// for _, tpos := range head_history {
		// fmt.Println(tpos);
	// }
	var count = CountVisited(head_history);

	return count;
}