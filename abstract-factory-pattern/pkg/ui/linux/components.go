package linux

// LinuxButton is a concrete product that implements the Button interface for Linux.
type LinuxButton struct{}

func (b *LinuxButton) Render() string {
	return "Rendering Linux Button"
}

// LinuxTextField is a concrete product that implements the TextField interface for Linux.
type LinuxTextField struct{}

func (t *LinuxTextField) Render() string {
	return "Rendering Linux TextField"
}
