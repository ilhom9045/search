package search

import (
	"context"
	"log"
	"testing"
)

func TestAll(t *testing.T) {
	log.Println(<-All(context.Background(), "2;", []string{"export.txt"}))
}
