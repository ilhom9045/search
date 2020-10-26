package search

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"sync"
)

type Result struct {
	//Фраза
	Phrase string
	//Строка
	Line string
	//Номер строки
	LineNum int64
	//Номер позиции
	ColNum int64
}

func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	part := len(files)
	ch := make(chan []Result)
	ctx, cansel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}
	for i := 0; i < part; i++ {
		wg.Add(1)
		go func(ctx context.Context, file string, val int, c chan<- []Result) {
			defer wg.Done()
			if results := ReadFile(file, phrase); len(results) > 0 {
				log.Print(files[val])
				c <- results
			}
		}(ctx, files[i], i, ch)
	}
	go func() {
		defer close(ch)
		wg.Wait()
	}()
	cansel()
	return ch
}

func Any(ctx context.Context, phrase string, files []string) <-chan Result {
	if files == nil {
		return nil
	}
	part := len(files)
	ch := make(chan Result, part)
	//defer close(ch)
	//ctxx, cansel := context.WithCancel(ctx)
	for i := 0; i < part; i++ {
		go func(ctx1 context.Context, fileOpen string, phrase string, c chan<- Result) {
			select {
			case <-ctx1.Done():
				return
			default:
				file, err := os.Open(fileOpen)
				if err != nil {
					return
				}
				defer func() {
					if cerr := file.Close(); cerr != nil {
						log.Print(cerr)
					}
				}()
				reader := bufio.NewReader(file)
				lineNum := 1
				for {
					line, _, err := reader.ReadLine()
					if err != nil || len(line) == 0 {
						break
					}
					if strings.Contains(string(line), phrase) {
						result := Result{}
						colNum := strings.Index(string(line), phrase)
						result.Phrase = phrase
						result.ColNum = int64(colNum)
						result.Line = string(line)
						result.LineNum = int64(lineNum)
						c <- result
						break
					}
					lineNum++
				}
			}
		}(ctx, files[i], phrase, ch)
	}
	//<-ch
	//cansel()
	return ch
}

func ReadFile(f string, phrase string) (results []Result) {
	file, err := os.Open(f)
	if err != nil {
		return results
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()
	reader := bufio.NewReader(file)
	lineNum := 1
	for {
		line, _, err := reader.ReadLine()
		if err != nil || len(line) == 0 {
			break
		}
		if strings.Contains(string(line), phrase) {
			results = append(results, Result{
				Phrase:  phrase,
				Line:    string(line),
				LineNum: int64(lineNum),
				ColNum:  int64(strings.Index(string(line), phrase) + 1),
			})
		}
		lineNum++
	}
	return results
}
