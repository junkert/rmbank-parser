package rmbank_parser

import (
	"bufio"
	"fmt"
	"golang.org/x/text/currency"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
    panic(e)
	}
}

type Transaction struct {
	date    string
	who     string
	amount  float64
	txType  string            // "DEPOSIT" or "WITHDRAWAL"
}

type MonthlyLedger struct {
	startDate       string
	endDate         string
	startingBalance float64
	endBalance      float64
	transactions    []Transaction
}

func main() {
	var (
		lineCounter   int64
		done          bool
		ledger        []MonthlyLedger
	)

	filePath := "../tmp/output.txt"
	file, err := os.Open(filePath)
	defer file.Close()
	check(err)

	fileReader := bufio.NewReader(file)

	done = false
	lineCounter = 1
	curMonth := new(MonthlyLedger)
	for {

		if lineCounter % 5 == 0 { fmt.Println("Completed lines: %d", lineCounter) }
		inputLine, err = fileReader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				done = true
			}
		} else {
			check(err)
		}

		splitLine := strings.Fields(strings.TrimLeftFunc(inputLine, unicode.IsSpace)
		if len(splitLine) == 0 {
			continue
		}
		if len(splitLine) >= 5 &&
			strings.HasPrefix(splitLine[0], "BEGINNING") &&
			strings.HasPrefix(splitLine[1], "BALANCE") {

				curMonth.startDate, err       = splitLine[2]
				check(err)
				curMonth.startingBalance, err = strconv.Atoi(splitLine[3])
				check(err)
			}
		if strings.HasPrefix(cleanLine, "asdf"){

		}


		if done {
			break
		}
		lineCounter++
		// Create new struct if needed
		if strings.HasPrefix(splitLine[0], "MEMBER") && strings.HasPrefix(splitLine[1], "FDIC") {
			ledger = append(ledger, curMonth)
			curMonth = new(MonthlyLedger)
		}
}