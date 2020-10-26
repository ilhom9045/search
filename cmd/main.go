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
		"data/export.txt",
		"data/export.txt",
		"data/export.txt",
		"data/export.txt",
		"export.txt",
		"export.txt",
		"export.txt",
	}
	log.Println(<-search.All(ctx, "2;", files))
}

//func Any(ctx context.Context, phrase string, files []string) <-chan Result {
//	if files == nil {
//		return nil
//	}
//	part := len(files)
//	ch := make(chan Result, part)
//	//defer close(ch)
//	//ctxx, cansel := context.WithCancel(ctx)
//	for i := 0; i < part; i++ {
//		go func(ctx1 context.Context, fileOpen string, phrase string, c chan<- Result) {
//			select {
//			case <-ctx1.Done():
//				return
//			default:
//				file, err := os.Open(fileOpen)
//				if err != nil {
//					return
//				}
//				defer func() {
//					if cerr := file.Close(); cerr != nil {
//						log.Print(cerr)
//					}
//				}()
//				reader := bufio.NewReader(file)
//				lineNum := 1
//				for {
//					line, _, err := reader.ReadLine()
//					if err != nil || len(line) == 0 {
//						break
//					}
//					if strings.Contains(string(line), phrase) {
//						result := Result{}
//						colNum := strings.Index(string(line), phrase)
//						result.Phrase = phrase
//						result.ColNum = int64(colNum)
//						result.Line = string(line)
//						result.LineNum = int64(lineNum)
//						c <- result
//						break
//					}
//					lineNum++
//				}
//			}
//		}(ctx, files[i], phrase, ch)
//	}
//	//<-ch
//	//cansel()
//	return ch
//}
