package tree_sitter_clojure_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-clojure"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_clojure.Language())
	if language == nil {
		t.Errorf("Error loading Clojure grammar")
	}
}
