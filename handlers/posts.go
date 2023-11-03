package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Home ハンドラはルートエンドポイントのGETリクエストを処理します。
func Posts(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Home Page"})
}