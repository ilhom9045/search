package search

import (
	"context"
	"testing"
)

var str = "aa sda dsd aw sd s aw  ffffff"

var files = []string{
	"data/ex.txt",
	"data/import",
	"data/import",
	"data/import",
	"data/import",
	"data/import",
	"data/import.txt",
	"data/import.txt",
	"data/export.txt",
}

var root = context.Background()

func BenchmarkAny(b *testing.B) {
	var ctx, cancel = context.WithCancel(root)
	for i := 0; i < b.N; i++ {
		_ = <-All(ctx, "2;", files)
		cancel()
	}
}
