package stdout

import (
	"fmt"
	"io"
)

type Writer struct {
	message string
}

var _ io.Writer = (*Writer)(nil)

func NewWriter(message string) (*Writer, error) {
	if message == "" {
		return nil, fmt.Errorf("stdout: empty string")
	}

	return &Writer{message: message}, nil
}

func (w *Writer) Write(p []byte) (int, error) {
	fmt.Printf("Stdout message=%s data=%s\n", w.message, p)
	return len(p), nil
}