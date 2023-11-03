package routers

import (
    "github.com/gin-gonic/gin"
    "memberProject/handlers" // ハンドラ関数が定義されているパッケージをインポート
)

func SetupRouter(router *gin.Engine) {
    // ルートエンドポイントへのGETリクエストを処理する
    router.GET("/", handlers.Home)

    // /api パスプレフィックスの下にエンドポイントを追加する例
    api := router.Group("/api")
    {
        api.GET("/users", handlers.GetUsers)
        api.POST("/users", handlers.CreateUser)
        api.GET("/users/:id", handlers.GetUserByID)
        api.PUT("/users/:id", handlers.UpdateUser)
        api.DELETE("/users/:id", handlers.DeleteUser)
		api.GET("/post", handlers.Posts)
    }

    // その他のルーティングルールを追加できます
}
