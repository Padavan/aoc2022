package day4

import "testing"

var testInput = []string{"2-4,6-8", "2-3,4-5", "5-7,7-9", "2-8,3-7", "6-6,4-6", "2-6,4-8" };

func TestPart1(t *testing.T){

    var got = GetFullOverlapCount(testInput)
    var want = 2

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestPart2(t *testing.T){

    var got = GetPartialOverlapCount(testInput)
    var want = 4

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}


