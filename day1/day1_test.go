package day1

import "testing"

var testInput = []string{ "1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000" };

func TestPart1GettingMaxGroupValue(t *testing.T){

    var got = GetMax(testInput)
    var want = 24000

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}


func TestPart2GettingMaxGroupValue(t *testing.T){

    var got = GetLargestN(testInput, 3)
    var want = 45000

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
