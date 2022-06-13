// this file shows how to take input from a file during LLD interview
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// get current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("error in getting the working directory: ", err.Error())
	} else {
		fmt.Println("working dir: ", wd)
	}

	// open the file
	file, err := os.Open(fmt.Sprintf("%v/input.txt", wd))
	if err != nil {
		log.Fatalf("error while opening file: %v", err.Error())
	}

	// define a scanner
	scanner := bufio.NewScanner(file)

	// split the input by words
	scanner.Split(bufio.ScanLines)

	// access the input file sentence by sentence
	// then access a sentence word by word
	// printOnConsole(scanner)
	writeOnOutputFile(wd+"/output.txt", scanner)
}

func writeOnOutputFile(path string, scanner *bufio.Scanner) {
	file, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		log.Fatalf("error while opening output file: %v", err.Error())
	}

	defer file.Close()

	// delete the content of the output file if any
	if err := os.Truncate(path, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

	// iterate through each line
	for scanner.Scan() {
		sentence := scanner.Text()
		words := strings.Split(sentence, " ")

		l := len(words)
		// iterate through each word
		for i := 0; i < l; i++ {
			if _, err = file.WriteString(words[i]); err != nil {
				log.Fatalf("error while writing to file: %v", err.Error())
			}

			// print the separator (SPACE or NEW_LINE)
			var separator string
			if i == l-1 {
				separator = "\n"
			} else {
				separator = " "
			}
			if _, err = file.WriteString(separator); err != nil {
				log.Fatalf("error while writing to file: %v", err.Error())
			}
		}
	}
}

func printOnConsole(scanner *bufio.Scanner) {
	for scanner.Scan() {
		sentence := scanner.Text()
		words := strings.Split(sentence, " ")
		for _, word := range words {
			fmt.Print(word, " ") // do not put a space after printing thr last word
		}
		fmt.Println()
	}
}
