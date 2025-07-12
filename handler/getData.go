package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Rishoban/crawler-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type CrawlerUrl struct {
	ID           int            `json:"id"`
	ObjectInfo   datatypes.JSON `json:"object_info"`
	ObjectStatus string         `json:"object_status"`
}

func (v *CrawlerService) GetAllUrls(ctx *gin.Context) {
	var records = []CrawlerUrl{}
	if err := v.DbConnection.Table("crawler_url").Find(&records).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch records: " + err.Error()})
		return
	}

	var urls = make([]map[string]interface{}, 0)
	for _, rec := range records {
		var crawlResp model.CrawlResponse
		if err := json.Unmarshal(rec.ObjectInfo, &crawlResp); err != nil {
			continue // skip invalid records
		}

		if rec.ObjectStatus == "Active" {
			urls = append(urls, map[string]interface{}{
				"id":                rec.ID,
				"url":               crawlResp.Url,
				"htmlVersion":       crawlResp.HTMLVersion,
				"title":             crawlResp.Title,
				"headings":          crawlResp.Headings,
				"internalLinks":     crawlResp.InternalLinks,
				"externalLinks":     crawlResp.ExternalLinks,
				"inaccessibleLinks": crawlResp.InaccessibleLinks,
				"hasLoginForm":      crawlResp.HasLoginForm,
				"checkbox":          crawlResp.Checkbox,
				"status":            crawlResp.Status,
			})
		}

	}

	ctx.JSON(http.StatusOK, gin.H{"urls": urls})
}
