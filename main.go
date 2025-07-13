package main

import (
	"fmt"
	"os"

	"github.com/Rishoban/crawler-api/handler"
	"github.com/Rishoban/crawler-api/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadConfig(path string) (*model.AppConfig, error) {
	var cfg model.AppConfig

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ConnectDB(user, password, host string, port int, dbname string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
	}

	return db
}

// InitStandaloneService init standalone service needed to have gin common router to support API end points.
func InitStandaloneService(router *gin.Engine) error {
	var cfg model.AppConfig

	configData, err := os.ReadFile("conf/config.yml")
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(configData, &cfg); err != nil {
		return err
	}

	db := ConnectDB(cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	var crawlerService = &handler.CrawlerService{DbConnection: db}

	crawlerService.InitRouter(router)

	return nil
}

func main() {
	router := gin.Default()

	// Set up CORS to allow requests from localhost:5173
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Initialize the standalone service
	if err := InitStandaloneService(router); err != nil {
		fmt.Println("Failed to initialize crawler service:", err)
	}

	// Start the server
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
