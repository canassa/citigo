package main

import (
    "fmt"
    // "net/http"
    // "io/ioutil"
    "github.com/PuerkitoBio/goquery"
)

// func add(x, y int) int {
//     return x + y
// }

func main() {
    doc, e := goquery.NewDocument("https://www.citibank.com.br/BRGCB/JPS/portal/Index.do")
    if e != nil {
        panic(e.Error())
    }

    loginUrl, _ := doc.Find("#SignonForm").Attr("action")
    syncToken, _ := doc.Find("input[name=SYNC_TOKEN]").Attr("value")
    // var syncToken, _ = doc.Find("input[name=SYNC_TOKEN]").Attr("value")
    // var syncToken, _ = s.Attr("value")
    fmt.Println("syncToken:", syncToken)
    fmt.Println("loginUrl", loginUrl)

    fmt.Println("Teste")

    // doc.Find("input[name=SYNC_TOKEN]").Each(func(i int, s *goquery.Selection) {
        // var syncToken, _ = s.Attr("value")
        // fmt.Printf(syncToken)
        // var band, title string
        // var score float64

        // // For each item found, get the band, title and score, and print it
        // band = s.Find("strong").Text()
        // title = s.Find("em").Text()
        // if score, e = strconv.ParseFloat(s.Find(".score").Text(), 64); e != nil {
        // // Not a valid float, ignore score
        // fmt.Printf("Review %d: %s - %s.\n", i, band, title)
        // } else {
        // // Print all, including score
        // fmt.Printf("Review %d: %s - %s (%2.1f).\n", i, band, title, score)
        // }
    // })

    // resp, err := http.Get("http://www.pudim.com.br/")

    // if err != nil {
    //     // handle error
    // }
    // defer resp.Body.Close()
    // body, err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    // }
    // fmt.Println(string(body))
}
