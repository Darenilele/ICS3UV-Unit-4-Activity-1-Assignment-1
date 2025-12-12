/**
 * @author Daren Ileleji
 * @version 1.0.0
 * @date 2025-12-09
 * @fileoverview This program will compute the average marks for given number of marks and calculate average 
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Variables
	var numberOfMarks int = 0
	var stringOfMarks string = ""
	var total float64 = 0
	var markAmountGiven float64 = 0
	var markStringGiven string = ""

	reader := bufio.NewReader(os.Stdin)

	// Input
	fmt.Print("How many marks will you enter today? ")
	stringOfMarks, _ = reader.ReadString('\n')
	stringOfMarks = strings.TrimSpace(stringOfMarks)
	numberOfMarks, _ = strconv.Atoi(stringOfMarks)

	// Processing
	for markCounter := 1; markCounter <= numberOfMarks; markCounter++ {
		fmt.Print("Enter mark ", markCounter, ": ")
		markStringGiven, _ = reader.ReadString('\n')
		markStringGiven = strings.TrimSpace(markStringGiven)
		markAmountGiven, _ = strconv.ParseFloat(markStringGiven, 64)
		total = total + markAmountGiven
	}

	total = total / float64(numberOfMarks)

	// Output
	fmt.Println("You have entered", numberOfMarks, "marks. The student's average is", total)
	if total <= 49 {
		fmt.Println("This student is failing")
	} else if total >= 50 && total <= 69 {
		fmt.Println("This student has performance just under average")
	} else if total >= 70 && total <= 79 {
		fmt.Println("This student has average performance")
	} else if total >= 80 {
		fmt.Println("This student has above average performance")
	}

	fmt.Println("\nDone.")
}
