package store

//just for storing the values, trying to separate
var WM map[string]int = make(map[string]int)   //words array, key for word so it has to be unique
var CM map[string]int = make(map[string]int)   //character array, key for char so it has to be unique ^

//mainly used for unit tests
func Clear() {
    WM = make(map[string]int)
    CM = make(map[string]int)
}