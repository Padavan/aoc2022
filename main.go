package main

import (
    "fmt"
    // "errors"
    // D1 "aoc2022/day1"
    // D2 "aoc2022/day2"
    // D3 "aoc2022/day3"
    // D4 "aoc2022/day4"
    // D5 "aoc2022/day5"
    // D6 "aoc2022/day6"
    // D7 "aoc2022/day7"
    // D8 "aoc2022/day8"
    // D9 "aoc2022/day9"
    // D10 "aoc2022/day10"
    // D11 "aoc2022/day11"
    // D12 "aoc2022/day12"
    // D13 "aoc2022/day13"
)

func b() {
    num, err := c();
    if (err) {
        fmt.Println(err);
    }
}

func c()(int) {
    err := errors.New("some error")
    // noterror := 69;
    panic(err);
    return 1
}

func a() {
    b();
    
}


func main() {
    fmt.Println("=== Advent of Go 2022 ===")
    // D1.Run();
    // D2.Run();
    // D3.Run();
    // D4.Run();
    // D5.Run();
    // D6.Run();
    // D7.Run();
    // D8.Run();
    // D9.Run();
    // D10.Run();
    // D11.Run();
    // D12.Run();
    // D13.Run();
    a();
}

