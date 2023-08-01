package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/wishwaprabodha/go-webscraper/internal/analyzer/domain"
	"github.com/wishwaprabodha/go-webscraper/internal/analyzer/service"
	"net/http"
	"sync"
	"time"
)

func AnalyzeWebPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	url := params.Get("url")
	if url == "" {
		http.Error(w, "empty URL", http.StatusBadRequest)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	var pageInfo domain.PageInfo
	go func() {
		defer wg.Done()
		pageInfo = service.GetPageInfo(url)
	}()

	timeout := make(chan bool, 1)
	go func() {
		wg.Wait()
		timeout <- false
	}()

	select {
	case <-timeout:
		if pageInfo.Error != nil {
			http.Error(w, pageInfo.Error.Error(), http.StatusBadRequest)
		} else {
			pageInfoResponse := domain.PageInfo{
				HTMLVersion:       pageInfo.HTMLVersion,
				PageTitle:         pageInfo.PageTitle,
				HeadingCounts:     pageInfo.HeadingCounts,
				InternalLinks:     pageInfo.InternalLinks,
				ExternalLinks:     pageInfo.ExternalLinks,
				InaccessibleLinks: pageInfo.InaccessibleLinks,
				HasLoginForm:      pageInfo.HasLoginForm,
				Error:             pageInfo.Error,
			}
			fmt.Println("Data Fetched")
			errJson := json.NewEncoder(w).Encode(pageInfoResponse)
			if errJson != nil {
				http.Error(w, "Error converting to JSON", http.StatusBadRequest)
			}
		}
	case <-time.After(30 * time.Second):
		http.Error(w, "Timeout: The request took too long to process", http.StatusRequestTimeout)
	}
}
