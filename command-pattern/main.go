package main

import (
	"github.com/antblood/design-patterns/command-pattern/pkg/paint"
	"github.com/antblood/design-patterns/command-pattern/pkg/commandmanager"
)


func main() {
	drawing := &paint.Drawing{}
	manager := commandmanager.NewCommandManager()

	// Create and execute commands
	addCircle := NewAddShapeCommand(drawing, "Circle")
	addSquare := NewAddShapeCommand(drawing, "Square")

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
