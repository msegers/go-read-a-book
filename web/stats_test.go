package web
import (
    "testing"
    "github.com/go-read-a-book/store"
    "github.com/go-read-a-book/web"
)

func TestTopResultTextReturnsExpected(t *testing.T) {
    store.WM = map[string]int{"chickensandwich": 9001, "chickenlesssandwich": 9000, "orval":10, "stroganoff"}
    //since map is random we're testing this a couple of times, it could randomly get the correct vals

    for i := 0; i<50; i++ {
        a, _ := web.HighestWords(4)
        if a[0] != "chickensandwich" {
            t.Fatalf("chickensandwich was not the top word in reading", a)
        }
    }
}

func TestTopResultCharReturnsExpected(t *testing.T) {
    store.WM = map[string]int{"z": 9001, "b": 9000, "x":10, "n"}
    //since map is random we're testing this a couple of times, it could randomly get the correct vals

    for i := 0; i<50; i++ {
        a, _ := web.HighestWords(4)
        if a[0] != "z" {
            t.Fatalf("z was not the top char in reading", a)
        }
    }
}
