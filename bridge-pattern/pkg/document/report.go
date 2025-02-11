package document

import "fmt"

// Report is a refined abstraction.
type Report struct {
	Title     string
	Content   string
	formatter Formatter
}

func (r *Report) SetFormatter(formatter Formatter) {
	r.formatter = formatter
}

func (r *Report) PrintContent() {
	if r.formatter == nil {
		fmt.Println("No formatter set")
		return
	}
	fmt.Println(r.formatter.Format("Report: " + r.Title + "\n" + r.Content))
}
