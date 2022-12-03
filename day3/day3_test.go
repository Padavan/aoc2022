package day3

import "testing"

var testInput = []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw" };

func TestPart1(t *testing.T){

    var got = GetTotalItemPriority(testInput)
    var want = 157

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestPart2(t *testing.T){

    var got = GetTotalGroupPriority(testInput)
    var want = 70

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetGetScorea(t *testing.T){

    var got = GetScore("a")
    var want = 1

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetGetScoreZ(t *testing.T){

    var got = GetScore("Z")
    var want = 52

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestFindCommonItem(t *testing.T){

    var got = FindCommonItem("abc", "defa")
    var want = "a"

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestFindCommonItemInThree(t *testing.T){

    var got = FindCommonItemInThree("abc", "defa", "gika")
    var want = "a"

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}


