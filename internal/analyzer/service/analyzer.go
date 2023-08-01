package service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/wishwaprabodha/go-webscraper/internal/analyzer/domain"
	"net/http"
	"strings"
)

func containsLoginForm(doc *goquery.Document) bool {
	return doc.Find("form input[type='password']").Length() > 0
}

func analyzeLinks(doc *goquery.Document, analyzedLinks chan domain.AnalyzeLinksResult) {
	internalLinkCounter, externalLinkCounter, inaccessibleLinkCounter := 0, 0, 0
	go doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		if strings.HasPrefix(link, "http") {
			externalLinkCounter++
			resp, err := http.Get(link)
			if err != nil || resp.StatusCode != http.StatusOK {
				inaccessibleLinkCounter++
			}
		} else if strings.HasPrefix(link, "/") {
			internalLinkCounter++
		}
	})
	analyzedLinks <- domain.AnalyzeLinksResult{
		InternalLinks:     internalLinkCounter,
		ExternalLinks:     externalLinkCounter,
		InaccessibleLinks: inaccessibleLinkCounter,
	}
	close(analyzedLinks)
}

func GetPageInfo(url string) domain.PageInfo {
	pageInfo := domain.PageInfo{}
	resp, err := http.Get(url)
	if err != nil {
		pageInfo.Error = err
		return pageInfo
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		pageInfo.Error = err
		return pageInfo
	}

	analyzedLinkChan := make(chan domain.AnalyzeLinksResult)
	go analyzeLinks(doc, analyzedLinkChan)
	results := <-analyzedLinkChan
	pageInfo.InternalLinks = results.InternalLinks
	pageInfo.ExternalLinks = results.ExternalLinks
	pageInfo.InaccessibleLinks = results.InaccessibleLinks

	pageInfo.HTMLVersion = doc.Find("html").First().AttrOr("version", "Unknown")
	pageInfo.PageTitle = doc.Find("title").First().Text()

	pageInfo.HeadingCounts = make(map[string]int)
	doc.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
		headingLevel := strings.ToLower(s.Nodes[0].Data)
		pageInfo.HeadingCounts[headingLevel]++
	})

	pageInfo.HasLoginForm = containsLoginForm(doc)

	return pageInfo
}
