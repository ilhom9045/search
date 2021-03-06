package search

import (
	"context"
	"io/ioutil"
	"log"
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
	ch := make(chan []Result)
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(ctx)

	for i := 0; i < len(files); i++ {
		wg.Add(1)

		go func(ctx context.Context, filename string, i int, ch chan<- []Result) {

			defer wg.Done()

			res := FindAllMatchTextInFile(phrase, filename)

			if len(res) > 0 {
				ch <- res
			}

		}(ctx, files[i], i, ch)
	}

	go func() {
		defer close(ch)
		wg.Wait()

	}()

	cancel()
	return ch
}
func Any(ctx context.Context, phrase string, files []string) <-chan Result {
	if files == nil {
		return nil
	}
	part := len(files)
	ch := make(chan Result)
	//ctx, cansel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}
	var result Result
	for i := 0; i < part; i++ {
		text, err := ioutil.ReadFile(files[i])
		if err != nil {
			break
		}
		lines := strings.Split(string(text), "\n")

		for index, value := range lines {
			if strings.Contains(value, phrase) {
				res := Result{
					Line:    value,
					LineNum: int64(index + 1),
					Phrase:  phrase,
					ColNum:  int64(strings.Index(value, phrase) + 1),
				}
				result = res
				break
			}
		}
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if (Result{}) != result {
			ch <- result
		}
	}()

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	//cansel()
	return ch
}

func FindAllMatchTextInFile(phrase, fileName string) (res []Result) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("error not opened file err => ", err)
		return res
	}

	file := string(data)

	temp := strings.Split(file, "\n")

	for i, line := range temp {
		if strings.Contains(line, phrase) {

			r := Result{
				Phrase:  phrase,
				Line:    line,
				LineNum: int64(i + 1),
				ColNum:  int64(strings.Index(line, phrase)) + 1,
			}

			res = append(res, r)
		}
	}

	return res
}
