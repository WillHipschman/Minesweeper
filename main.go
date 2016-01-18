package main

import "fmt"
import "github.com/willhipschman/minesweeper/resources"

func main() {
    printWelcomeMessage()
    
    board := new(Board)
    board.Init()
}

func printWelcomeMessage(){
    fmt.Println()
    fmt.Println("Hello and welcome to Minesweeper!")
    fmt.Println()
    fmt.Println(resources.Broom) 
    fmt.Println()
}