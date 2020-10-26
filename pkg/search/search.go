package search

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
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
	if files == nil {
		log.Println("Nil")
		return nil
	}
	part := len(files)
	ch := make(chan []Result, 1)
	//defer close(ch)
	var result []Result
	//ctx, cancel := context.WithCancel(ctx)
	//wg := sync.WaitGroup{}
	for i := 0; i < part; i++ {
		//wg.Add(1)
		go func(ctx context.Context, val int) {
			//defer wg.Done()
			select {
			case <-ctx.Done():
				log.Println("Done")
				return
			default:
				file, err := os.Open(files[val])
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
						res := Result{}
						colNum := strings.Index(string(line), phrase)
						res.ColNum = int64(colNum)
						log.Println(files[val])
						result = append(result, res)
						break
					}
					lineNum++
				}
				ch <- result
			}
		}(ctx, i)
	}
	//wg.Wait()
	//close(ch)
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
				log.Println("done")
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

func ReadFile(f string, phrase string) (Result, error) {
	result := Result{}
	file, err := os.Open(f)
	if err != nil {
		return result, err
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
			colNum := strings.Index(string(line), phrase)
			result.Phrase = phrase
			result.ColNum = int64(colNum)
			result.Line = string(line)
			result.LineNum = int64(lineNum)
			return result, nil
		}
		lineNum++
	}
	return result, err
}
