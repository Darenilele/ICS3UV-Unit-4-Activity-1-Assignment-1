/**
 * @author Daren Ileleji
 * @version 1.0.0
 * @date 2025-12-09
 * @fileoverview This program will compute the average marks for given number of marks and calculate average 
 */

// Variables
let numberOfMarks: number = 0;
let stringOfMarks: string = "";
let total: number = 0;
// The below variables are ones for when the program gets the amount of marks the user gave
let markAmountGiven: number = 0;
let markStringGiven: string = "";


// Input
stringOfMarks = prompt("How many marks will you enter today?") || "0";
numberOfMarks = parseInt(stringOfMarks);

// Processing
for (let markCounter = 1; markCounter <= numberOfMarks; markCounter++) {
  markStringGiven = prompt(`Enter mark ${markCounter}:`) || "0";
  markAmountGiven = parseInt(markStringGiven);
  total = markAmountGiven + total;
}

total = total / numberOfMarks 

// Output
console.log(`You have entered ${numberOfMarks} marks`)
if (total <= 49) {
  console.log("This student is failing")
} else if (total >= 50 && total <= 69) {
  console.log("This student has performance just under average")
} else if (total >= 70 && total <= 79) {
  console.log("This student has average performance")
} else if (total >= 80) {
  console.log("This student has above average performance")
}

console.log("\nDone.")