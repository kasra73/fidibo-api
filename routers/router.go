package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/kasra73/fidibo-api/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/kasra73/fidibo-api/pkg/export"
	"github.com/kasra73/fidibo-api/routers/api"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))

	r.POST("/auth", api.Authenticate)
	r.GET("/search/book", api.SearchBooks)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
