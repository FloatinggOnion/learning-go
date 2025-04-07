package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)


type Expense struct {
	Date string`json:"date"`
	Description string `json:"description"`
	Amount float64 `json:"amount"`
}
var filename = "expenses.json"

func main() {

	// flags
	descriptionFlag := flag.String("description", "some expense", "Description of the expense")
	amountFlag := flag.Float64("amount", 0.0, "Amount of the expense")
	idFlag := flag.Int("id", 0, "ID of the expense to delete")
	monthFlag := flag.Int("month", 0, "Month to summarise")
	
	fmt.Println("Golang Expense Tracker")

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: add or list")
		return
	}

	command := os.Args[1]
	
	flag.CommandLine.Parse(os.Args[2:])
	
	switch command {
		case "add":
			fmt.Println(*descriptionFlag, *amountFlag)
			if *descriptionFlag == "" || *amountFlag <= 0 {
				fmt.Println("Please provide a valid description and amount.")
				return
			}
			add(*descriptionFlag, *amountFlag)
		
		case "list":
			list()
		
		case "summary":
			summary(*monthFlag)

		case "delete":
			delete(*idFlag)

		default:
			fmt.Println("Unknown command:", command)
			fmt.Println("Available commands: add, list, summary, delete")
			return
	}
}

func add(description string, amount float64) {
    // Open the file for reading and writing
    file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Read existing expenses
    var expenses []Expense
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&expenses)
    if err != nil && err.Error() != "EOF" {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    // Add the new expense
    newExpense := Expense{
		Date: time.Now().Format("2006-01-02"),
        Description: description,
        Amount:      amount,
    }
    expenses = append(expenses, newExpense)

    // Truncate the file and write the updated expenses
    file.Truncate(0)
    file.Seek(0, 0)
    encoder := json.NewEncoder(file)
    err = encoder.Encode(expenses)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }

    // Print success message
    fmt.Println("# Expense added successfully!")
    fmt.Printf("# Description: %s, Amount: %.2f\n", description, amount)
}

func open_file() ([]Expense) {
	var expenses []Expense

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&expenses)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	return (expenses)
}

func list() {
	fmt.Println("Listing all expenses...")

	expenses := open_file()

	fmt.Println("Expenses:")
	fmt.Println("# ID\t Date\t\t Description\t Amount")
	for id, expense := range expenses {
		fmt.Printf("# %d\t %s\t %s\t %.2f\t\n", id+1, expense.Date, expense.Description, expense.Amount)
	}

	return
}

func summary(month ...int) {
	fmt.Println("Summary of expenses...")

	expenses := open_file()
	amount := 0.0


	if month == nil || month[0] == 0 {
		fmt.Println("No month provided. Showing all expenses.")
		for _, expense := range expenses {
			amount += expense.Amount
		}

		fmt.Printf("# Total expenses: %.2f\n", amount)
	} else {
		fmt.Println("Month provided:", month[0])
		for _, expense := range expenses {
			expenseMonth, _ := time.Parse("2006-01-02", expense.Date)
			if int(expenseMonth.Month()) == month[0] {
				amount += expense.Amount
			}
		}

		fmt.Printf("Total expenses for %s: %.2f\n", time.Month(month[0]), amount)
	}

	

	return
}

func delete(id int) {
	fmt.Println("Deleting an expense...")

	expenses := open_file()

	if id < 1 || id > len(expenses) {
        fmt.Println("Invalid ID. Please provide a valid expense ID.")
        return
    }

	expenses = append(expenses[:id-1], expenses[id:]...)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(expenses)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Expense deleted successfully!")
	
	return
}