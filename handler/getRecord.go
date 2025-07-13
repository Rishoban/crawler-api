package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /crawler/record/:id
func (v *CrawlerService) GetRecordByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record id"})
		return
	}

	record, err := GetRecordByID(v.DbConnection, "crawler_url", id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(record.ObjectInfo, &obj); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse record data"})
		return
	}

	internal, _ := obj["internal_links"].(float64) // JSON numbers are float64
	external, _ := obj["external_links"].(float64)
	brokenLinks, _ := obj["broken_links"].([]interface{})

	linksData := []map[string]interface{}{}
	for _, item := range brokenLinks {
		if m, ok := item.(map[string]interface{}); ok {
			linksData = append(linksData, map[string]interface{}{
				"url":    m["url"],
				"status": m["status"],
			})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"title": obj["url"],
		"chartData": gin.H{
			"internal": int(internal),
			"external": int(external),
		},
		"linksData": linksData,
	})
}
