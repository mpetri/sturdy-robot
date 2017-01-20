package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type txtFile struct {
	fname     string
	lineCount int
	wordCount int
}

func lineWordCounter(fname, keyWord string) (int, int, error) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineCount int
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		lineCount++
	}
	wordCount := findKeyWord(lines, keyWord)
	return lineCount, wordCount, scanner.Err()
}

func findKeyWord(lines []string, keyWord string) int {
	wordCount := make(map[string]int)
	for _, line := range lines {
		words := strings.FieldsFunc(line, Split)
		for _, word := range words {
			wordLow := strings.ToLower(word)
			(wordCount[wordLow])++
		}
	}
	keyWordLow := strings.ToLower(keyWord)
	if count, ok := wordCount[keyWordLow]; ok {
		return count
	}
	return 0
}

func Split(r rune) bool {
	return r == '.' || r == '"' || r == '?' || r == ',' || r == ' ' || r == '-' || r == '/'
}

func printFile(fp *os.File, word string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if !f.IsDir() {
			matched, err := regexp.MatchString(".txt", f.Name())
			if err == nil && matched {
				tF := new(txtFile)
				tF.fname = f.Name()
				tF.lineCount, tF.wordCount, err = lineWordCounter(path, word)
				if err != nil {
					log.Fatal(err)
				}
				writeToFile(fp, tF)
			}
		}
		return nil
	}
}

func writeToFile(file *os.File, tf *txtFile) {
	fmt.Fprintf(file, "%s,%d,%d\n", tf.fname, tf.lineCount, tf.wordCount)
}

func writeHeader(file *os.File) {
	fmt.Fprintf(file, "File Name, Number of Lines, Number of Words\n")
}
func safe_open_file() *os.File {
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
func close_file(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	start := time.Now()

	var filedir, keyWord string
	flag.StringVar(&filedir, "filedir", "", "Directory you wish to access")
	flag.StringVar(&keyWord, "word", "gutenberg", "A specific word you wish to find")
	flag.Parse()
	if filedir == "" || keyWord == "" {
		log.Fatal("not enough argument")
	}

	fp := safe_open_file()
	writeHeader(fp)
	err := filepath.Walk(filedir, printFile(fp, keyWord))
	if err != nil {
		log.Fatal(err)
	}
	close_file(fp)

	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}
