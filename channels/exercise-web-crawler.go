package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type CrawlUrlMap struct {
	// 已经抓取页面url
	urlMap map[string]bool
	// 保证并发安全
	mutex sync.Mutex
}

func (m *CrawlUrlMap) exist(url string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, ok := m.urlMap[url]
	m.urlMap[url] = true
	return ok
}

// 防止重复抓取页面
var crawlUrlMap = &CrawlUrlMap{
	urlMap: make(map[string]bool),
}

// 保证多个协程完成执行
var wg = sync.WaitGroup{}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		if !crawlUrlMap.exist(u) {
			wg.Add(1)
			go func(url string, depth int, fetcher Fetcher) {
				Crawl(url, depth, fetcher)
				wg.Done()
			}(u, depth-1, fetcher)
		}
	}
	return
}

func main() {
	if !crawlUrlMap.exist("https://golang.org/") {
		Crawl("https://golang.org/", 4, fetcher)
	}
	wg.Wait()
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
