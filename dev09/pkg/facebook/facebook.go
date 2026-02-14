package facebook

import (
	"fmt"
	"io"
)

type Writer struct {
	fileName string
}

//compile-time assertion, проверка на этапе компиляции, что тип действительно реализует интерфейс
var _ io.Writer = (*Writer)(nil)

func NewWriter(fileName string) (*Writer, error) {
	if fileName == "" {
		return nil, fmt.Errorf("facebook: empty filename")
	}

	return &Writer{fileName: fileName}, nil
}

func (w *Writer) Write(p []byte) (int, error) {
	fmt.Printf("Facebook file: %s, data: %s\n", w.fileName, p)
    
	return len(p), nil
}