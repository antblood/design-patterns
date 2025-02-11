package document

import "fmt"

// Book is another refined abstraction.
type Book struct {
	Title     string
	Author    string
	formatter Formatter
}

func (b *Book) SetFormatter(formatter Formatter) {
	b.formatter = formatter
}

func (b *Book) PrintContent() {
	if b.formatter == nil {
		fmt.Println("No formatter set")
		return
	}
	fmt.Println(b.formatter.Format("Book: " + b.Title + " by " + b.Author))
}
