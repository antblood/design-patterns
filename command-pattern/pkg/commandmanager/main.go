package commandmanager

import "fmt"

// Command interface with Execute, Undo, and Redo methods
type Command interface {
	Execute()
	Undo()
	Redo()
}

// CommandManager manages the command history for undo/redo
type CommandManager struct {
	history      []Command
	redoStack    []Command
	currentIndex int
}

func NewCommandManager() *CommandManager {
	return &CommandManager{history: []Command{}, redoStack: []Command{}}
}

func (cm *CommandManager) ExecuteCommand(command Command) {
	command.Execute()
	cm.history = append(cm.history[:cm.currentIndex], command)
	cm.currentIndex++
	cm.redoStack = []Command{} // Clear the redo stack when a new command is executed
}

func (cm *CommandManager) Undo() {
	if cm.currentIndex == 0 {
		fmt.Println("Nothing to undo")
		return
	}
	cm.currentIndex--
	command := cm.history[cm.currentIndex]
	command.Undo()
	cm.redoStack = append(cm.redoStack, command)
}

func (cm *CommandManager) Redo() {
	if len(cm.redoStack) == 0 {
		fmt.Println("Nothing to redo")
		return
	}
	command := cm.redoStack[len(cm.redoStack)-1]
	cm.redoStack = cm.redoStack[:len(cm.redoStack)-1]
	command.Redo()
	cm.currentIndex++
}
