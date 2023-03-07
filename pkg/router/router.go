package router

import (
	"net/http"

	"github.com/choisangh/image_crud_api/pkg/api"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("public/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	images := r.Group("/images")
	{
		images.POST("", api.CreateImage)
		images.GET("/:id", api.ReadImage)
		images.PUT("/:id", api.PutImage)
		images.DELETE("/:id", api.DeleteImage)
	}
	return r
}
