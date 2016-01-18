package main

import "testing"
import "fmt"
import "github.com/willhipschman/minesweeper/resources"

func TestSetRowToExplore(t *testing.T){
    board := new(Board)
    board.SetHeight(20)
    board.SetWidth(20)
    
    assertIsEqual(t, board.rowToExplore, 0, "Row should have been %d but is %d.")
    
    failIfNull(t, board.SetRowToExplore(resources.MinInt), "MinInt is an illegal row.")
    failIfNull(t, board.SetRowToExplore(-1), "-1 is an illegal row.")
    failIfNull(t, board.SetRowToExplore(0), "0 is an illegal row.")
    failIfNull(t, board.SetRowToExplore(21), "21 is an illegal row.")
    failIfNull(t, board.SetRowToExplore(resources.MaxInt), "MaxInt is an illegal row.")
    
    failIfNotNull(t, board.SetRowToExplore(1), "1 should be a legal value.")
    assertIsEqual(t, board.rowToExplore, 0, "Row should have been %d but is %d.")
    
    failIfNotNull(t, board.SetRowToExplore(17), "17 should be a legal value.")
    assertIsEqual(t, board.rowToExplore, 16, "Row should have been %d but is %d.")
    
    failIfNotNull(t, board.SetRowToExplore(20), "20 should be a legal value.")
    assertIsEqual(t, board.rowToExplore, 19, "Row should have been %d but is %d.")
}

func TestSetColToExplore(t *testing.T){
    board := new(Board)
    board.SetHeight(20)
    board.SetWidth(20)
    
    assertIsEqual(t, board.colToExplore, 0, "Col should have been %d but is %d.")
    
    failIfNull(t, board.SetColToExplore(resources.MinInt), "MinInt is an illegal col.")
    failIfNull(t, board.SetColToExplore(-1), "-1 is an illegal col.")
    failIfNull(t, board.SetColToExplore(0), "0 is an illegal col.")
    failIfNull(t, board.SetColToExplore(21), "21 is an illegal col.")
    failIfNull(t, board.SetColToExplore(resources.MaxInt), "MaxInt is an illegal col.")
    
    failIfNotNull(t, board.SetColToExplore(1), "1 should be a legal value.")
    assertIsEqual(t, board.colToExplore, 0, "Col should have been %d but is %d.")
    
    failIfNotNull(t, board.SetColToExplore(17), "17 should be a legal value.")
    assertIsEqual(t, board.colToExplore, 16, "Col should have been %d but is %d.")
    
    failIfNotNull(t, board.SetColToExplore(20), "20 should be a legal value.")
    assertIsEqual(t, board.colToExplore, 19, "Col should have been %d but is %d.")
}

func TestIsSolved(t *testing.T){
    board := new(Board)
    board.SetHeight(2)
    board.SetWidth(2)
    board.SetBombs(2)
    
    assertIsFalse(t, board.IsSolved(), "Board should not be solved on init")
    board.exploredCells++
    assertIsFalse(t, board.IsSolved(), "Board should not be solved when there is 1 unexplored cell")
    board.exploredCells++
    
    assertIsTrue(t, board.IsSolved(), "Board should be solved")
}

func TestValidateHeightAndWidthAreSet(t *testing.T){
    board := new(Board)
    
    failIfNull(t, board.validateHeightAndWidthAreSet(), "ValidateHeightAndWidthAreSet should fail before init.")
    
    board.SetHeight(20)
    board.SetWidth(20)
    
    failIfNotNull(t, board.validateHeightAndWidthAreSet(), "ValidateHeightAndWidthAreSet should succeed after init.")
}

func TestSetup(t *testing.T){
    board := new(Board)
    
    failIfNull(t, board.Setup(), "GenerateBombPositions should fail before Setup.")
    
    board.SetHeight(20)
    board.SetWidth(20)
    board.SetBombs(10)
    board.Setup()
    numBombs, numNonBombs := countTypes(t, board)
    assertIsEqual(t, numBombs, 10, "Number of bombs should be 10 but was %d.")
    assertIsEqual(t, numNonBombs, 390, "Number of non-bombs should be 390 but was %d.")
    
    board.SetBombs(300)
    board.Setup()
    numBombs, numNonBombs = countTypes(t, board)
    assertIsEqual(t, numBombs, 300, "Number of bombs should be 300 but was %d.")
    assertIsEqual(t, numNonBombs, 100, "Number of non-bombs should be 100 but was %d.")
}

func TestSetBombs(t *testing.T){
    board := new(Board)
    
    failIfNull(t, board.SetBombs(10), "SetBombs should fail when height and weight aren't initialized.")
    
    board.SetHeight(20)
    board.SetWidth(20)
    
    failIfNull(t, board.SetBombs(-1), "-1 is an illegal value for bombs.")
    failIfNull(t, board.SetBombs(0), "0 is an illegal value for bombs.")
    failIfNull(t, board.SetBombs(400), "400 is an illegal value for bombs.")
    failIfNull(t, board.SetBombs(401), "401 is an illegal value for bombs.")
    
    failIfNotNull(t, board.SetBombs(1), "1 should be a legal value.")
    assertIsEqual(t, board.numOfBombs, 1, "Bomb num should have been %d but is %d.")
    
    failIfNotNull(t, board.SetBombs(163), "163 should be a legal value.")
    assertIsEqual(t, board.numOfBombs, 163, "Bomb num should have been %d but is %d.")
}

func TestSetHeight(t *testing.T){
    board := new(Board)
    
    assertIsEqual(t, board.height, 0, "Height should have been %d but is %d.")
    
    failIfNull(t, board.SetHeight(resources.MinInt), "MinInt is an illegal height.")
    failIfNull(t, board.SetHeight(-1), "-1 is an illegal height.")
    failIfNull(t, board.SetHeight(0), "0 is an illegal height.")
    failIfNull(t, board.SetHeight(1), "1 is an illegal height.")
    failIfNull(t, board.SetHeight(41), "41 is an illegal height.")
    failIfNull(t, board.SetHeight(resources.MaxInt), "MaxInt is an illegal height.")
    
    failIfNotNull(t, board.SetHeight(2), "2 should be a legal value.")
    assertIsEqual(t, board.height, 2, "Height should have been %d but is %d.")
    
    failIfNotNull(t, board.SetHeight(27), "27 should be a legal value.")
    assertIsEqual(t, board.height, 27, "Height should have been %d but is %d.")
    
    failIfNotNull(t, board.SetHeight(40), "40 should be a legal value.")
    assertIsEqual(t, board.height, 40, "Height should have been %d but is %d.")
}

func TestSetWidth(t *testing.T){
    board := new(Board)
    
    assertIsEqual(t, board.width, 0, "Width should have been %d but is %d.")
    
    failIfNull(t, board.SetWidth(resources.MinInt), "MinInt is an illegal width.")
    failIfNull(t, board.SetWidth(-1), "-1 is an illegal width.")
    failIfNull(t, board.SetWidth(0), "0 is an illegal width.")
    failIfNull(t, board.SetWidth(1), "1 is an illegal width.")
    failIfNull(t, board.SetWidth(81), "81 is an illegal width.")
    failIfNull(t, board.SetWidth(resources.MaxInt), "MaxInt is an illegal width.")
    
    failIfNotNull(t, board.SetWidth(2), "2 should be a legal value.")
    assertIsEqual(t, board.width, 2, "Width should have been %d but is %d.")
    
    failIfNotNull(t, board.SetWidth(37), "37 should be a legal value.")
    assertIsEqual(t, board.width, 37, "Width should have been %d but is %d.")
    
    failIfNotNull(t, board.SetWidth(80), "80 should be a legal value.")
    assertIsEqual(t, board.width, 80, "Width should have been %d but is %d.")
}

func assertIsFalse(t *testing.T, actual bool, message string){
    if(actual){
        t.Errorf(message)
    }
}

func assertIsTrue(t *testing.T, actual bool, message string){
    if(!actual){
        t.Errorf(message)
    }
}

func assertIsEqual(t *testing.T, actual int, expected int, message string){
    if(actual != expected){
        t.Errorf(fmt.Sprintf(message, expected, actual))
    }
}

func failIfNotNull(t *testing.T, err error, message string){
    if(err != nil){
        t.Errorf(message)
    }
}

func failIfNull(t *testing.T, err error, message string){
    if(err == nil){
        t.Errorf(message)
    }
}

func countTypes(t *testing.T, board *Board) (bombs int, nonBombs int){
    for i := 0; i < board.width; i++{
        for j := 0 ; j < board.height; j++{
            if board.privateField[i][j] == 0{
                nonBombs++
            }else if board.privateField[i][j] == '*'{
                bombs++
            }else{
                t.Errorf(fmt.Sprintf("Unexpected value %d found in field.", board.privateField[i][j]))
            }
        }
    }
    
    return bombs, nonBombs
}