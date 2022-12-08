package day8

import "testing"

var testInput = []string{""};

func TestPart1(t *testing.T){

    var got = len(testInput)
    var want = 1

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}


