package day2

import "testing"

var testInput = []string{ "A Y", "B X", "C Z" };

func TestPart1(t *testing.T){

    var got = GetTotalScore(testInput)
    var want = 15

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestPart2(t *testing.T){

    var got = GetTotalScore2(testInput)
    var want = 12

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetOutcomeIndexWin(t *testing.T){
    var got = GetOutcomeIndex(0, 1)
    var want = 2

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetOutcomeIndexLose(t *testing.T){
    var got = GetOutcomeIndex(0, 2)
    var want = 0

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetOutcomeIndexDraw(t *testing.T){
    var got = GetOutcomeIndex(2, 2)
    var want = 1

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetShapeScoreScissors(t *testing.T){
    var got = GetShapeScore(2)
    var want = 3

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetShapeScoreRock(t *testing.T){
    var got = GetShapeScore(0)
    var want = 1

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestOutcomeScoreLose(t *testing.T){
    var got = GetOutcomeScore(0)
    var want = 0

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}


func TestOutcomeScoreDraw(t *testing.T){
    var got = GetOutcomeScore(1)
    var want = 3

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestOutcomeScoreWin(t *testing.T){
    var got = GetOutcomeScore(2)
    var want = 6

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetPlayerHandIndexLose(t *testing.T){
    var got = GetPlayerHandIndex(0, 0)
    var want = 2

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetPlayerHandIndexLoseAgain(t *testing.T){
    var got = GetPlayerHandIndex(2, 0)
    var want = 1

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetPlayerHandIndexDraw(t *testing.T){
    var got = GetPlayerHandIndex(1, 1)
    var want = 1

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetPlayerHandIndexWin(t *testing.T){
    var got = GetPlayerHandIndex(2, 2)
    var want = 0

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestGetPlayerHandIndexWinAgain(t *testing.T){
    var got = GetPlayerHandIndex(1, 2)
    var want = 2

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}