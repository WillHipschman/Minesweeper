package main

import "testing"
import "fmt"
import "github.com/willhipschman/minesweeper/resources"

func TestSetWidth(t *testing.T){
    board := new(Board)
    
    failIfNull(t, board.SetWidth(resources.MinInt), "MinInt is an illegal width.")
    failIfNull(t, board.SetWidth(-1), "-1 is an illegal width.")
    failIfNull(t, board.SetWidth(0), "0 is an illegal width.")
    
    
    failIfNotNull(t, board.SetWidth(1), "1 should be a legal value.")
    assertIsEqual(t, board.width, 1, "Width should have been %d but is %d")
    
    failIfNotNull(t, board.SetWidth(resources.MaxInt), "MaxInt should be a legal value.")
    assertIsEqual(t, board.width, resources.MaxInt, "Width should have been %d but is %d")
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