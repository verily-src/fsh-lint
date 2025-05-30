package fsh

import (
	"fmt"

	"github.com/verily-src/fsh-lint/internal/fsh/internal/grammar"
	"github.com/verily-src/fsh-lint/internal/fsh/internal/parser"
	"github.com/verily-src/fsh-lint/internal/fsh/types"

	"github.com/antlr4-go/antlr/v4"
)

// Parse parses a FSH doc to a FSHDocument object.
func Parse(fshData string) (*types.FSHDocument, error) {
	stream := antlr.NewInputStream(fshData)

	// Lex the input stream
	lexer := grammar.NewFSHLexer(stream)
	errorListener := &parser.FSHErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Parse the tokens
	p := grammar.NewFSHParser(tokens)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
	tree := p.Doc()
	v := &parser.FSHVisitor{}

	doc, err := v.VisitDoc(tree)
	if err != nil {
		return nil, fmt.Errorf("fsh parse tree: %w", err)
	}

	if err := errorListener.Error(); err != nil {
		return nil, err
	}

	return doc, nil
}
