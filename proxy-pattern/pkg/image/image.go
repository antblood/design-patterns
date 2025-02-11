package image

import "fmt"

// Image defines the interface for displaying an image.
type Image interface {
	Display()
}

// image is the real object that loads and displays an image.
type image struct {
	filename string
}

// Newimage creates a new image and loads the file.
func Newimage(filename string) *image {
	image := &image{filename: filename}
	image.loadFromDisk()
	return image
}

// Display displays the image.
func (r *image) Display() {
	fmt.Println("Displaying", r.filename)
}

// ============
// Following methods aren't accessible from the outside of the package.
// ============

// loadFromDisk simulates loading an image from disk.
func (r *image) loadFromDisk() {
	fmt.Println("Loading", r.filename)
}

// deleteFromDisk simulates deleting an image from disk.
func (r *image) deleteFromDisk() {
	// This isn't accessible from the outside of the package.
	fmt.Println("Deleting", r.filename)
}
