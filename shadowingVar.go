package main 

import "fmt"

func main() {
//The statement n := 1 declares a new variable which shadows the original n throughout the scope of the if statement.
//To reuse n from the outer block, write n = 1 instead.
    n := 0
    if true {
        n = 1
        n++
    }
	//%v == value of n
	//%T == data type of n
    fmt.Printf("%v, %T", n, n) // 0
}
