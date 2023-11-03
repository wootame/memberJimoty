package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Home ハンドラはルートエンドポイントのGETリクエストを処理します。
func Home(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Home Page"})
}

// GetUserByID ハンドラは指定されたユーザーIDに対応するユーザー情報を取得します。
func GetUserByID(c *gin.Context) {
    userID := c.Param("id") // URLからIDを取得

    // ユーザー情報をデータベースから取得するなどの処理を行う

    // ダミーデータを返す例
    user := gin.H{
        "id":   userID,
        "name": "John Doe",
    }
    c.JSON(http.StatusOK, user)
}

// GetUsers ハンドラはすべてのユーザー情報を取得します。
func GetUsers(c *gin.Context) {
    // ユーザー情報をデータベースから取得するなどの処理を行う

    // ダミーデータを返す例
    users := []gin.H{
        {
            "id":   "1",
            "name": "Alice",
        },
        {
            "id":   "2",
            "name": "Bob",
        },
    }
	c.JSON(http.StatusOK, users)
}

// CreateUser ハンドラは新しいユーザーを作成します。
func CreateUser(c *gin.Context) {
    // リクエストから新しいユーザーの情報を受け取り、データベースに保存するなどの処理を行う

    // ダミーデータを返す例
    newUser := gin.H{
        "id":   "3",
        "name": "New User",
    }
    c.JSON(http.StatusCreated, newUser)
}

// UpdateUser ハンドラは指定されたユーザーの情報を更新します。
func UpdateUser(c *gin.Context) {
    userID := c.Param("id") // URLからIDを取得

    // リクエストから新しいユーザー情報を受け取り、データベースを更新するなどの処理を行う

    // ダミーデータを返す例
    updatedUser := gin.H{
        "id":   userID,
        "name": "Updated User",
    }
    c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser ハンドラは指定されたユーザーを削除します。
func DeleteUser(c *gin.Context) {
    // userID := c.Param("id") // URLからIDを取得

    // データベースからユーザーを削除するなどの処理を行う

    c.JSON(http.StatusNoContent, nil)
}
