package main

import "fmt"
import "errors"

var(
    INVALID_BOMBS = "The number of bombs must be greater than 0 and less than width * height."
    DIMENSION_TOO_SMALL = "The dimension cannot be less than 0."
    DIMENSION_TOO_LARGE = "The dimension cannot be greater than 80 for width and 40 for height."
    HEIGHT_WIDTH_NOT_SET = "You can declare the number of bombs until height and width have been set."
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
    if(numOfBombs <= 0){
        return errors.New(INVALID_BOMBS)
    }
    
    if(board.width == 0  || board.height == 0){
        return errors.New(HEIGHT_WIDTH_NOT_SET)
    }
    
    if(numOfBombs >= board.width * board.height){
        return errors.New(INVALID_BOMBS)
    }
    
    board.numOfBombs = numOfBombs
    return nil
}

func (board *Board) SetHeight(height int)(error){
    if(height <= 0){
        return errors.New(DIMENSION_TOO_SMALL)
    }
    
    if(height > 40){
        return errors.New(DIMENSION_TOO_LARGE)
    }
    
    board.height = height
    return nil
}

func (board *Board) SetWidth(width int)(error){
    if(width <= 0){
        return errors.New(DIMENSION_TOO_SMALL)
    }
    
    if(width > 80){
        return errors.New(DIMENSION_TOO_LARGE)
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