package document

// Formatter is the formatter interface.
type Formatter interface {
	Format(content string) string
}

// PlainTextFormatter is a concrete implementor.
type PlainTextFormatter struct{}

func (f *PlainTextFormatter) Format(text string) string {
	return text
}

// HTMLFormatter is another concrete implementor.
type HTMLFormatter struct{}

func (f *HTMLFormatter) Format(text string) string {
	return "<p>" + text + "</p>"
}
