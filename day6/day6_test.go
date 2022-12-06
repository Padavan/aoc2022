package day6

import "testing"

// PART 1
func TestGetMarkerIndex1(t *testing.T){

    var got = GetMarkerIndex("bvwbjplbgvbhsrlpgdmjqwftvncz", 4)
    var want = 5

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetMarkerIndex2(t *testing.T){

    var got = GetMarkerIndex("nppdvjthqldpwncqszvftbrmjlhg", 4)
    var want = 6

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetMarkerIndex3(t *testing.T){

    var got = GetMarkerIndex("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4)
    var want = 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}


func TestGetMarkerIndex4(t *testing.T){

    var got = GetMarkerIndex("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4)
    var want = 11

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}


func TestGetMarkerIndex1Part2(t *testing.T){

    var got = GetMarkerIndex("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14)
    var want = 19

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}
