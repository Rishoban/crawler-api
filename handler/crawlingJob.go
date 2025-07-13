package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron/v3"
	"gorm.io/datatypes"
)

func (v *CrawlerService) OnTimer() {

	c := cron.New()
	c.AddFunc("@every 10s", func() { v.makeCrawlingJob() })
	c.Start()
}

func (v *CrawlerService) makeCrawlingJob() {
	var records = []GeneralObject{}
	if err := v.DbConnection.Table("crawler_url").Find(&records).Error; err != nil {
		// Log the error since ctx is not available here
		fmt.Println("Failed to fetch records:", err)
		return
	}

	for _, record := range records {
		var obj = make(map[string]interface{})
		json.Unmarshal(record.ObjectInfo, &obj)

		if record.ObjectStatus == "Active" {
			currentStatus, ok := obj["status"].(string)
			if ok && currentStatus == "Queued" {
				obj["checkbox"] = false
				obj["status"] = "Running"
				objBytes, _ := json.Marshal(obj)
				UpdateRecord(v.DbConnection, "crawler_url", record.Id, datatypes.JSON(objBytes), record.CreatedBy, "Active")

				// --- Begin goquery analysis (same as CreateNewRecord) ---
				url, ok := obj["url"].(string)
				if !ok || url == "" {
					obj["checkbox"] = false
					obj["status"] = "Error"
					objBytes, _ = json.Marshal(obj)
					UpdateRecord(v.DbConnection, "crawler_url", record.Id, datatypes.JSON(objBytes), record.CreatedBy, "Active")
					continue
				}

				// Check for stop signal before crawling
				var checkedCrawlerStatus = GetStatusOfCrawler(v.DbConnection, record.Id)
				if !checkedCrawlerStatus {
					continue
				}

				resp, err := http.Get(url)
				if err != nil {
					obj["checkbox"] = false
					obj["status"] = "Error"
					objBytes, _ = json.Marshal(obj)
					UpdateRecord(v.DbConnection, "crawler_url", record.Id, datatypes.JSON(objBytes), record.CreatedBy, "Active")
					continue
				}
				defer resp.Body.Close()
				if resp.StatusCode < 200 || resp.StatusCode >= 400 {
					obj["checkbox"] = false
					obj["status"] = "Error"
					objBytes, _ = json.Marshal(obj)
					UpdateRecord(v.DbConnection, "crawler_url", record.Id, datatypes.JSON(objBytes), record.CreatedBy, "Active")
					continue
				}

				doc, err := goquery.NewDocumentFromReader(resp.Body)
				if err != nil {
					obj["checkbox"] = false
					obj["status"] = "Error"
					objBytes, _ = json.Marshal(obj)
					UpdateRecord(v.DbConnection, "crawler_url", record.Id, datatypes.JSON(objBytes), record.CreatedBy, "Active")
					continue
				}

				// HTML version
				htmlVersion := ""
				doc.Find("html").Each(func(i int, s *goquery.Selection) {
					if n := s.Nodes[0]; n != nil && n.PrevSibling != nil && n.PrevSibling.Type == 4 /* DoctypeNode */ {
						htmlVersion = n.PrevSibling.Data
					}
				})
				if htmlVersion == "" {
					htmlVersion = "Unknown"
				}

				// Title
				title := strings.TrimSpace(doc.Find("title").First().Text())

				// Headings count
				headings := map[string]int{
					"h1": doc.Find("h1").Length(),
					"h2": doc.Find("h2").Length(),
					"h3": doc.Find("h3").Length(),
					"h4": doc.Find("h4").Length(),
					"h5": doc.Find("h5").Length(),
					"h6": doc.Find("h6").Length(),
				}

				// Links
				baseUrl := url
				internalLinks := 0
				externalLinks := 0
				inaccessibleLinks := 0
				links := []string{}
				doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
					href, _ := s.Attr("href")
					if strings.HasPrefix(href, "http") {
						if strings.HasPrefix(href, baseUrl) {
							internalLinks++
						} else {
							externalLinks++
						}
						links = append(links, href)
					} else if strings.HasPrefix(href, "/") {
						internalLinks++
						links = append(links, baseUrl+href)
					}
				})

				// Check inaccessible links (4xx/5xx) and collect broken links with status codes
				brokenLinks := []map[string]interface{}{}
				for _, link := range links {
					r, err := http.Head(link)
					if err != nil {
						inaccessibleLinks++
						brokenLinks = append(brokenLinks, map[string]interface{}{
							"url":    link,
							"status": 0, // 0 for network error
						})
					} else if r.StatusCode >= 400 && r.StatusCode < 600 {
						inaccessibleLinks++
						brokenLinks = append(brokenLinks, map[string]interface{}{
							"url":    link,
							"status": r.StatusCode,
						})
					}
					if r != nil {
						r.Body.Close()
					}
				}

				// If stopped during link checking, skip the rest
				checkedCrawlerStatus = GetStatusOfCrawler(v.DbConnection, record.Id)
				if !checkedCrawlerStatus {
					continue
				}

				// Login form detection
				hasLoginForm := false
				doc.Find("form").EachWithBreak(func(i int, s *goquery.Selection) bool {
					if s.Find("input[type='password']").Length() > 0 {
						hasLoginForm = true
						return false
					}
					return true
				})

				// Update object_info with new analysis and status
				obj["html_version"] = htmlVersion
				obj["title"] = title
				obj["headings"] = headings
				obj["internal_links"] = internalLinks
				obj["external_links"] = externalLinks
				obj["inaccessible_links"] = inaccessibleLinks
				obj["has_login_form"] = hasLoginForm
				obj["broken_links"] = brokenLinks
				obj["status"] = "Done"
				obj["checkbox"] = false
				objBytes, _ = json.Marshal(obj)

				UpdateRecord(v.DbConnection, "crawler_url", record.Id, datatypes.JSON(objBytes), record.CreatedBy, "Active")
			}
		}
	}

}
