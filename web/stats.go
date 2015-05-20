package web
import (
    "net/http"
    "strconv"
    "fmt"
    "github.com/go-read-a-book/valsorter"
    "github.com/go-read-a-book/store"
    "encoding/json"
    "errors"
)

func Stats(w http.ResponseWriter, r *http.Request) {
    n, err := strconv.Atoi(r.URL.Query().Get("N"))

    //if no values are set return 5 values (well if 5 are available ofcourse)
    if (err != nil) {
        n = 5
    }

    dat := make(map[string]interface{})

    warnings := [2]string{}
    warningsLen := 0

    dat["total"] = int(len(store.WM))
    highestWords, err := HighestWords(n) //slice it!

    if err != nil {
        warningsLen ++
        warnings[warningsLen] = err.Error()
    }

    highestChars, err := HighestChars(n) //slice this one too!

    if err != nil {
        warningsLen ++
        warnings[warningsLen] = err.Error()
    }

    dat[fmt.Sprint("top_", n ,"_words")] = highestWords
    dat[fmt.Sprint("top_", n ,"_letters")] = highestChars

    j, err := json.Marshal(dat)

    if (err == nil) {
        w.Write([]byte(j));
    } else {
        w.Write([]byte("Could not handle it"))
        w.Write([]byte(err.Error()))
    }
}

func HighestWords(n int) ([]string, error) {
    wSorter := valsorter.NewValSorter(store.WM)
    warning := nil

    wn := len(store.WM)
    if (wn > n) {
        wn = n
    } else if (wn != n) {
        warning = errors.New(fmt.Sprint("Could not get ", n , " words, only  ", wn, " available"))
    }

    topWords := make([]string, wn)

    for i :=0 ; i < n ; i++ {
        if wn > i {
            topWords[i] = wSorter.Keys[wSorter.Len() - i - 1]
        }
    }

    return topWords, warning
}

func HighestChars(n int) ([]string, error) {
    wSorter := valsorter.NewValSorter(store.CM)
    warning := nil

    cn := len(store.CM)
    if (cn > n) {
        cn = n
    } else if (cn != n) {
        warning = errors.New(fmt.Sprint("Could not get ", n , " chars, only  ", cn, " available"))
    }

    topChars := make([]string, cn)

    for i :=0 ; i < n ; i++ {
        if cn > i {
            topChars[i] = wSorter.Keys[wSorter.Len() - i - 1]
        }
    }

    return topChars, warning
}