package main

import "testing"
import "fmt"
import "github.com/willhipschman/minesweeper/resources"

func TestValidateHeightAndWidthAreSet(t *testing.T){
    board := new(Board)
    
    failIfNull(t, board.validateHeightAndWidthAreSet(), "ValidateHeightAndWidthAreSet should fail before init.")
    
    board.SetHeight(20)
    board.SetWidth(20)
    
    failIfNotNull(t, board.validateHeightAndWidthAreSet(), "ValidateHeightAndWidthAreSet should succeed after init.")
}

func TestGenerateBombPositions(t *testing.T){
    board := new(Board)
    
    board.SetHeight(20)
    board.SetWidth(20)
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
    assertIsEqual(t, board.numOfBombs, 1, "Bomb num should have been %d but is %d")
    
    failIfNotNull(t, board.SetBombs(163), "163 should be a legal value.")
    assertIsEqual(t, board.numOfBombs, 163, "Bomb num should have been %d but is %d")
}

func TestSetHeight(t *testing.T){
    board := new(Board)
    
    assertIsEqual(t, board.height, 0, "Height should have been %d but is %d")
    
    failIfNull(t, board.SetHeight(resources.MinInt), "MinInt is an illegal height.")
    failIfNull(t, board.SetHeight(-1), "-1 is an illegal height.")
    failIfNull(t, board.SetHeight(0), "0 is an illegal height.")
    failIfNull(t, board.SetHeight(1), "1 is an illegal height.")
    failIfNull(t, board.SetHeight(41), "41 is an illegal height.")
    failIfNull(t, board.SetHeight(resources.MaxInt), "MaxInt is an illegal height.")
    
    failIfNotNull(t, board.SetHeight(2), "2 should be a legal value.")
    assertIsEqual(t, board.height, 2, "Height should have been %d but is %d")
    
    failIfNotNull(t, board.SetHeight(27), "27 should be a legal value.")
    assertIsEqual(t, board.height, 27, "Height should have been %d but is %d")
    
    failIfNotNull(t, board.SetHeight(40), "40 should be a legal value.")
    assertIsEqual(t, board.height, 40, "Height should have been %d but is %d")
}

func TestSetWidth(t *testing.T){
    board := new(Board)
    
    assertIsEqual(t, board.width, 0, "Width should have been %d but is %d")
    
    failIfNull(t, board.SetWidth(resources.MinInt), "MinInt is an illegal width.")
    failIfNull(t, board.SetWidth(-1), "-1 is an illegal width.")
    failIfNull(t, board.SetWidth(0), "0 is an illegal width.")
    failIfNull(t, board.SetWidth(1), "1 is an illegal width.")
    failIfNull(t, board.SetWidth(81), "81 is an illegal width.")
    failIfNull(t, board.SetWidth(resources.MaxInt), "MaxInt is an illegal width.")
    
    failIfNotNull(t, board.SetWidth(2), "2 should be a legal value.")
    assertIsEqual(t, board.width, 2, "Width should have been %d but is %d")
    
    failIfNotNull(t, board.SetWidth(37), "37 should be a legal value.")
    assertIsEqual(t, board.width, 37, "Width should have been %d but is %d")
    
    failIfNotNull(t, board.SetWidth(80), "80 should be a legal value.")
    assertIsEqual(t, board.width, 80, "Width should have been %d but is %d")
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