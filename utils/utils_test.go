package utils

import "testing"

var testCollection = []string{"s", "h", "i", "t"};

func TestIndexOfSuccess(t *testing.T){

    var got = IndexOf("t", testCollection);
    var want = 3

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestIndexOfFail(t *testing.T){

    var got = IndexOf("u", testCollection);
    var want = -1

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

