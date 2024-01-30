package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Data struct {
	Text string `json:"text"`
}

func main() {
	start := time.Now() // Start recording the execution time.

	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	// Fire up the collector using the colly library
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"), // only permit wikipedia.org English version
	)

	// For loop to hit each URL in a list
	for _, pageURL := range urls {
		fmt.Println("Scraping:", pageURL)

		// Pull the file name from the URLs
		parsedURL, err := url.Parse(pageURL)
		if err != nil {
			panic(err)
		}
		pathSegments := strings.Split(parsedURL.Path, "/")
		title := pathSegments[len(pathSegments)-1]
		fileName := title + ".jl"

		// Create a single file for each corresponding URL in the list
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}

		// On every  element which has the 'mw-parser-output' class, callback
		c.OnHTML(".mw-parser-output p", func(e *colly.HTMLElement) {
			data := Data{
				Text: e.Text,
			}

			jsonData, err := json.Marshal(data)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Record the JSON data to a new line for each line in the scrapped URL
			file.WriteString(string(jsonData) + "\n")
		})

		// Fire up scraping
		c.Visit(pageURL)
		file.Close()
	}

	elapsed := time.Since(start) // Close the execution time measurement
	fmt.Printf("URL web scraping is complete. The execution time was: %s\n", elapsed)
}
