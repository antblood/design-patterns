package image

type ExternalImage struct {
	filename string
	image    *image
}

// NewExternalImage creates a new ExternalImage.
func NewExternalImage(filename string) *ExternalImage {
	return &ExternalImage{
		filename: filename,
	}
}

// Display displays the ExternalImage, loading it if necessary.
func (p *ExternalImage) Display() {
	if p.image == nil {
		p.image = Newimage(p.filename)
	}
	p.image.Display()
}
