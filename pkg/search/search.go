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
		return nil
	}
	part := len(files)
	ch := make(chan []Result, 1)
	var result []Result
	for i := 0; i < part; i++ {
		select {
		case <-ctx.Done():
			continue
		default:
			file, err := os.Open(files[i])
			if err != nil {
				continue
			}
			defer func() {
				if cerr := file.Close(); cerr != nil {
					log.Print(cerr)
				}
			}()
			reader := bufio.NewReader(file)
			lines := ""
			for {
				line, _, err := reader.ReadLine()
				if err != nil || len(line) == 0 {
					break
				}
				lines += string(line)
			}
			index := int64(strings.Index(lines, phrase))

			if index != -1 {
				result = append(result, Result{ColNum: index})
			}
		}
	}

	log.Print(result)
	ch <- result
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
