package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	filePtr := flag.String("f", "input", "Input file if not 'input'")

	flag.Parse()
	readFile, err := os.Open(*filePtr)

	if err != nil {
		fmt.Println("Fatal:", err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// Insert code here
	var num int = 0
	for _, v := range lines {
		s := ConvertFromSnafu(v)
		fmt.Printf("%s --> %d\n", v, s)
		num += s
	}
	fmt.Printf("Total: %d\n", num)
	fmt.Printf("Converted: %s\n", ConvertToSnafu(num))

	fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
