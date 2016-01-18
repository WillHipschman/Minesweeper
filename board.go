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
    width           int
    height          int
    numOfBombs      int
    privateField    [][]rune
    displayField    [][]rune
    exploredCells   int
    rowToExplore    int
    colToExplore    int
}

func (board *Board) Explore() error {
    return nil
}

func (board *Board) IsSolved() bool {
    return (board.exploredCells + board.numOfBombs) == (board.height * board.width)
}

func (board *Board) Print(){
    fmt.Println()
    fmt.Println("The board is:")
    
    for i := range board.displayField {
        for j := range board.displayField {
            if (board.displayField[i][j] == 0){
                fmt.Print("X")   
            }else if (board.displayField[i][j] > 0){
                fmt.Print(board.displayField[i][j])   
            } else {
                fmt.Print(".")   
            }
        }
        fmt.Println()
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

func (board *Board) SetWidth(width int) error {
    if (width <= 1){
        return errors.New(DIMENSION_TOO_SMALL)
    }
    
    if (width > 80){
        return errors.New(DIMENSION_TOO_LARGE)
    }
    
    board.width = width
    return nil
}

func (board *Board) SetColToExplore(col int) error{
    // move from 1 based to 0 base index
    col = col -1
    
    if col < 0 {
        return errors.New(DIMENSION_TOO_SMALL)
    }
    
    if col >= board.width {
        return errors.New(DIMENSION_TOO_LARGE)
    }
    
    board.colToExplore = col
    return nil
}

func (board *Board) SetRowToExplore(row int) error{
    // move from 1 based to 0 base index
    row = row -1
    
    if row < 0 {
        return errors.New(DIMENSION_TOO_SMALL)
    }
    
    if row >= board.height {
        return errors.New(DIMENSION_TOO_LARGE)
    }
    
    board.rowToExplore = row
    return nil
}

func (board *Board) Setup() error {
    
    // A cell   containing a * is a bomb
    // A cell   containing 0 is undiscovered and may be explorer
    //  A cell containi  ng 1 is discovered and may not be explored    
    board.privateField = make([][]rune, board.height)
    board.displayField = make([][] rune, board.height)
         
    for i := range board.privateField {
        board.privateField[i] = make([]rune, board.width)
        board.displayField[i] = make([]rune, board.width)
    }
    
    if err := board.validateHeightAndWidthAreSet(); err != nil{
        return err
    }
    
    for numberToGenerate := board.numOfBombs; numberToGenerate > 0;{
            row := rand.Intn(board.height)
                col := rand.Intn(board.width)
                
        if (board.privateField[row][col] == 0){
            board.privateField[row][col] = '*'
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