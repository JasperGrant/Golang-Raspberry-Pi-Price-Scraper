/*
scraper.go: Simple webscraper in go.
Will make a search for the top raspberry pi listings on amazon.
Will print an extra line for products on sale and notify
of price difference.
By Jasper Grant Feb 14th 2023
*/

//Main package
package main

import (
	"fmt"                      //Standard input/output
	"github.com/gocolly/colly" //Web scraping library
	"strconv"                  //String conversion library
	"strings"                  //General use string library
)

// Main function
func main() {
	//Initialize colly collector with amazon URL
	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.ca"),
	)

	//On colly request print URL to stdout
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("URL: ", r.URL)
	})

	//On HTML of webpage output names and prices of top results for Raspberry Pi on amazon
	c.OnHTML("div.s-main-slot.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		//For each element in search
		h.ForEach("div.sg-col-4-of-24.sg-col-4-of-12.s-result-item", func(_ int, h *colly.HTMLElement) {
			//Print line for ease of reading different entries
			fmt.Println("---------------------------------------------------------------------------------------------")
			//Get name from HTML
			name := h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
			//Declare max string length before covering by ...
			maxstrlen := 64
			//If string is too long cut it off at maxstrlen and add ...
			if len(name) > maxstrlen {
				name = name[:maxstrlen-1] + "..."
			}
			//Print name of product
			fmt.Println("Product Name: ", name)
			//Get price from HTML
			price := h.ChildText("span.a-price-symbol, span.a-price-whole, span.a-price-fraction")
			//Print price of product
			fmt.Println("Price: ", price)
			//Get non-sale price from HTML
			nonsaleprice := h.ChildText("span.a-price.a-text-price > span.a-offscreen")
			//If there is a non sale price
			if len(nonsaleprice) != 0 {
				//Print non sale price
				fmt.Println("On sale from: ", nonsaleprice)
				//Declare floats to find difference between non-sale price and sale price
				nonsalepricefloat, _ := strconv.ParseFloat(strings.Trim(nonsaleprice, "$"), 32)
				pricefloat, _ := strconv.ParseFloat(strings.Trim(price, "$"), 32)
				//Print amount saved by sale price
				fmt.Printf("You save: $%.2f\n", nonsalepricefloat-pricefloat)
			}
		})
	})

	//Load webpage of raspberry pi search
	c.Visit("https://www.amazon.ca/s?k=Raspberry+Pi")
}
