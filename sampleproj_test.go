package main

import (
    "testing"
)

type counterTest struct{
    text string
    expected int
}

var counterTests = []counterTest{
    counterTest{"/....path.../test1.txt",15},
    counterTest{"/...path.../test2.txt",5},
    counterTest{"/...path.../test3.txt",10},
    }

func TestLineCounter(t *testing.T) {
    for _, cT:= range counterTests{
        actual,_:=lineCounter(cT.text)
        if actual != cT.expected{
            t.Errorf("lineCounter(%s): expected %d, actual %d",cT.text,
            cT.expected,actual)
        }
    }
} 