package scrapers

import "fmt"

type GlobaldataScraper struct{
    Domain string
}

func NewGlobaldataScraper() GlobaldataScraper{
    scrp := GlobaldataScraper{
        Domain: "https://www.globaldata.pt/",
    }
    return scrp
}

func (scrpr GlobaldataScraper) Scrape(){
    fmt.Println("Scraping: " + scrpr.Domain)
}
