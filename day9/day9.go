package day9

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
	"strconv"
)

type position struct {
	x int
	y int
}

func part1(input []string) {
	var count = GetVisitedPositions(input)
	fmt.Println("\tPart 1: ", count);
} 
//  5960

func part2(input []string) {
	var score = GetVisitedPositionsWithLongSnake(input)
	fmt.Println("\tPart 2: ", score);
}
// 2327

func Run() {
	fmt.Println("--- Day 9 ---")
	absPath, _ := filepath.Abs("./input/day9.txt")
	var multiLineInput = utils.ReadFile(absPath)
	part1(multiLineInput)
	part2(multiLineInput)
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

func GetNewPosition(head position, tail position)(position) {
	var diffx = utils.Abs(head.x - tail.x);
	var diffy = utils.Abs(head.y - tail.y);

	if (diffx <= 1 && diffy <= 1) {
		return tail;
	} else if (diffx != 0 && diffy == 0) {
		var diffstepx = (head.x - tail.x) / diffx;
		return position{tail.x + diffstepx,tail.y}
	} else if (diffx == 0 && diffy != 0) {
		var diffstepy = (head.y - tail.y) / diffy;
		return position{tail.x,tail.y + diffstepy};
	} else {
		var diffstepx = (head.x - tail.x) / diffx;
		var diffstepy = (head.y - tail.y) / diffy;

		return position{tail.x + diffstepx, tail.y + diffstepy};
	}
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

	var tail_position = position{0, 0};
	var tail_history = []position{tail_position};

	for _, head := range head_history {
		var last_tail_position =  tail_history[len(tail_history)-1];
		var tail_position = GetNewPosition(head, last_tail_position);
		tail_history = append(tail_history, tail_position);
	}

	var count = CountVisited(tail_history);

	return count;
}


// PART 2 
var SNAKE_LENGTH = 10;

func CreateSnake(start position) ([]position) {
	var snake = []position{}
	for i := 0; i< SNAKE_LENGTH; i++ {
		snake = append(snake, start);
	}

	return snake;
}

func GetVisitedPositionsWithLongSnake(input []string) (int) {
	var start_position = position{0, 0};
	var head_history = []position{start_position};

	for _,line := range input {
		var last_position = head_history[len(head_history)-1];
		var newpath = GetNewHeadPosition(line, last_position);
		head_history = append(head_history, newpath...);
	}

	var snake = CreateSnake(start_position);
	var tail_history = []position{start_position};

	for _, head_position := range head_history {
		snake[0] = head_position;
		for i := 1; i < SNAKE_LENGTH; i++ {
			snake[i] = GetNewPosition(snake[i-1], snake[i]);
		}
		tail_history = append(tail_history, snake[SNAKE_LENGTH - 1]);
	}

	var count = CountVisited(tail_history);

	return count;
}