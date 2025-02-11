package filesystem

import "fmt"

// File struct
type File struct {
	Name string
}

// Show method for File
func (f *File) Show(indentation string) {
	fmt.Println(indentation + f.Name)
}
