package day10

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
	"strconv"
	// "sort"
)


func part1(input []string) {
	var count = GetSignalStrenghtSum(input)
	fmt.Println("\tPart 1: ", count);
} 
//  13520

func part2(input []string) {
	fmt.Println("\tPart 2: ");
	GetCRTImage(input)
}
// PGPHBEAB


func Run() {
	absPath, _ := filepath.Abs("./day10/day10.txt")
	var input = utils.ReadFile(absPath)
	part1(input)
	part2(input)
}

func CatchDuring(pikes []int, cycle int, x int) ([]int) {
	var new_pikes = pikes;
	if (cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220) {
		new_pikes = append(new_pikes, x*cycle);
	}

	return new_pikes;
}

func CollectPikes(pikes []int) (int) {
	var total int;

	for _, strength := range pikes {
		total = total + strength;
	}

	return total;
}

func GetSignalStrenghtSum(multiline []string)(int) {
	var pikes = []int{};

	var the_x = 1;
	var cycle = 1;

	for _, line := range multiline {
		if (line == "noop") {
			pikes = CatchDuring(pikes, cycle, the_x);
			cycle = cycle + 1;
		} else {
			pikes = CatchDuring(pikes, cycle, the_x);
			var splitted_line = strings.Split(line, " ");
			var diff, err = strconv.Atoi(splitted_line[1]);
			if (err != nil) {
				panic(err);
			}

			cycle = cycle + 1;

			pikes = CatchDuring(pikes, cycle, the_x);
			the_x = the_x + diff;
			cycle = cycle + 1;
		};
	};

	var total = CollectPikes(pikes);

	return total;
};

// part 2 

func Display(pixels [240]bool) {
	for idx, pixel := range pixels {
		var symbol string;
		if (pixel) {
			symbol = "#"
		} else {
			symbol = "."
		}
		fmt.Printf("%s", symbol);

		if ((idx+1) % 40 == 0) {
			fmt.Printf("\n");
		}
	}
}

func GetPixel(cycle int, position [3]int)(bool) {
	var pixel = false;
	for _, xcoord := range position {
		if (xcoord+1 == (cycle % 40)) {
			pixel = true;
		}
	} 

	return pixel;
}

func MoveSprite(current [3]int, xdiff int)([3]int) {
	var new_sprite_position = [3]int{}
	for idx, cur := range current {
		new_sprite_position[idx] = cur + xdiff;
	}
	return new_sprite_position;
}

func GetCRTImage(input []string) {

	var CRT = [240]bool{};

	var the_x = 1;
	var cycle = 1;
	var sprite_position = [3]int{0, 1, 2};
	var pixel bool;

	for _, line := range input {
		if (line == "noop") {
			// fmt.Printf("Cycle before %d, value %d \n", cycle, the_x);
			pixel = GetPixel(cycle, sprite_position);
			CRT[cycle-1] = pixel;

			cycle = cycle + 1;
		} else {
			// fmt.Printf("Cycle before %d, value %d \n", cycle, the_x);
			pixel = GetPixel(cycle, sprite_position);
			CRT[cycle-1] = pixel;

			var splitted_line = strings.Split(line, " ");
			var diff, err = strconv.Atoi(splitted_line[1]);
			if (err != nil) {
				panic(err);
			}

			cycle = cycle + 1;

			// fmt.Printf("Cycle before %d, value %d \n", cycle, the_x);
			pixel = GetPixel(cycle, sprite_position);
			CRT[cycle-1] = pixel;
			the_x = the_x + diff;
			sprite_position = MoveSprite(sprite_position, diff);
			cycle = cycle + 1;
		};
	};

	Display(CRT);

}