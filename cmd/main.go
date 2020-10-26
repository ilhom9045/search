package main

import (
	"context"
	"github.com/ilhom9045/search/pkg/search"
	"log"
)

type Rest struct {
	Name     string `json:"name"`
	FullName string `json:"fullname"`
}

func main() {
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
	log.Println(<-search.All(ctx, "2;", files))
}
