package main

import (
    "flag"
    "net/http"
    "fmt"
    "log"
    "github.com/go-read-a-book/web"
    "github.com/go-read-a-book/read"
)

//configurable ports
var tport = flag.Int("tport", 5555, "Listening port for text parsing")
var wport = flag.Int("wport", 8080, "Port for Restfull API refering to it as webport in code")

func init() {
    //parse flags
    flag.Parse()
}

func main() {
    //start in goroutine to allow multiple instead of getting stuck at the first
    //port to post data to
    go SetTextPort()
    //port to get stats from
    SetWebPort()
}

func SetTextPort() {
    http.HandleFunc("/", read.WordsAndChars)

    txtPort := fmt.Sprint(":" , *tport);
    fmt.Println("textReadingPort = " + txtPort)
    log.Fatal(http.ListenAndServe(txtPort, nil))
}

func SetWebPort() {
    http.HandleFunc("/stats", web.Stats)

    webPort := fmt.Sprint(":" , *wport);
    fmt.Println("Webport = " + webPort)
    log.Fatal(http.ListenAndServe(webPort, nil))
}

