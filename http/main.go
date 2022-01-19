package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

func main() {
	// http.Handle("/", middleware(http.HandlerFunc(helloWorldHandler)))
	// http.ListenAndServe(":8080", nil)
	httpClient()
}

func helloWorldHandler(w http.ResponseWriter, request *http.Request) {

	reqId := request.Context().Value("reqId")
	if reqId != nil {
		w.Write([]byte(fmt.Sprintf("hello world with request id! %s", reqId)))
		return
	}

	w.Write([]byte(fmt.Sprintf("hello world!")))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		ctx := context.WithValue(r.Context(), "reqId", id)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}

type Data struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Author string `json:"author"`
}

type Articles struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []Data `json:"data"`
}

func (a Articles) GetByAuthor(author string) []Data {
	articles := []Data{}
	for _, data := range a.Data {
		if data.Author == author {
			articles = append(articles, data)
		}
	}
	return articles
}

func httpClient() {

	allArticleChannels := []<-chan Articles{}

	for i := 0; i <= 1000; i++ {
		allArticleChannels = append(allArticleChannels, getArticles(i))
	}

	merged := articleChReaderMerger(allArticleChannels)

	for final := range articleFilter("epaga", merged) {
		fmt.Println(final)
	}
}

func getArticles(pageNo int) <-chan Articles {

	ch := make(chan Articles)

	go func() {

		defer close(ch)

		str := fmt.Sprintf("https://jsonmock.hackerrank.com/api/articles?page=%d", pageNo)

		res, err := http.Get(str)

		if err != nil {
			log.Fatalf(err.Error())
		}

		defer res.Body.Close()

		var articles Articles

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalf(err.Error())
		}

		err = json.Unmarshal(body, &articles)

		if err != nil {
			log.Fatalf(err.Error())
		}
		ch <- articles
	}()

	return ch
}

func articleChReaderMerger(articlesCh []<-chan Articles) <-chan Articles {

	var wg sync.WaitGroup

	finalCh := make(chan Articles)

	readrFn := func(articlesCh <-chan Articles) {
		defer wg.Done()
		for article := range articlesCh {
			finalCh <- article
		}
	}

	wg.Add(len(articlesCh))

	for _, ch := range articlesCh {
		go readrFn(ch)
	}

	go func() {
		wg.Wait()
		close(finalCh)
	}()

	return finalCh
}

func articleFilter(authorName string, articlesChn <-chan Articles) <-chan []Data {

	filteredCh := make(chan []Data)

	go func() {
		defer close(filteredCh)
		for article := range articlesChn {
			log.Println(time.Now())
			filtered := article.GetByAuthor(authorName)
			if len(filtered) > 0 {
				filteredCh <- filtered
			}
		}
	}()

	return filteredCh
}
