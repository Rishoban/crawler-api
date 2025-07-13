package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

// {"urls": [1,2,5]}
type ReRunRequest struct {
	URLs []int `json:"url"`
}

// POST /rerun-analysis
func (v *CrawlerService) ReRunAnalysis(ctx *gin.Context) {
	var req ReRunRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// For each URL, start crawling and check for stop signal
	for _, id := range req.URLs {
		// Fetch the record
		record, err := GetRecordByID(v.DbConnection, "crawler_url", id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found for id: " + string(rune(id))})
			continue
		}

		// Set status to Running
		var obj map[string]interface{}
		_ = json.Unmarshal(record.ObjectInfo, &obj)
		if obj == nil {
			obj = map[string]interface{}{}
		}
		obj["checkbox"] = false
		obj["status"] = "Queued"
		objBytes, _ := json.Marshal(obj)
		UpdateRecord(v.DbConnection, "crawler_url", id, datatypes.JSON(objBytes), record.CreatedBy, "Active")

	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Re-analysis started for provided URLs. Status will be updated."})
}
