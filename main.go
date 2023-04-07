package main

import (
    "fmt"
    "encoding/csv"
    "github.com/gocolly/colly"
    "log"
    "os"
)

// initializing a data structure to keep the scraped data
type PokemonProduct struct {
    title, price string
    // title, price, area, price_per_m2, bedroom, toilet string
}

func main() {
    // initializing the slice of structs to store the data to scrape
    var pokemonProducts []PokemonProduct
    var url = "https://batdongsan.com.vn/ban-nha-rieng-ha-noi"

    // creating a new Colly instance
    c := colly.NewCollector()

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.OnHTML(".re__body", func(e *colly.HTMLElement) {
        fmt.Println("DT")
    })

    // scraping logic
    c.OnHTML(".js__product-link-for-product-id", func(e *colly.HTMLElement) {
        fmt.Println(e)
        pokemonProduct := PokemonProduct{}

        pokemonProduct.title = e.ChildText(".pr-title js__card-title")
        pokemonProduct.price = e.ChildText(".re__card-config-price")
        // pokemonProduct.price = e.ChildText(".re__card-config-price")
        // pokemonProduct.url = e.ChildAttr("a", "href")
        // pokemonProduct.image = e.ChildAttr("img", "src")
        // pokemonProduct.price = e.ChildText(".price")

        pokemonProducts = append(pokemonProducts, pokemonProduct)
    })

    // visiting the target page
    c.Visit(url)

    fmt.Println(pokemonProducts)

    // opening the CSV file
    file, err := os.Create("products.csv")
    if err != nil {
        log.Fatalln("Failed to create output CSV file", err)
    }
    defer file.Close()

    // initializing a file writer
    writer := csv.NewWriter(file)

    // writing the CSV headers
    headers := []string{
        "title",
        "price",
    }
    writer.Write(headers)

    // writing each Pokemon product as a CSV row
    for _, pokemonProduct := range pokemonProducts {
        // converting a PokemonProduct to an array of strings
        record := []string{
            pokemonProduct.title,
            pokemonProduct.price,
        }

        // adding a CSV record to the output file
        writer.Write(record)
    }
    defer writer.Flush()
}
