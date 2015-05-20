package read
//testing the reader, mostly trying to see if we don't get invalid chars
import (
    "net/http"
    "io/ioutil"
    "strings"
    "regexp"
    "log"
    "github.com/go-read-a-book/store"
)

var regAlphaWhitespace *regexp.Regexp
var regAlpha *regexp.Regexp

func init() {
    //create regex to get only chars & whitespaces/newlines
    rW, err := regexp.Compile("[^a-z\\s]+")
    if err != nil {
        log.Fatal(err)
    }
    //one to only allow chars
    regAlphaWhitespace = rW

    r, err := regexp.Compile("[^a-z]")
    if err != nil {
        log.Fatal(err)
    }
    regAlpha = r
}

//web rest endpoint
func WordsAndChars(w http.ResponseWriter, r *http.Request) {
    go parseBody(r)
}

//just so we can wrap it in a goroutine
func parseBody(r *http.Request) {
    bytes, err := ioutil.ReadAll(r.Body);

    if err == nil {
        ReadWordsAndCharsFrom(bytes)
    }
}

//actual reading code
func ReadWordsAndCharsFrom(bytes []byte) {
    str := string(bytes)
    //only store lowercase
    str = regAlphaWhitespace.ReplaceAllString(strings.ToLower(str), "")
    //gets all words (contains duplicates)
    wCounts := strings.Fields(str)

    for _, word := range wCounts {
        store.WM[word] ++
    }

    //string is clean we need it clean
    bytes = []byte(regAlpha.ReplaceAllString(str, ""))

    //iterate over array and add to char array, I feel like this is not the best way to do it
    for i := 0; i < len(bytes); i++ {
        char := string(bytes[i])
        //regex allows spaces to be there
        if char != "\\s" {
            store.CM[char] ++
        }
    }
}