package main

import (
    "fmt"
    "techprices/pkg/scrapers"
)

func main() {
    fmt.Println("Entry point")
    done := scrapers.InitScrapers()
    fmt.Println(done)
    scrapers.StartScraping()

}
