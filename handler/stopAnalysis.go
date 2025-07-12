package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POST /stop-analysis
func (v *CrawlerService) StopAnalysis(ctx *gin.Context) {
	var req ReRunRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	for _, id := range req.URLs {
		record, err := GetRecordByID(v.DbConnection, "crawler_url", id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found for id: " + string(rune(id))})
			continue
		}

		// Set status to Running
		var obj map[string]interface{}
		_ = json.Unmarshal(record.ObjectInfo, &obj)

		obj["checkbox"] = false
		obj["status"] = "Stopped"
		objBytes, _ := json.Marshal(obj)
		UpdateRecord(v.DbConnection, "crawler_url", id, objBytes, record.CreatedBy, "Active")
		continue
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Stop signal sent for provided URLs."})
}
