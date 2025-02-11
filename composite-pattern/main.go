package main

import (
	"fmt"

	"github.com/antblood/design-patterns/composite-pattern/pkg/filesystem"
)

func main() {
	// Create files
	file1 := &filesystem.File{Name: "File1.txt"}
	file2 := &filesystem.File{Name: "File2.txt"}
	file3 := &filesystem.File{Name: "File3.txt"}

	// Create directories
	dir1 := &filesystem.Directory{Name: "Dir1"}
	dir2 := &filesystem.Directory{Name: "Dir2"}
	dir3 := &filesystem.Directory{Name: "Dir3"}

	// Construct the tree structure
	dir1.Add(file1)
	dir1.Add(file2)
	dir2.Add(file3)
	dir2.Add(dir1)
	dir3.Add(dir2)

	// Display the structure
	fmt.Println("File System Structure:")
	dir3.Show("")
}
