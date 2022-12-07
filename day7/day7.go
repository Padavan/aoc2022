package day7

import (
	"fmt"
	"path/filepath"
	"aoc2022/utils"
	"strings"
	"strconv"
	"sort"
)

type file struct {
	name string
	size int
}

type dir struct {
	name string
	parent *dir
	folders []*dir
	files []file
	size int
}

func part1(line []string) {
	var size = GetTotalSize(line)
	fmt.Println("\tPart 1: ", size);
} 
// 43562874 wrong 
// 1886043 right

func part2(line []string) {
	var size = GetRemoveFolderSize(line)
	fmt.Println("\tPart 2: ", size);
}
// 3915660 too high
// 3842121 right

func Run() {
	absPath, _ := filepath.Abs("./day7/day7.txt")
	var multiLineInput = utils.ReadFile(absPath)
	part1(multiLineInput)
	part2(multiLineInput)
}

func CreateNode(line string, parent *dir) (dir) {
	var spliited = strings.Split(line, " ");
	var folderName = spliited[1];

	return dir{folderName, parent, []*dir{}, []file{}, 0};
}

func CreateFileInTree(line string)(file) {
	var spliited = strings.Split(line, " ");
	var size, err = strconv.Atoi(spliited[0]);
	if err != nil {
    // ... handle error
    panic(err)
	};
	var fileName = spliited[1];

	return file{fileName, size};
}

func IsFile(line string)(bool) {
	if (line[0:3] == "dir") {
		return false;
	} else {
		return true;
	}
}

func FindIndex(folders []*dir, searchString string)(int) {
	for idx, folder := range folders {
		if (folder.name == searchString) {
			return idx;
		}
	}

	return -1;
}

func IsCommand(line string)(bool) {
	if (line[0:1] == "$") {
		return true;
	}
	return false;
}

func GetCommand(line string)(string, string) {
	var splitted = strings.Split(line, " ");
	var command = splitted[1];
	var args = "";
	if (command == "cd") {
		args = splitted[2];
	}
	return command, args;
}

func ParseInput(input []string)(dir) {
	var rootNode = dir{"/", nil, []*dir{}, []file{}, 0};
	var currentNode = &rootNode;

	for _, line := range input {
		if (IsCommand(line)) {
			var command, args = GetCommand(line);
			if (command == "cd" && args == "..") {
				currentNode.parent.size = currentNode.parent.size + currentNode.size;
				currentNode = currentNode.parent;
			} else if (command == "cd" && args == "/") {
				continue;
			} else if (command == "cd") {
				var folderName = args;
				var index = FindIndex(currentNode.folders, folderName);
				if (index != -1) {
					currentNode = currentNode.folders[index];
				} else {
					fmt.Println("not panic", args);
				}
			}
		} else {
			if (IsFile(line)) {
				var file = CreateFileInTree(line);
				currentNode.files = append(currentNode.files, file)
				currentNode.size = currentNode.size + file.size;
			} else {
				var folder = CreateNode(line, currentNode)
				currentNode.folders = append(currentNode.folders, &folder)
			}
		}
	}

	return rootNode;
}


func PrintFolder(tree *dir, depth int) {
		fmt.Printf(strings.Repeat("  ", depth) + "- %s (dir size=%d) \n", tree.name, tree.size)
		for _, file := range tree.files  {
			fmt.Printf(strings.Repeat("  ", depth+1) + "- %s (file, size=%d) \n", file.name, file.size);
		}
		for _, folder := range tree.folders {
			PrintFolder(folder, depth + 1);
		};
}

// PART 1
var PART1_LIMIT = 100000;

func GetFolderOver100KSumSize(tree *dir, sum int) (int) {
	var currentSum = sum;

	if (tree.size < PART1_LIMIT) {
  	currentSum = tree.size + sum;
  };

	for _, folder := range tree.folders {
		 currentSum = GetFolderOver100KSumSize(folder, currentSum);
	};
 
  return currentSum;
}

func GetTotalSize(input []string) (int) {

	var tree = ParseInput(input);

	PrintFolder(&tree, 1);

	var size = GetFolderOver100KSumSize(&tree, 0);

	return size;
}

// PART 2
func GetDiskUsage(currentFolder *dir, sum int) (int) {
	var currentSum = sum;

  for _, file := range currentFolder.files {
  	currentSum = currentSum + file.size;
  }

	for _, folder := range currentFolder.folders {
		 currentSum = GetDiskUsage(folder, currentSum);
	};
 
  return currentSum;
}

func GetFolderListForRemoval(currentFolder *dir, target int, list []int) ([]int) {
	if (currentFolder.size > target) {
		list = append(list, currentFolder.size);
	};

	for _, folder := range currentFolder.folders {
		list = GetFolderListForRemoval(folder, target, list);
	};
 
  return list;
}

var STORAGE_SIZE = 70000000;
var NEED_SPACE = 30000000;

func GetRemoveFolderSize(input []string) (int) {
	var tree = ParseInput(input);

	var disk_usage = GetDiskUsage(&tree, 0);
	var free_space = STORAGE_SIZE - disk_usage;
	var deficiency = NEED_SPACE - free_space;

	var foldersToDelete = GetFolderListForRemoval(&tree, deficiency, []int{});

	sort.Ints(foldersToDelete)

	return foldersToDelete[0];
}


