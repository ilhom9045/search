package search

import (
	"context"
	"log"
	"testing"
)

func BenchmarkAny(b *testing.B) {
	root, cansel := context.WithCancel(context.Background())
	files := []string{
		"import.txt",
		"import.txt",
		"import.txt",
		"import.txt",
		"import.txt",
		"import.txt",
		"export.txt",
		"export.txt",
		"export.txt",
		"export.txt",
		"export.txt",
		"export.txt",
		"export.txt",
	}
	log.Println(<-All(root, "2;", files))
	cansel()
}
