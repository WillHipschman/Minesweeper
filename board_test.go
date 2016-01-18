package main

import "testing"
import "fmt"
import "github.com/willhipschman/minesweeper/resources"


///
/// Some tests contain comments indicating board positions to help reason about tests.
/// * indicates a bomb.
/// . indicates an unexplored cell.
/// X indicates an explored cell with no adjacent bombs.
/// {1-8} indicates the number of adjacent bombs.
///

func TestE2E(t *testing.T){
    board := new(Board)
    
    // **..
    // ....
    // ....
    // ..*.
    testCountOfBombs(t, board, 4, 4)
    board.privateField[0][0] = BOMB
    board.privateField[0][1] = BOMB
    board.privateField[3][2] = BOMB
    board.numOfBombs = 3
    
    
    board.explore(1, 3)
    
    // **1X
    // ..1X
    // ..11
    // ..*.
    
    // discovered
    assertIsEqual(t, int(board.displayField[0][3]), -1, "Position 0,3 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[0][2]), 1, "Position 0,2 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[1][3]), -1, "Position 1,3 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[1][2]), 1, "Position 1,2 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[2][3]), 1, "Position 2,3 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[2][2]), 1, "Position 2,2 should contain %d but contained %d")

    // undiscovered
    assertIsEqual(t, int(board.displayField[0][1]), 0, "Position 0,1 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[1][1]), 0, "Position 1,1 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[2][1]), 0, "Position 2,1 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[3][1]), 0, "Position 3,1 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[3][2]), 0, "Position 3,2 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[3][3]), 0, "Position 3,3 should contain %d but contained %d")
    
    board.explore(2, 0)
    
    // **1X
    // 221X
    // X111
    // X1*.
    
    // discovered
    assertIsEqual(t, int(board.displayField[1][0]), 2, "Position 1,0 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[1][1]), 2, "Position 1,1 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[2][0]), -1, "Position 2,0 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[2][1]), 1, "Position 1,1 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[3][0]), -1, "Position 2,1 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[3][1]), 1, "Position 3,1 should contain %d but contained %d")
    
    // undiscovered
    assertIsEqual(t, int(board.displayField[0][0]), 0, "Position 0,0 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[0][1]), 0, "Position 0,1 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[3][2]), 0, "Position 3,2 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[3][3]), 0, "Position 3,3 should contain %d but contained %d")
    
    // sanity--make sure nothing else changed
    assertIsEqual(t, int(board.displayField[1][3]), -1, "Position 1,3 should contain %d but contained %d")
    assertIsEqual(t, int(board.displayField[0][2]), 1, "Position 0,2 should contain %d but contained %d")
    
    board.explore(3, 3)
    
    // **1X
    // 221X
    // X111
    // X1*1
    
    assertIsEqual(t, int(board.displayField[3][3]), 1, "Position 3,3 should contain %d but contained %d")
    assertIsTrue(t, board.IsSolved(), "Board should be solved")
}

func TestCountOfBombs(t *testing.T){
    board := new(Board)
    
    // *.
    // ..
    // ..
    testCountOfBombs(t, board, 3, 2)
    board.privateField[0][0] = BOMB
    assertIsEqual(t, board.countOfBombs(0,1), 1, "3*2: There should be %d bomb but there were %d.")
    assertIsEqual(t, board.countOfBombs(1,0), 1, "3*2: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(1,1), 1, "3*2: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(2,0), 0, "3*2: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(2,1), 0, "3*2: There should be %d bombs but there were %d.")
    
    // **.
    // ...
    testCountOfBombs(t, board, 2, 3)
    board.privateField[0][0] = BOMB
    board.privateField[0][1] = BOMB
    assertIsEqual(t, board.countOfBombs(0,2), 1, "2*3: There should be %d bomb but there were %d.")
    assertIsEqual(t, board.countOfBombs(1,0), 2, "2*3: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(1,1), 2, "2*3: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(1,2), 1, "2*3: There should be %d bombs but there were %d.")
    
    // .*
    // **
    testCountOfBombs(t, board, 2, 2)
    board.privateField[0][1] = BOMB
    board.privateField[1][0] = BOMB
    board.privateField[1][1] = BOMB
    assertIsEqual(t, board.countOfBombs(0,0), 3, "2*2: There should be %d bombs but there were %d.")
    
    // ***
    // *..
    // ...
    testCountOfBombs(t, board, 3, 3)
    board.privateField[0][0] = BOMB
    board.privateField[0][1] = BOMB
    board.privateField[0][2] = BOMB
    board.privateField[1][0] = BOMB
    assertIsEqual(t, board.countOfBombs(1,1), 4, "3*3: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(1,2), 2, "3*3: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(2,0), 1, "3*3: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(2,1), 1, "3*3: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(2,2), 0, "3*3: There should be %d bombs but there were %d.")
    
    // ***..
    // *.*..
    // ***..
    // .....
    // .....
    testCountOfBombs(t, board, 4, 5)
    board.privateField[0][0] = BOMB
    board.privateField[0][1] = BOMB
    board.privateField[0][2] = BOMB
    board.privateField[1][0] = BOMB
    board.privateField[1][2] = BOMB
    board.privateField[2][0] = BOMB
    board.privateField[2][1] = BOMB
    board.privateField[2][2] = BOMB
    assertIsEqual(t, board.countOfBombs(1,1), 8, "5*5: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(0,3), 2, "5*5: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(0,5), 0, "5*5: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(2,3), 2, "5*5: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(2,5), 0, "5*5: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(3,0), 2, "5*5: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(3,1), 3, "5*5: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(3,2), 2, "5*5: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(3,3), 1, "5*5: There should be %d bombs but there were %d.")
    assertIsEqual(t, board.countOfBombs(3,4), 0, "5*5: There should be %d bombs but there were %d.")
}

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

func testCountOfBombs(t *testing.T, board *Board, height, width int){
    board.SetHeight(height)
    board.SetWidth(width)
    if err := board.Setup(); err != nil {
        t.Errorf(err.Error())
    }
}