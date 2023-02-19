package scrapers

import "fmt"

type Scraper interface{
    Scrape()
}
var scrapers []Scraper

func InitScrapers() string {
	pcdiga := NewPcdigaScraper()
    scrapers = append(scrapers, pcdiga)
    globaldata := NewGlobaldataScraper()
    scrapers = append(scrapers, globaldata)
    fmt.Println(scrapers)
	return "Initialiazed!"
}

func StartScraping(){
    for _, scraper := range scrapers{
        scraper.Scrape()
    }
}

