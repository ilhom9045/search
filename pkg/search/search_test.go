package search

import (
	"context"
	"log"
	"testing"
)

func BenchmarkAny(b *testing.B) {
	ctx := context.Background()
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
	log.Println(<-All(ctx, "2;", files))
}
