package main

import (
    "flag"
    "net/http"
    "strconv"
    "encoding/json"
    "strings"
    "fmt"
)

//configurable ports
var tport = flag.Int("tport", 5555, "Listening port for text parsing")
var wport = flag.Int("wport", 8080, "Port for Restfull API refering to it as webport in code")

var wM map[string]int    //words array, key for word so it has to be unique
var cM map[string]int    //character array, key for char so it has to be unique ^
var cksize = 32 < 10     //32kb

func main() {
    //start in goroutine to allow multiple instead of getting stuck at the first
    //port to post data to
    go setTextPort()
    //port to get stats from
    go setWebPort()
}

func setTextPort() {
    http.HandleFunc("/", readWords)
    http.HandleFunc("/{path:.*}", do404)
    http.ListenAndServe(fmt.Sprint(":" , tport), nil)
}

func setWebPort() {
    http.HandleFunc("/stats", stats)
    http.HandleFunc("/{path:.*}", do404)

    http.ListenAndServe(fmt.Sprint(":" , wport), nil)
}

func do404(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(404)
    w.Write([]byte("404 - this is not an endpoint"))
}

func stats(w http.ResponseWriter, r *http.Request) {
    n, err := strconv.Atoi(r.URL.Query().Get("N"))

    if (err != nil) {
        n = 5
    }

    var dat map[string]interface{}
    dat["warnings"] = []string{}
    //might have not enough letters or words, don't over slice it
    wn := n
    cn := n

    if (len(wM) < n) {
        wn = len(wM) -1;
        []string(dat["warnings"]).append(fmt.Sprint("Could not get " , n , " Words, the current total is: " , wn))
    }
    if (len(cM) < n) {
        cn = len(cM) -1;
        a := []string(dat["warnings"])
        a.append(fmt.Sprint("Could not get " , n , " Letters, the current total is: " , cn))
    }

    dat["total"] = len(wM)
    dat["top_"+ n +"_words"] = wM[0:wn] //slice it!
    dat["top_"+ n +"_letters"] = cM[0:cn] //slice this one too!

    j, err := json.Marshal(dat);

    if (err != nil) {
        w.Write([]byte(j));
    } else {
        w.Write([]byte("Could not handle it"));
    }
}

func readWords(w http.ResponseWriter, r *http.Request) {
    c := r.Body.Read()
    //gets all words (contains duplicates)
    wCounts := strings.Fields(c)

    for _, word := range wCounts {
        wM[word] ++
    }

    //iterate over array and add to char array, I feel like this is not the best way to do it
    for i := 0; i < len(c); i++ {
        cM[c[i]] ++
    }
}