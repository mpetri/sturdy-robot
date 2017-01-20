package main

import "testing"

type counterTest struct {
        Name              string
        expectedLineCount int
        expectedWordCount int
}

var tests = []counterTest{
        {"test1.txt", 36, 13},
        {"test2.txt", 43, 23},
        {"test3.txt", 22, 4},
}

func TestlineWordCounter(t *testing.T) {
        for _, cT := range tests {
                actualLineCount, actualWordCount, _ := lineWordCounter(cT.Name, "GUTENberG")
                if actualLineCount != cT.expectedLineCount || actualWordCount != cT.expectedWordCount {
                        t.Errorf("lineWordCounter(%s): expectedLineCount %d, expectedWordCount %d, actualLineCount %d ,actualWordCount %d", cT.Name, cT.expectedLineCount, cT.expectedWordCount, actualLineCount, actualWordCount)
                }
        }
}
