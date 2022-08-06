package ioHandler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	outputFileName = "output.txt"
	inputFileName  = "input.txt"
)

var (
	InputFilePath  string
	OutputFilePath string
)

func init() {
	wd, err := getCurrentWorkingDirectory()
	if err != nil {
		log.Fatalln("error in getting the working directory: ", err.Error())
	}
	InputFilePath = fmt.Sprintf("%v/ioHandler/%v", wd, inputFileName)
	OutputFilePath = fmt.Sprintf("%v/ioHandler/%v", wd, outputFileName)

	// delete the content of the output file if any
	if err := os.Truncate(OutputFilePath, 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}
}

func getCurrentWorkingDirectory() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return wd, nil
}

// Write writes to the output file
func Write(filePath string, text string) {
	file, err := os.OpenFile(OutputFilePath, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		log.Fatalf("error in opening output file: %v", err.Error())
	}

	defer file.Close()

	if _, err = file.WriteString(fmt.Sprintf("%v\n", text)); err != nil {
		log.Fatalf("error while writing to output file: %v", err.Error())
	}
}

func TakeInput(snakeCnt *int, ladderCnt *int, snakes *[][]int, ladders *[][]int,
	playerCnt *int, playerNames *[]string,
) {
	file, err := os.Open(InputFilePath)
	if err != nil {
		log.Fatalf("error while opening input file: %v", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// snakes input
	scanner.Scan()
	snCnt, _ := strconv.Atoi(scanner.Text())
	*snakeCnt = snCnt
	var snks [][]int
	for i := 0; i < *snakeCnt; i++ {
		scanner.Scan()
		line := scanner.Text()
		nums := strings.Split(line, " ")
		snakeFrom, _ := strconv.Atoi(nums[0])
		snakeTo, _ := strconv.Atoi(nums[1])
		snks = append(snks, []int{snakeFrom, snakeTo})
	}
	*snakes = snks

	// ladder input
	scanner.Scan()
	lddrCnt, _ := strconv.Atoi(scanner.Text())
	*ladderCnt = lddrCnt
	for i := 0; i < *ladderCnt; i++ {
		scanner.Scan()
		line := scanner.Text()
		nums := strings.Split(line, " ")
		ladderFrom, _ := strconv.Atoi(nums[0])
		ladderTo, _ := strconv.Atoi(nums[1])
		lddrs := append(*snakes, []int{ladderFrom, ladderTo})
		*ladders = lddrs
	}

	// player info input
	scanner.Scan()
	plyrCnt, _ := strconv.Atoi(scanner.Text())
	*playerCnt = plyrCnt
	for i := 0; i < *playerCnt; i++ {
		scanner.Scan()
		name := scanner.Text()
		plrs := append(*playerNames, name)
		*playerNames = plrs
	}
}
