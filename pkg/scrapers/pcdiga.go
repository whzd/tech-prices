package scrapers

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type PcdigaScraper struct {
    Domain string
}

func NewPcdigaScraper() PcdigaScraper{
    scrp := PcdigaScraper{
        Domain: "www.pcdiga.com",
    }
    return scrp
}

func (scrpr PcdigaScraper) Scrape(){
    useragent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"
    c := colly.NewCollector(colly.UserAgent(useragent), colly.Debugger(&debug.LogDebugger{}))
    
    // On every a element which has href attribute call callback
	c.OnHTML("span.price__amount", func(e *colly.HTMLElement) {
        fmt.Println("-----------------")
        fmt.Println(e.Text)
	})
    c.OnRequest(func(r *colly.Request) {
       fmt.Println("Visiting", r.URL)
    })
    c.OnResponse(func(r *colly.Response) {
       fmt.Println(r.StatusCode)
    })
    // Set error handler
	c.OnError(func(r *colly.Response, err error) {
        fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.Headers, "\nError:", err)
	})
    c.Visit("https://www.pcdiga.com")

    fmt.Println("Scraping: " + scrpr.Domain)
}
