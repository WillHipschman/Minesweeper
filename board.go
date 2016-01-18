package main

import "fmt"
import "errors"
import "math/rand"

var(
    INVALID_BOMBS = "The number of bombs must be greater than 0 and less than width * height."
    DIMENSION_TOO_SMALL = "The dimension cannot be less than 1."
    DIMENSION_TOO_LARGE = "The dimension cannot be greater than 80 for width and 40 for height."
    HEIGHT_WIDTH_NOT_SET = "Height and Width must be set for this operation to complete."
)

type setAction func (val int)(error)

type Board struct{
    width int
    height int
    numOfBombs int
    field [][]rune
}

func (board *Board) Init(){
    setBoardAttribute(board.SetWidth, "Enter the width of the board.")
    setBoardAttribute(board.SetHeight, "Enter the height of the board.")
    setBoardAttribute(board.SetBombs, "Enter the number of bombs.")
    
    if err := board.initBombs(); err != nil{
        board.Init()
    }
}

func (board *Board) SetBombs(numOfBombs int)(error){
    if (numOfBombs <= 0){
        return errors.New(INVALID_BOMBS)
    }
    
    if err := board.validateHeightAndWidthAreSet(); err != nil {
        return err
    }
    
    if (numOfBombs >= board.width * board.height){
        return errors.New(INVALID_BOMBS)
    }
    
    board.numOfBombs = numOfBombs
    return nil
}

func (board *Board) SetHeight(height int)(error){
    if (height <= 1){
        return errors.New(DIMENSION_TOO_SMALL)
    }
    
    if (height > 40){
        return errors.New(DIMENSION_TOO_LARGE)
    }
    
    board.height = height
    return nil
}

func (board *Board) SetWidth(width int)(error){
    if (width <= 1){
        return errors.New(DIMENSION_TOO_SMALL)
    }
    
    if (width > 80){
        return errors.New(DIMENSION_TOO_LARGE)
    }
    
    board.width = width
    return nil
}

func (board *Board) initBombs() error{
    
    // A cell containing a * is a bomb
    // A cell containing 0 is undiscovered and may be explorer
    // A cell containing 1 is discovered and may not be explored
    board.field = make([][]rune, board.width, board.height)
    
    if err := board.generateBombPositions(); err != nil{
        return err
    }
    
    return nil
}

func (board *Board) generateBombPositions() error{
    if err := board.validateHeightAndWidthAreSet(); err != nil{
        return err
    }
    
    for numberToGenerate := board.numOfBombs; numberToGenerate > 0;{
        row := rand.Int()
        col := rand.Int()
        
        if (board.field[row][col] == 0){
            board.field[row][col] = '*'
            numberToGenerate--
        }
    }
    
    return nil
}

func (board *Board) validateHeightAndWidthAreSet() error{
    if(board.width == 0  || board.height == 0){
        return errors.New(HEIGHT_WIDTH_NOT_SET)
    }
    
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