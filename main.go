package main

import (
	"fmt"
	"os"
)

func main() {
	// for loop that goes forever
	for {
		fmt.Println("Choose a CSV file:")
		fmt.Println("1. NBAData22_23.csv")
		fmt.Println("2. Salary_dataset.csv")
		fmt.Println("3. 50_Startups.csv")
		fmt.Println("4. Exit")

		// get input form the user
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		// get the choice and perform the things with that data
		switch choice {
		case 1:
			processCSV("NBAData22_23.csv")
		case 2:
			processCSV("Salary_dataset.csv")
		case 3:
			processCSV("50_Startups.csv")
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}

func processCSV(fileName string) {
	data := LoadData(fileName)
	coef := train(data)

	// test the results on the training data
	results := test(coef, data)

	numRow, _ := results.Dims()

	// for every row, display the predictions and actual results
	for i := 0; i < numRow; i++ {
		// display the predictions
		fmt.Printf("Prediction for %dth row of data is: \n", i+1)
		fmt.Printf("%.2f\n", results.At(i, 0))

		// display the actual results
		fmt.Println("Actual result is: ")
		fmt.Printf("%.2f\n", results.At(i, 1))
		fmt.Println()
	}
}
