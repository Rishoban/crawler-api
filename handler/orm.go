package handler

import (
	"database/sql"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ResourceTime is a custom time type for resource timestamps (you can adjust as needed)
type ResourceTime time.Time

// GeneralObject is a generic struct for tables with the given structure
type GeneralObject struct {
	Id            int            `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt     ResourceTime   `json:"created_at" gorm:"type:TIMESTAMP;null;default:'CURRENT_TIMESTAMP'"`
	LastUpdatedAt ResourceTime   `json:"last_updated_at" gorm:"type:TIMESTAMP;null;default:CURRENT_TIMESTAMP"`
	CreatedBy     int            `json:"created_by" gorm:"type:int"`
	LastUpdatedBy int            `json:"lastUpdated_by" gorm:"type:int"`
	ObjectStatus  string         `json:"object_status" gorm:"type:varchar(45)"`
	ObjectInfo    datatypes.JSON `json:"object_info" gorm:"type:json"`
}

func (GeneralObject) TableName() string {
	return "inventory_item_description" // override as needed
}

// CreateRecord inserts a new record into the given table using GORM
func CreateRecord(db *gorm.DB, tableName string, objectInfo datatypes.JSON, createdBy int, objectStatus string) (int, error) {
	record := GeneralObject{
		ObjectInfo:   objectInfo,
		CreatedBy:    createdBy,
		ObjectStatus: objectStatus,
	}
	tx := db.Table(tableName).Create(&record)
	return record.Id, tx.Error
}

// UpdateRecord updates a record by id using GORM
func UpdateRecord(db *gorm.DB, tableName string, id int, objectInfo datatypes.JSON, updatedBy int, objectStatus string) error {
	updates := map[string]interface{}{
		"object_info":     objectInfo,
		"last_updated_by": updatedBy,
		"object_status":   objectStatus,
		"last_updated_at": gorm.Expr("NOW()"),
	}
	return db.Table(tableName).Where("id = ?", id).Updates(updates).Error
}

// DeleteRecord deletes a record by id using GORM
func DeleteRecord(db *gorm.DB, tableName string, id int) error {
	return db.Table(tableName).Where("id = ?", id).Delete(&GeneralObject{}).Error
}

// GetRecordByID fetches a record by id using GORM
func GetRecordByID(db *gorm.DB, tableName string, id int) (*GeneralObject, error) {
	var record GeneralObject
	err := db.Table(tableName).Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// GetAllRecords fetches all records from a table using GORM
func GetAllRecords(db *gorm.DB, tableName string) ([]GeneralObject, error) {
	var records []GeneralObject
	err := db.Table(tableName).Find(&records).Error
	return records, err
}

// GetRecordsByCondition fetches records by a custom condition string using GORM
func GetRecordsByCondition(db *gorm.DB, tableName string, condition string, args ...interface{}) ([]GeneralObject, error) {
	var records []GeneralObject
	err := db.Table(tableName).Where(condition, args...).Find(&records).Error
	return records, err
}

// scanRows scans multiple rows into a slice of maps
func scanRows(rows *sql.Rows) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	for rows.Next() {
		var id, createdBy, lastUpdatedBy sql.NullInt64
		var objectInfo sql.NullString
		var createdAt, lastUpdatedAt sql.NullTime
		var objectStatus sql.NullString
		err := rows.Scan(&id, &objectInfo, &createdAt, &lastUpdatedAt, &createdBy, &lastUpdatedBy, &objectStatus)
		if err != nil {
			return nil, err
		}
		results = append(results, map[string]interface{}{
			"id":              id.Int64,
			"object_info":     objectInfo.String,
			"created_at":      createdAt.Time,
			"last_updated_at": lastUpdatedAt.Time,
			"created_by":      createdBy.Int64,
			"last_updated_by": lastUpdatedBy.Int64,
			"object_status":   objectStatus.String,
		})
	}
	return results, nil
}
