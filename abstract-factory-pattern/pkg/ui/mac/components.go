package mac

// MacButton is a concrete product that implements the Button interface for macOS.
type MacButton struct{}

func (b *MacButton) Render() string {
	return "Rendering Mac Button"
}

// MacTextField is a concrete product that implements the TextField interface for macOS.
type MacTextField struct{}

func (t *MacTextField) Render() string {
	return "Rendering Mac TextField"
}
