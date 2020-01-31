package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var entries [][]string
	var balance, charges float64

	// Open file with charges
	file, err := os.Open("entries.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Go through the file and remove non-alphanumeric or decimals
	for scanner.Scan() {
		cleanRow := strings.Split(cleanString(scanner.Text()), " ")
		entries = append(entries, cleanRow)
	}

	// Go through entries and adjust balances
	for k, v := range entries {
		// Sets the first row as the starting balance and move to charges
		if k == 0 {
			balance, err = strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Original_Balance: %0.2f\n", balance)
			continue
		}
		charge, err := strconv.ParseFloat(v[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		charges += charge
		balance = balance - charge
		fmt.Printf("%v %v %v Balance %0.2f\n", v[0], v[1], v[2], balance)
	}
	fmt.Printf("Total expense %0.2f\n", charges)
	fmt.Printf("Average expense %0.2f\n", charges/float64(len(entries)-1))
}

func cleanString(entry string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9.]+")
	if err != nil {
		log.Fatal(err)
	}
	cleanString := reg.ReplaceAllString(entry, " ")
	return cleanString
}
