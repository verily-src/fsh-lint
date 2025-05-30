package parser

import (
	"github.com/verily-src/fsh-lint/internal/fsh/types"
)

// keyValue is a simple struct to hold a key-value pair.
type keyValue struct {
	key   string
	value *types.ParsedElement[string]
}
