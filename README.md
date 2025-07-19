# Expense Tracker CLI

## A powerful and minimal command-line tool to track your expenses, budgets, and summaries right from the terminal.

## Features

- Add new expenses with description, amount, category, and date
- List expenses by category, date, or month
- Update existing expenses by ID
- Delete specific expenses
- Delete all expenses
- Export expenses to CSV or JSON
- Manage monthly budgets
- View summaries by category or month
- Run system diagnostics
- Initialize configuration
- Built using [Cobra](https://github.com/spf13/cobra) & [Viper](https://github.com/spf13/viper)
- Built for [Roadmap.sh](https://roadmap.sh/projects/expense-tracker)

---

## Installation

> Requires Go 1.20+

```bash
git clone https://github.com/Moukhtar-youssef/Expense-Tracker.git
cd Expense-Tracker
go install
```

Once installed, run:

```bash
expense-tracker help
```

---

## Usage

### Add Expense

```bash
expense_tracker add --description "Lunch" --amount 12.50 --category "Food" --date 2025-07-11
```

### List Expenses

```bash
expense_tracker list --category "Transport" --from 2025-07-01 --to 2025-07-10
```

### Update Expense

```bash
expense_tracker update --id 3 --amount 20 --description "Dinner"
```

### Delete Expense

```bash
expense_tracker delete --id 5
```

### Delete All Expenses

```bash
expense_tracker delete all
```

### Export Expenses

```bash
expense_tracker export --format csv --output expenses.csv --month 7
```

### Set a Budget

```bash
expense_tracker budget set --month 7 --amount 500
```

### Check Budget

```bash
expense_tracker budget check --month 7
```

### View Summary

```bash
expense_tracker summary --month 7 --category "Food"
```

### Run Doctor

```bash
expense_tracker doctor
```

### Initialize Config

```bash
expense_tracker init
```

---

## Available Commands

| Command        | Description                    |
| -------------- | ------------------------------ |
| `add`          | Add a new expense              |
| `list`         | List expenses                  |
| `update`       | Update an existing expense     |
| `delete`       | Delete an expense              |
| `delete all`   | Delete all expenses            |
| `export`       | Export expenses to CSV or JSON |
| `budget set`   | Set a monthly budget           |
| `budget check` | Check if budget is exceeded    |
| `summary`      | View summary of expenses       |
| `doctor`       | Check system configuration     |
| `init`         | Setup config and data file     |
| `version`      | Display current CLI version    |
