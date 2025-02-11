package paint

// AddShapeCommand adds a shape to the drawing
type AddShapeCommand struct {
	drawing *Drawing
	shape   string
}

func NewAddShapeCommand(drawing *Drawing, shape string) *AddShapeCommand {
	return &AddShapeCommand{
		drawing: drawing,
		shape:   shape,
	}
}

func (c *AddShapeCommand) Execute() {
	c.drawing.AddShape(c.shape)
}

func (c *AddShapeCommand) Undo() {
	c.drawing.RemoveShape()
}

func (c *AddShapeCommand) Redo() {
	c.Execute()
}
