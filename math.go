package main

import "C"

//export Add
func Add(x, y C.int) C.int {
    return C.int(x + y)
}
func main(){}