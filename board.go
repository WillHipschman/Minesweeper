package main

import "fmt"
import "errors"

var(
    TOO_MANY_BOMBS = "The number of bombs cannot be greater than the number of spaces on the board."
    INVALID_DIMENSIONS = "The dimension was less than 0."
)

type setAction func (val int)(error)

type Board struct{
    width int
    height int
    numOfBombs int
}

func (board *Board) Init(){
    setBoardAttribute(board.SetWidth, "Enter the width of the board.")
    setBoardAttribute(board.SetHeight, "Enter the height of the board.")
    setBoardAttribute(board.SetBombs, "Enter the number of bombs.")
}

func (board *Board) SetBombs(numOfBombs int)(error){
    if(numOfBombs < 0){
        return errors.New(INVALID_DIMENSIONS)
    }
    
    if(numOfBombs > board.width * board.height){
        return errors.New(TOO_MANY_BOMBS)
    }
    
    board.numOfBombs = numOfBombs
    return nil
}

func (board *Board) SetHeight(height int)(error){
    if(height <= 0){
        return errors.New(INVALID_DIMENSIONS)
    }
    
    board.height = height
    return nil
}

func (board *Board) SetWidth(width int)(error){
    if(width <= 0){
        return errors.New(INVALID_DIMENSIONS)
    }
    
    board.width = width
    return nil
}

func setBoardAttribute(action setAction, message string){
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