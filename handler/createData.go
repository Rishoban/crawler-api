package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Rishoban/crawler-api/model"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func (v *CrawlerService) CreateNewRecord(ctx *gin.Context) {
	var req model.CrawlRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Prepare the default object_info structure
	objectInfo := map[string]interface{}{
		"url":      req.URL,
		"title":    "",
		"status":   "Pending",
		"checkbox": true,
		"headings": map[string]int{
			"h1": 0,
			"h2": 0,
			"h3": 0,
			"h4": 0,
			"h5": 0,
			"h6": 0,
		},
		"html_version":        "",
		"external_links":      0,
		"has_login_form":      false,
		"internal_links":      0,
		"inaccessible_links":  0,
		"presenceOfLoginForm": false,
		"broken_links":        make([]map[string]interface{}, 0),
	}

	// Marshal objectInfo to JSON for object_info
	jsonBytes, err := json.Marshal(objectInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create record"})
		return
	}

	// Insert into object_info table using your reusable function
	_, err = CreateRecord(v.DbConnection, "crawler_url", datatypes.JSON(jsonBytes), 1, "Active")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create record: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Record is successfully created"})
}
