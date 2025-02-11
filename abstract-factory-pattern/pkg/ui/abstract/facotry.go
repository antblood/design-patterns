package abstract

// GUIFactory is the abstract factory interface.
type GUIFactory interface {
	CreateButton() Button
	CreateTextField() TextField
}

// Button is the abstract product interface for buttons.
type Button interface {
	Render() string
}

// TextField is the abstract product interface for text fields.
type TextField interface {
	Render() string
}
