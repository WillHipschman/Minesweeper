package main

import "fmt"
import "github.com/willhipschman/minesweeper/resources"

func main() {
    printWelcomeMessage()
    
    board := new(Board)
    setup(board)
    
    for solved := board.IsSolved(); !solved;{
        board.Print()
        promptExplore(board)
        board.Explore()
        solved = board.IsSolved()
    } 
}

func printWelcomeMessage(){
    fmt.Println()
    fmt.Println("Hello and welcome to Minesweeper!")
    fmt.Println()
    fmt.Println(resources.Broom) 
    fmt.Println()
}

func printSuccessMessage(){
    fmt.Println()
    fmt.Println("Great job solving Minesweeper!")
    fmt.Println()
}

func setup(board *Board){
    promptForSetting(board.SetWidth, "Enter the width of the board.")
    promptForSetting(board.SetHeight, "Enter the height of the board.")
    promptForSetting(board.SetBombs, "Enter the number of bombs.")
    
    if err := board.Setup(); err != nil{
        err = board.Setup()
    }
}

func promptExplore(board *Board){
    fmt.Println()
    promptForSetting(board.SetRowToExplore, "Enter the row to  explore")
    promptForSetting(board.SetColToExplore, "Enter the col to  explore")
}

func promptForSetting(action setAction, message string){
    temp := -1
    firstTime := true
    for err := action(temp); err != nil; {
        if(!firstTime){
            fmt.Println(err.Error())
        }else{
            firstTime = false
        }
        
        fmt.Println(message)
        _, err2 := fmt.Scan(&temp) 
        
        if(err2 != nil){
            panic(err2)
        }
        
        err = action(temp)
    }
}