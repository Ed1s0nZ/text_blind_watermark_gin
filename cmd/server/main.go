package main

import (
	"log"
	"net/http"
	"text_blind_watermark/internal/watermark"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 提供静态文件
	r.Static("/static", "./web/static")
	r.LoadHTMLGlob("web/templates/*")

	// 主页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// API路由
	api := r.Group("/api")
	{
		api.POST("/encode", func(c *gin.Context) {
			var req struct {
				Text      string `json:"text"`
				Watermark string `json:"watermark"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			encoded := watermark.Encode(req.Text, req.Watermark)
			c.JSON(http.StatusOK, gin.H{"result": encoded})
		})

		api.POST("/decode", func(c *gin.Context) {
			var req struct {
				Text string `json:"text"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			decoded := watermark.Decode(req.Text)
			c.JSON(http.StatusOK, gin.H{"result": decoded})
		})
	}

	// 启动服务器
	log.Println("Server starting on http://localhost:8080")
	r.Run(":8080")
}
