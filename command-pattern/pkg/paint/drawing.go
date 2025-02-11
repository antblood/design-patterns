package paint

import "fmt"

// Drawing is the Receiver that maintains the state
type Drawing struct {
	shapes []string
}

func (d *Drawing) AddShape(shape string) {
	d.shapes = append(d.shapes, shape)
	fmt.Println("Added shape:", shape)
}

func (d *Drawing) RemoveShape() string {
	if len(d.shapes) == 0 {
		return ""
	}
	shape := d.shapes[len(d.shapes)-1]
	d.shapes = d.shapes[:len(d.shapes)-1]
	fmt.Println("Removed shape:", shape)
	return shape
}

func (d *Drawing) Show() {
	fmt.Println("Current Drawing:", d.shapes)
}
