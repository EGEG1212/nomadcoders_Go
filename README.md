# nomadcoders_Go

**REFERENCE goquery** <br>
https://github.com/PuerkitoBio/goquery

```golang
package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
  // Request the HTML page.
  res, err := http.Get("http://metalsucks.net")
  if err != nil { //🧐에러체크
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {  //🧐goquery document 만들때도 에러체크 ✔에러계속체크체크해야하기때문에, ☢checkErr함수만드는것이 편할것이다~~
    log.Fatal(err)
  }

  // Find the review items
  doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

func main() {
  ExampleScrape()
}
```
