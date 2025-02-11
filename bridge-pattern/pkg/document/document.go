package document

// Document is the abstraction interface.
type Document interface {
	PrintContent()
	SetFormatter(formatter Formatter)
}
