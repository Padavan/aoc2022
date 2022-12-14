package day12

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
	var count = GetPath(input)
	fmt.Println("\tPart 1: ", count);
} 
//  385 too low
// 771 too hight
// 391

func part2(input []string) {
	var count = GetPathForHiking(input)
	fmt.Println("\tPart 2: ", count);
}
// 386


func Run() {
	fmt.Println("--- Day 12 ---")
	// absPath, _ := filepath.Abs("./input/day12.example.txt")
	absPath, _ := filepath.Abs("./input/day12.txt")
	var input = utils.ReadFile(absPath)
	part1(input);
	part2(input);
}

func ParseInputToMatrix(input []string)([][]string) {

	var vsize = len(input);
	var matrix [][]string = make([][]string, vsize);

	for idx, line := range input {
		var hsize = len(line);

		var arr = make([]string, hsize);

		for i := 0; i < hsize; i++ {
			var letter = line[i:i+1];
			arr[i] = letter;
		}

		matrix[idx] = arr;
	}

	return matrix;
}

type pos struct {
	x int
	y int
}

var ALPHABET = []string {"S", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "E"}

func PrintMatrix(matrix [][]int) {
	for _, line := range matrix {
		// fmt.Println(line);
		for _, letter := range line {
			fmt.Printf("%4d", letter);
		}
		fmt.Printf("\n");
	}
}

func GetStartPosition(matrix [][]string)(pos) {
	var start pos;

	for i, line := range matrix {
		for j, letter := range line {
			if (letter == "S") {
				start = pos{j, i}
			}
		}
	}

	return start
}

func GetStartPositions(matrix [][]string)([]pos) {
	var start_arr []pos;

	for i, line := range matrix {
		for j, letter := range line {
			if (letter == "a") {
				start_arr = append(start_arr, pos{j, i});
			}
		}
	}

	return start_arr;
}

func IsReachable(letter string, target string)(bool) {
	var letter_index int;
	var target_index int;
	for idx, bookva := range ALPHABET {
		if (bookva == letter) {
			letter_index = idx;
		} 

		if (bookva == target) {
			target_index = idx;
		}
	}

	var diff = target_index - letter_index;
	if (diff <= 1) {
		return true;
	}
	return false;
}

func IsInsideMatrix(point pos, matrix [][]string)(bool) {
	var vsize = len(matrix);
	var hsize = len(matrix[0]);

	if (point.x >= 0 && point.x < hsize && point.y >= 0 && point.y < vsize) {
		return true;
	};

	return false;
}

func IsNotVisited(point pos, step int)(bool) {
	if (VisitedMatrix[point.y][point.x] == 0 || VisitedMatrix[point.y][point.x] > step) {
		return true;
	} 
	return false;
}

var VisitedMatrix [][]int;

func GetAvailablePoint(current pos, matrix [][]string, step int, minimal_step int)(int) {
	var up = pos{current.x, current.y+1};
	var down = pos{current.x, current.y-1};
	var left = pos{current.x-1, current.y};
	var right = pos{current.x+1, current.y};

	var candidates = []pos{up, down, left, right};


	var successfulSteps = []pos{};
	for _, candidate := range candidates {


		if (IsInsideMatrix(candidate, matrix) && IsReachable(matrix[current.y][current.x], matrix[candidate.y][candidate.x]) && IsNotVisited(candidate, step)) {

			if (matrix[candidate.y][candidate.x] == "E") {
				if (step < minimal_step) {
					minimal_step = step;
				}
			}

			VisitedMatrix[candidate.y][candidate.x] = step;
			successfulSteps = append(successfulSteps, candidate);
		}
	}

	for _, new_point := range successfulSteps {
		var new_end_step = GetAvailablePoint(new_point, matrix, step+1, minimal_step);
		if (minimal_step > step) {
			minimal_step = new_end_step
		}
	}

	return minimal_step;
}

func CreateVisitedMatrix(matrix [][]string) {

	for _, line := range matrix {
		var arr = make([]int, len(line));
		VisitedMatrix = append(VisitedMatrix, arr);
	}
}

func GetPathForHiking(input []string)(int) {
	var height_matrix = ParseInputToMatrix(input);
	CreateVisitedMatrix(height_matrix);

	var start_arr = GetStartPositions(height_matrix);

	var minimal_hiking_distance = 10000;
	for _, start := range start_arr {
		var length = GetAvailablePoint(start, height_matrix, 1, 10000);
		if (minimal_hiking_distance > length) {
			minimal_hiking_distance = length;
		}
	}

	// PrintMatrix(height_matrix);
	// PrintMatrix(VisitedMatrix);

	return minimal_hiking_distance;
}

func GetPath(input []string)(int) {
	var height_matrix = ParseInputToMatrix(input);
	CreateVisitedMatrix(height_matrix);

	var start = GetStartPosition(height_matrix);

	var answer = GetAvailablePoint(start, height_matrix, 1, 10000);
	// PrintMatrix(height_matrix);
	// PrintMatrix(VisitedMatrix);

	return answer;
}