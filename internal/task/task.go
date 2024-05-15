package task

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"explorer_webarchiv/internal/utils"

	"github.com/PuerkitoBio/goquery"
)

const BASE_URL = "http://web.archive.org/web/"

type Task struct {
	sem chan struct{}
}

func New(maxWorkers int) *Task {
	return &Task{
		sem: make(chan struct{}, maxWorkers),
	}
}

func (t *Task) Run(ctx context.Context, wg *sync.WaitGroup, numWorker uint, rootDir string, inputURL string, fileName string) {
	t.sem <- struct{}{}
	defer func() { <-t.sem }()

	defer wg.Done()

	select {
	case <-ctx.Done():
		log.Printf("worker %d was done", numWorker)
	default:
		log.Printf("worker: %d, downloading: %s", numWorker+1, inputURL)

		fileName := utils.UrlToFileName(inputURL) + ".txt"
		pathFileName := filepath.Join(rootDir, fileName)
		if utils.PathExists(pathFileName) != false {
			log.Printf("skipping: %s", inputURL)
			return
		}

		doc, err := goquery.NewDocument(BASE_URL + inputURL)
		if err != nil {
			log.Printf("error encountered: %s while retrieving content from URL: %s", err, inputURL)
			return
		}

		result := strings.Builder{}

		doc.Find("body").Each(func(i int, s *goquery.Selection) {
			s.Contents().Each(func(j int, c *goquery.Selection) {
				if c.Is("script, style") {
					return
				}

				if content := strings.TrimSpace(c.Text()); content != "" {
					re := regexp.MustCompile(`(\s{2,}|\n{3,})`)
					cleanedText := re.ReplaceAllString(content, "\n\n")

					result.WriteString(cleanedText)
				}
			})
		})

		file, err := os.Create(pathFileName)
		if err != nil {
			log.Printf("error creating file: %s", err)
			return
		}

		file.WriteString(result.String())
		file.Close()

		log.Printf("done: %s", pathFileName)
	}
}
