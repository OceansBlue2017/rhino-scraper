// finally got Git to work for me on 11-01-2020
// editing in GitHub.  testing 123
package main

import (
	"fmt"
	"log"
	"strconv"
	"github.com/gocolly/colly"
	"encoding/json"
	//"os"
	"io/ioutil" // to be able to write to a file 
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)


	collector.OnHTML(".factsList li", func(element *colly.HTMLElement){
		FactID, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Could not get id")
		}

		factDesc := element.Text

		fact := Fact{
			ID: FactID,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
		
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")

	/*
	enc := json.NewEncoder(os.Stdout)   
	enc.SetIndent("", " ")
	enc.Encode(allFacts) // Replacing this section with "writeJSON(allFacts)"  To write to the JSON file instead of the Terminal.   
	*/
	writeJSON(allFacts)
}

func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create JSON file")
		return 
	}

	_= ioutil.WriteFile("rhinofacts.json", file, 0644) // 0644 is a type of file format. "http://permissions-calculator.org/decode/0644/"
}
