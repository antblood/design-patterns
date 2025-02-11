package main

import "github.com/antblood/design-patterns/proxy-pattern/pkg/image"

func main() {
	image1 := image.NewExternalImage("image1.png")
	image2 := image.NewExternalImage("image2.png")

	// The images will be loaded and displayed only when Display is called.
	image1.Display() // Loading image1.png, Displaying image1.png
	image2.Display() // Loading image2.png, Displaying image2.png

	// The images are already loaded, so Display will not load them again.
	image1.Display() // Displaying image1.png
	image2.Display() // Displaying image2.png
}
