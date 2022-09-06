package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Sample for File Read")
	filePath := "/Users/rohitphatak/Rohit/MyResults.csv"
	scanFile(filePath)
}
func scanFile(filePath string) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l := sc.Text() // GET the line string
		// Print the length of string
		fmt.Println("Line size : ", len(l))
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}
