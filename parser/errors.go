package parser

import "fmt"

// Diagnostic is a diagnostic message
type Diagnostic struct {
	Position
	Message string
	Severity
}

// Severity is the severity of the diagnostic message
type Severity int // 1: Error, 2: Warning, 3: Information, 4: Hint
const (
	Error       Severity = 1
	Warning     Severity = 2
	Information Severity = 3
	Hint        Severity = 4
)

func (d *Diagnostic) String() string {
	return fmt.Sprintf("%d Line: %d, Col: %d, %s", d.Severity, d.line, d.col, d.Message)
}
