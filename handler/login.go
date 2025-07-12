package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// POST /login
func (v *CrawlerService) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	// Find user by email
	var user GeneralObject
	err := v.DbConnection.Table("user").Where("JSON_UNQUOTE(JSON_EXTRACT(object_info, '$.email')) = ?", req.Email).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	// Parse object_info
	var obj map[string]interface{}
	// If ObjectInfo is datatypes.JSON, use .Bytes() for unmarshalling
	var objBytes []byte
	switch v := any(user.ObjectInfo).(type) {
	case []byte:
		objBytes = v
	case string:
		objBytes = []byte(v)
	case interface{ Bytes() []byte }:
		objBytes = v.Bytes()
	default:
		objBytes, _ = json.Marshal(user.ObjectInfo)
	}
	if err := json.Unmarshal(objBytes, &obj); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}
	hash, ok := obj["password"].(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.Id,
		"email":   req.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	// Use a secure secret in production!
	secret := []byte("your_secret_key")
	tokenString, err := token.SignedString(secret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}
