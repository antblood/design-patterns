package main

import (
	command_manager "github.com/antblood/design-patterns/command-pattern/pkg/commandmanager"
	"github.com/antblood/design-patterns/command-pattern/pkg/paint"
)

func main() {
	manager := command_manager.NewCommandManager()

	// Create a drawing
	drawing := &paint.Drawing{}

	// Create commands
	addCircle := paint.NewAddShapeCommand(drawing, "Circle")
	addSquare := paint.NewAddShapeCommand(drawing, "Square")

	// Execute commands
	manager.ExecuteCommand(addCircle)
	manager.ExecuteCommand(addSquare)
	drawing.Show()

	// Undo the last command
	manager.Undo()
	drawing.Show()

	// Redo the last undone command
	manager.Redo()
	drawing.Show()

	// Undo twice
	manager.Undo()
	manager.Undo()
	drawing.Show()
}
