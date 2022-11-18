package day1

import "testing"

func TestGetDistanceFromCenter(t *testing.T){

    var got = GetDistanceFromCenter(2, -2)
    var want = 4

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
