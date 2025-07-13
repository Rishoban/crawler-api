package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteRequest struct {
	URLs []int `json:"urls"`
}

// POST /archive-crawlers
func (v *CrawlerService) ArchiveCrawlers(ctx *gin.Context) {
	var req DeleteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	for _, id := range req.URLs {

		record, err := GetRecordByID(v.DbConnection, "crawler_url", id)
		if err != nil {
			continue
		}

		// Only update object_status to 'Archived'
		err = UpdateRecord(v.DbConnection, "crawler_url", id, record.ObjectInfo, 0, "Archived")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to archive id: " + string(rune(id))})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Selected records archived successfully."})
}
