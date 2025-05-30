package lint

import (
	"fmt"
	"os"

	"github.com/verily-src/fsh-lint/internal/fsh"
	"github.com/verily-src/fsh-lint/internal/fsh/types"
)

// FileContext represents the context of a file being linted.
type FileContext struct {
	// Path is the path to the file being linted.
	Path string

	// Data is the file data in bytes
	Data []byte

	// ParsedFSH is the data in parsed form
	ParsedFSH *types.FSHDocument
}

// NewFileContext creates a new FileContext with the given path.
func NewFileContext(path string) (*FileContext, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	f := string(data)
	parsedFSH, err := fsh.Parse(f)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse FSH file %v: %w", path, err)
	}

	return &FileContext{
		Path:      path,
		Data:      data,
		ParsedFSH: parsedFSH,
	}, nil
}
