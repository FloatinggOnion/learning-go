# Golang Expense Tracker
*Author: Jesse-Paul Osemeke*
This is a simple command-line application written in Go to track expenses. It allows you to add, list, summarize, and delete expenses stored in a JSON file.


## Features
- **Add Expenses**: Add a new expense with a description, amount, and the current date.
- **List Expenses**: Display all recorded expenses with their ID, date, description, and amount.
- **Summarize Expenses**: Calculate the total expenses for a specific month or all expenses.
- **Delete Expenses**: Remove an expense by its ID.


## Prerequisites
- Go installed on your system (version 1.16 or later).
- Basic knowledge of using the terminal/command line.


## How It Works
The application uses a JSON file (expenses.json) to store expense data. Each expense includes:

**date**: The date the expense was added (in YYYY-MM-DD format).
**description**: A brief description of the expense.
**amount**: The amount spent.

The commands and flags allow you to interact with the application to manage your expenses.

### Commands
1. Add an Expense
Add a new expense to the tracker.

```zsh
./expense-tracker add -description="Lunch" -amount=15.50
```

-description: A description of the expense (default: "some expense").
-amount: The amount of the expense (default: 0.0).

2. List All Expenses
List all recorded expenses.

```zsh
./expense-tracker list
```

This will display all expenses in the following format:

```zsh
# ID    Date        Description     Amount
# 1     2025-04-07  Groceries       12.00
# 2     2025-04-07  Gym             3.50
# 3     2025-04-07  Laptop          400.00
```

3. Summarize Expenses
Summarize the total expenses for a specific month or all expenses.

Summarize All Expenses

```zsh
./expense-tracker summary
```

Summarize for a Specific Month
```zsh
./expense-tracker summary -month=4
```
`-month`: The month (1–12) for which to calculate the total expenses.

4. Delete an Expense
Delete an expense by its ID.

```zsh
./expense-tracker delete -id=2
```

`-id`: The ID of the expense to delete.


## How to Run
- Build the Application: Compile the Go program into an executable file.

- Run the Application: Use the compiled executable to run commands.

    - Replace `<command>` with one of the supported commands (add, list, summary, delete).

- Example Usage:

    - Add an expense:
    ```zsh
    ./expense-tracker add -description="Coffee" -amount=4.50
    ```
    - List all expenses:
    ```zsh
    ./expense-tracker list
    ```
    - Summarize expenses for April:
    ```zsh
    ./expense-tracker summary -month=4
    ```
    - Delete an expense:
    ```zsh
    ./expense-tracker delete -id=1
    ```

## File Structure
`main.go`: The main application logic.
`expenses.json`: The JSON file where expenses are stored.
`Readme.md`: Documentation for the application.

## Notes
*Ensure the expenses.json file exists in the same directory as the executable. If it doesn't exist, the application will create it automatically.*
*The application assumes valid input for commands and flags. Invalid input may result in errors.*