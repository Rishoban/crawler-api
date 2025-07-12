package handler

import (
	"encoding/json"

	"gorm.io/gorm"
)

// GetStatusOfCrawler checks if the crawler status is "Running" for the given ID
func GetStatusOfCrawler(dbConnection *gorm.DB, crawlerId int) bool {
	generalObject, err := GetRecordByID(dbConnection, "crawler_url", crawlerId)
	if err != nil {
		return false
	}
	var obj map[string]interface{}
	err = json.Unmarshal(generalObject.ObjectInfo, &obj)
	if err != nil {
		return false
	}
	status, _ := obj["status"].(string)

	return status == "Running"
}
