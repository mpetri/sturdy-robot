package main

import "testing" 

   

func TestlineWordCounter(t *testing.T) {
    counterTests := []counterTest struct{
    fName string
    expectedLineCount int
    expectedWordCount int
}
    {
    counterTest{"test1.txt",36,13,},
    
    counterTest{"test2.txt",43,23,},
    
    counterTest{"test3.txt",22,4,},
    }

    for _, cT:= range counterTests{
        actualLineCount, actualWordCount,_:=lineWordCounter(cT.fName,"GUTENberG")
        if actualLineCount != cT.expectedLineCount||actualWordCount != cT.expectedWordCount {
            t.Errorf("lineWordCounter(%s): expectedLineCount %d, expectedWordCount %d, actualLineCount %d
            ,actualWordCount %d",cT.fName,cT.expectedLineCount,cT.expectedWordCount,actualLineCount,actualWordCount)
        }
    }
} 
