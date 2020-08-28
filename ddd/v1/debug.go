package ddd

import (
	"fmt"
	"runtime"
)

// Pos is a debug information, to capture a source code origin.
type Pos struct {
	Name string
	File string
	Line int
}

// String returns a text which printed into the console, is usually clickable by your favorite IDE.
func (p Pos) String() string {
	return fmt.Sprintf("%s: %s:%d", p.Name, p.File, p.Line)
}

// capturePos reads the current calling file and line information. Using skip, you can walk up the stack frame.
func capturePos(name string, skip int) Pos {
	_, fn, line, _ := runtime.Caller(1 + skip)
	return Pos{
		Name: name,
		File: fn,
		Line: line,
	}
}
