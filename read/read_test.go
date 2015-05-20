package read
import (
    "testing"
    "github.com/go-read-a-book/store"
)

var txtContainingTopWordChickenSandWichTenTimes string =
"ChickenSandwich Frank Zappa, Who was eating those fine burgers ChickenSandwich, " +
"Would you rather have a ChickenSandwich? Yeah we can use marks 'n stuff " +
"ChickenSandwich is my favorite kind of ChickenSandwich especially Cheese ChickenSandwich " +
"Chicken Sandwich is a different combo though, ChickenSandwich ChickenSandwich ChickenSandwich ChickenSandwich "

func TestIfValuesAreBeingStored(t *testing.T) {
    store.Clear()
    ReadWordsAndCharsFrom([]byte("abcdefghijklmnopqrstuvwxyz"))
    if len(store.WM) != 1 {
        t.Fatalf("Word count is not 1, count: %d", len(store.WM), store.WM)
    }

    if len (store.CM) != 26 {
        t.Fatalf("Char count is not 26, count: %d", len(store.CM), store.CM)
    }
}

func TestOnlyAlphaNumericValuesAreBeingStored(t *testing.T) {
    store.Clear()
    ReadWordsAndCharsFrom([]byte("!@#$%^&*()_+1234567890-=`~[]{}\\|'\";:,.<>/?"))

    if len(store.WM) > 0 {
        t.Fatalf("Word count is bigger than 0, count: %d", len(store.WM), store.WM)
    }
    if len(store.CM) > 0 {
        t.Fatalf("Character count is bigger than 0, count: %d", len(store.CM), store.CM)
    }
}

func TestTotalWordCountMatchesReading(t *testing.T) {
    store.Clear()
    ReadWordsAndCharsFrom([]byte(txtContainingTopWordChickenSandWichTenTimes))

    if store.WM["chickensandwich"] != 10 {
        t.Fatalf("chickensandwich count was not matching 10, count: %d", store.WM["chickensandwich"])
    }

    if store.CM["c"] != 37 {
        t.Fatalf("c count was not matching 37, count: %d", store.CM["c"])
    }
}

func TestCapitalLettersAreIgnoredAndNotAddedToCount(t *testing.T) {
    store.Clear()
    ReadWordsAndCharsFrom([]byte("FRANK frank Frank\nfRank fRANK fRaNk"))

    if store.WM["frank"] != 6 {
        t.Fatalf("frank count was not matching 6, count: %d", store.WM["frank"])
    }

    if len(store.CM) != 5 {
        t.Fatalf("Different capitalized franks did not return a value of 5", store.CM)
    }

}