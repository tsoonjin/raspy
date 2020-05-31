package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
    "github.com/PuerkitoBio/goquery"
    "encoding/json"
)

type Page struct {
    Imgs []string `json: "imgs"`
    Links []string `json: "links"`
    Videos []string `json: "vids"`
}

// This will get called for each HTML element found
func processImage(index int, element *goquery.Selection) {
    imgSrc, imgExists := element.Attr("src")
    linkSrc, linkExists := element.Attr("href")
    videoSrc, videoExists := element.Attr("data-video-id")
    if imgExists {
        fmt.Printf("Img: %s\n\n", imgSrc)
    }
    if linkExists {
        fmt.Printf("Link: %s\n\n", linkSrc)
    }
    if videoExists {
        fmt.Printf("Video: %s\n\n", videoSrc)
    }
}

func extractPage(doc *goquery.Document) Page {
    imgs  := []string{}
    links := []string{}
    vids  := []string{}

	if doc != nil {
        doc.Find("a,img,video").Each(func(i int, element *goquery.Selection) {
            elementType := element.Nodes[0].Data
            src, srcExists := element.Attr("src")
            link, linkExists := element.Attr("href")
            // imgSrc, imgExists := element.Attr("src")
            // linkSrc, linkExists := element.Attr("href")
            // videoSrc, videoExists := element.Attr("data-video-id")
            switch elementType {
                case "img":
                    if srcExists {
                        imgs = append(imgs, src)
                    }
                case "a":
                    if linkExists {
                        links = append(links, link)
                    }
                case "video":
                    if srcExists {
                        vids = append(vids, src)
                    }
            }
            // if imgExists {
            //     imgs = append(imgs, imgSrc)
            //     fmt.Printf("Img: %s\n\n", imgSrc)
            // }
            // if linkExists {
            //     links = append(links, linkSrc)
            //     fmt.Printf("Link: %s\n\n", linkSrc)
            // }
            // if videoExists {
            //     vids = append(vids, videoSrc)
            //     fmt.Printf("Video: %s\n\n", videoSrc)
            // }

        })
    }
    return Page{imgs, links, vids}
}

func main() {
    // Arguments Parsing

    url := flag.String("url", "", "Site URL to crawl and extract")
    flag.Parse()
    fmt.Printf("Analyzing URL: %s\n", *url)

    // Process url
    response, err := http.Get(*url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Create a goquery document from the HTTP response
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }

    page := extractPage(document)
    pageJSON, err := json.Marshal(page)
    fmt.Println(string(pageJSON))
}
