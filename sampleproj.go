package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

func checkTxt(searchDir string) {
	filepath.Walk(searchDir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			matched, err := regexp.MatchString(".txt", f.Name())
			if err == nil && matched {
				numLines, err := lineCounter(path)
				if err != nil {
					log.Fatal(err)
				}
				writeFile(f.Name(), numLines)
			}
		}
		return nil
	})
}

func lineCounter(fname string) (int, error) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var counter int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counter++
	}
	return counter, scanner.Err()
}

func writeFile(fname string, numLines int) {
	file, err := os.OpenFile("proj.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := []byte(fname + "," + strconv.Itoa(numLines) + "\n")
	_, err = file.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
}

func writeHeader() {
	err := ioutil.WriteFile("proj.csv", []byte("File Name, Number of Lines\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	start := time.Now()

	writeHeader()
	searchDir := os.Args[1]
	checkTxt(searchDir)

	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)

}
