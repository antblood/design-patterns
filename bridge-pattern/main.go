package main

import (
	"fmt"

	"github.com/antblood/design-patterns/bridge-pattern/pkg/document"
)

func main() {
	// Create a report and a book.
	report := &document.Report{
		Title:   "Annual Report",
		Content: "This is the annual report content.",
	}

	book := &document.Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
	}

	// Create formatters.
	plainTextFormatter := &document.PlainTextFormatter{}
	htmlFormatter := &document.HTMLFormatter{}

	// Assign formatters and print contents.
	report.SetFormatter(plainTextFormatter)
	book.SetFormatter(plainTextFormatter)

	fmt.Println("Printing in Plain Text Format:")
	report.PrintContent()
	book.PrintContent()

	report.SetFormatter(htmlFormatter)
	book.SetFormatter(htmlFormatter)

	fmt.Println("\nPrinting in HTML Format:")
	report.PrintContent()
	book.PrintContent()
}
