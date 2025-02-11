package filesystem

import "fmt"

// Directory struct
type Directory struct {
	Name     string
	Children []Component
}

// Add method for Directory to add a component
func (d *Directory) Add(component Component) {
	d.Children = append(d.Children, component)
}

// Show method for Directory
func (d *Directory) Show(indentation string) {
	fmt.Println(indentation + d.Name)
	for _, component := range d.Children {
		component.Show(indentation + "  ")
	}
}
