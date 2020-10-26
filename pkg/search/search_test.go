package search

import (
	"context"
	"log"
	"testing"
)

func TestAll(t *testing.T) {
	ch := All(context.Background(), "2;", []string{"../../data/export.txt"})
	read, ok := <-ch
	if !ok {
		t.Error(ok)
	}
	log.Println(read)
}
