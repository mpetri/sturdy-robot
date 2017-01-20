package main

import(
    "flag"
    "fmt"
    "os"
    "log"
    "path/filepath"
	"regexp"
    "bufio"
    "time"
)

type txtFile struct{
    fname string
    numLines int
    //wordCount int   
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
func printFile(fp *os.File) filepath.WalkFunc{
    return func(path string, f os.FileInfo, err error) error{
        if err != nil {
            log.Print(err)
            return nil
        }
        if !f.IsDir() {
			matched, err := regexp.MatchString(".txt", f.Name())
			if err == nil && matched {
                tF := new(txtFile)
                tF.fname = f.Name()
                tF.numLines,err = lineCounter(path)
				if err != nil {
					log.Fatal(err)
				}
                writeToFile(fp, tF)
			}
		}
		return nil
    }
}

func  writeToFile(file *os.File, tf *txtFile){
    fmt.Fprintf(file, "%s,%d,%d\n",tf.fname,tf.numLines)
}

func writeHeader(file *os.File) {
    fmt.Fprintf(file,"File Name, Number of Lines\n")
}
func safe_open_file() *os.File{
    file, err := os.Create("result.csv")
    if err != nil{
        log.Fatal(err)
    }
    return file
}
func close_file(file *os.File){
    err:=file.Close()
    if err!=nil{
        log.Fatal(err)
    }
}

func main(){
    start := time.Now()
    
    var filedir string
    flag.StringVar(&filedir,"filedir","","Directory you wish to access")
    //flag.StringVar(&wordPtr,"word","gutenberg","A specific word you wish to find")
    //flag.Parse()
    //filedir := flag.Arg[1]
    //word:=flag.Args(2)
    flag.Parse()
    if filedir == ""{
        log.Fatal("not enough argument")
    }

    fp := safe_open_file()
    writeHeader(fp)
    err:=filepath.Walk(filedir, printFile(fp))
    if err!=nil{
        log.Fatal(err)
    }
    close_file(fp)

    elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

