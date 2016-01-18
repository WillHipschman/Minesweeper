package main

import "fmt"
import "errors"
import "math/rand"

var(
    INVALID_BOMBS = "The number of bombs must be greater than 0 and less than width * height."
    DIMENSION_TOO_SMALL = "The dimension cannot be less than 1."
    DIMENSION_TOO_LARGE = "The dimension cannot be greater than 80 for width and 40 for height."
    HEIGHT_WIDTH_NOT_SET = "Height and Width must be set for this operation to complete."
    YOU_LOSE = "Boom! You lose!"
    BOMB = '*'
)

type setAction func (val int)(error)

type Board struct{
    width           int
    height          int
    numOfBombs      int
    field           [][]rune
    exploredCells   int
    rowToExplore    int
    colToExplore    int
}

func (board *Board) Explore() (bool, error) {
    return board.explore(board.rowToExplore, board.colToExplore)
}
    
func (board *Board) explore(row, col int) (bool, error) {
    
    if board.field[row][col] == BOMB {
        return false, errors.New(YOU_LOSE)
    }
    
    if (board.field[row][col] != 0){
        return false, nil
    }
    
    count := board.countOfBombs(row, col)
    
    if count == 0 {
        board.field[row][col] = -1
        
        // If there are no bombs then we have to explore the neighbors
        board.exploreNeighbors(row, col)
        
    } else {
        board.field[row][col] = rune(count)
    }
    
    board.exploredCells++
    
    return true, nil
}

func (board *Board) IsSolved() bool {
    return (board.exploredCells + board.numOfBombs) == (board.height * board.width)
}

func (board *Board) Print(){
    fmt.Println()
    fmt.Println("The board is:")
    
    for i := range board.field {
        for j := range board.field[i] {
            if (board.field[i][j] == 0){
                fmt.Print("X")   
            }else if (board.field[i][j] > 0 && board.field[i][j] != BOMB){
                fmt.Print(int(board.field[i][j]))
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
    if err := board.validateHeightAndWidthAreSet(); err != nil{
        return err
    }
    
    // A cell   containing a * is a bomb
    // A cell   containing 0 is undiscovered and may be explored
    // A cell containi  ng 1 is discovered and may not be explored    
    board.field = make([][]rune, board.height)
    
    for i := range board.field {
        board.field[i] = make([]rune, board.width)
    }
    
    for numberToGenerate := board.numOfBombs; numberToGenerate > 0;{
        row := rand.Intn(board.height)
            col := rand.Intn(board.width)
                
        if (board.field[row][col] == 0){
            board.field[row][col] = BOMB
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

func (board *Board) exploreNeighbors(row, col int) {    
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            if (i == 0 && j == 0){
                continue;
            }
            
            if !board.validatePosition(row + i, col + j){
                continue
            }
            
            if (board.field[row + i][col + j] == '*') {
                panic("Internal Error: Found unexepected bomb.")
            }
            
            board.explore(row + i, col + j)
        }
    }
}

func (board *Board) countOfBombs(row, col int) int{
    countOfBombs := 0 
    
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            if (i == 0 && j == 0){
                continue;
            }
            
            if !board.validatePosition(row + i, col + j){
                continue
            }
            
            if (board.field[row + i][col + j] == '*') {
                countOfBombs++;
            }
        }
    }
    
    return countOfBombs
}

func (board *Board) validatePosition(row, col int) bool{
    if (row < 0 || row >= board.height) {
        return false;   
    }
    
    if (col < 0 || col >= board.width) {
        return false
    }
    
    return true
}