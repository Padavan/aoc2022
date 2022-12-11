package day8

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	// "strings"
	"strconv"
	// "sort"
)

func part1(line []string) {
	var count = GetVisibleTreeCount(line)
	fmt.Println("\tPart 1: ", count);
} 
//  1809

func part2(line []string) {
	var score = GetHighestScenicScore(line)
	fmt.Println("\tPart 2: ", score);
}
// 3915660 too high


func Run() {
	fmt.Println("--- Day 8 ---")
	absPath, _ := filepath.Abs("./input/day8.txt")
	var multiLineInput = utils.ReadFile(absPath)
	part1(multiLineInput)
	part2(multiLineInput)
}

func Parse(multiline []string) ([][]int) {
	var matrix = [][]int{};

	for _, line := range multiline {
		var row = []int{};

		for i:=0; i< len(line); i++ {
			var substr = line[i:i+1];
			var height, err = strconv.Atoi(substr);
			if (err != nil) {
				panic(err);
			}
			row = append(row, height);
		}

		matrix = append(matrix, row);
	}

	return matrix;
}


func PrintMatrix(matrix [][]int) {
	for _, line := range matrix {
		fmt.Println(line);
	}
}

// PART1
func IsVisible(list []int, target_height int)(bool) {
	var result = true;
	for _, height := range list {
		if (height >= target_height) {
			result = false;
		}
	}
	return result;
}

func GetVisibilityMatrix(matrix [][]int)([][]int) {
	var size = len(matrix);

	var newMatrix = make([][]int, size)
	for i := range newMatrix {
	   newMatrix[i] = make([]int, size)
	}

	// var newMatrix = [size][size]uint8{}

	for rowIndex := 0; rowIndex < size; rowIndex++ {
		for columnIndex := 0; columnIndex < size; columnIndex++ {
			var target_tree_height = matrix[rowIndex][columnIndex];
			var left_trees = matrix[rowIndex][:columnIndex];
			var right_trees = matrix[rowIndex][columnIndex+1:];


			var top_trees = []int{}
			var bottom_trees = []int{};

			for dy := 0; dy < size; dy++ {
				if (dy < rowIndex) {
					top_trees = append(top_trees, matrix[dy][columnIndex]);
				} else if (dy > rowIndex) {
					bottom_trees = append(bottom_trees, matrix[dy][columnIndex]);
				}
			}

			var visibility = 0;
			if (len(left_trees) ==0 || len(right_trees) == 0 || len(top_trees) == 0 || len(bottom_trees) == 0) {
				visibility = 1;
			}

			if (IsVisible(left_trees, target_tree_height) || IsVisible(right_trees, target_tree_height) || IsVisible(top_trees, target_tree_height) || IsVisible(bottom_trees, target_tree_height)) {
				visibility = 1;
			}
			
			newMatrix[rowIndex][columnIndex] = visibility;

		}
	}

	return newMatrix
}

func CountTrees(matrix [][]int)(int) {
	var sum = 0;

	for _, row := range matrix {
		for _, element := range row {
			sum = sum + element;
		}
	}

	return sum;
}

func GetVisibleTreeCount(multiline []string)(int) {
	var matrix = Parse(multiline);

	var visibilityMatrix = GetVisibilityMatrix(matrix);

	var count = CountTrees(visibilityMatrix);

	// PrintMatrix(matrix)

	return count;
}



// PART2

//spizzheno
// func reverse(numbers []int) []int {
// 	for i := 0; i < len(numbers)/2; i++ {
// 		j := len(numbers) - i - 1
// 		numbers[i], numbers[j] = numbers[j], numbers[i]
// 	}
// 	return numbers
// }

func reverse(numbers []int) ([]int) {
	var length = len(numbers);
	var result = []int{};

	for i := length-1; i>=0; i-- {
		result = append(result, numbers[i]);
	}

	return result;
}

func CountVisibleTrees(tree_line []int, target_tree int)(int) {
	for idx, tree := range tree_line {
		if (tree >= target_tree) {
			return idx + 1;
		}
		// count = count + 1;
	}

	return len(tree_line);
} 

func GetScenicMatrix(matrix [][]int) ([][]int) {
	var size = len(matrix);

	var scenicMatrix = make([][]int, size)
	for i := range scenicMatrix {
	   scenicMatrix[i] = make([]int, size)
	}

	for rowIndex := 0; rowIndex < size; rowIndex++ {
		for columnIndex := 0; columnIndex < size; columnIndex++ {
			var target_tree_height = matrix[rowIndex][columnIndex];
			var left_trees = matrix[rowIndex][:columnIndex];
			var right_trees = matrix[rowIndex][columnIndex+1:];

			var top_trees = []int{}
			var bottom_trees = []int{};

			for dy := 0; dy < size; dy++ {
				if (dy < rowIndex) {
					top_trees = append(top_trees, matrix[dy][columnIndex]);
				} else if (dy > rowIndex) {
					bottom_trees = append(bottom_trees, matrix[dy][columnIndex]);
				}
			}

			left_trees = reverse(left_trees);
			top_trees = reverse(top_trees);

			var left_visibility = CountVisibleTrees(left_trees, target_tree_height);
			var right_visibility = CountVisibleTrees(right_trees, target_tree_height);
			var top_visibility = CountVisibleTrees(top_trees, target_tree_height);
			var bottom_visibility = CountVisibleTrees(bottom_trees, target_tree_height);

			var scenicScore = left_visibility * right_visibility * top_visibility * bottom_visibility;
			scenicMatrix[rowIndex][columnIndex] = scenicScore;
		}
	}

	return scenicMatrix
}

func GetHighestScore(matrix [][]int)(int) {
	var top = 0;

	for _, row := range matrix {
		for _, element := range row {
			if (element > top) {
				top = element
			}
		}
	}

	return top;
}

func GetHighestScenicScore(multiline []string)(int) {
	var matrix = Parse(multiline);

	var scenicMatrix = GetScenicMatrix(matrix);

	var top_score = GetHighestScore(scenicMatrix); 

	return top_score;
}