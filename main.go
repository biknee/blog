package main

import (
	"html/template"
	"log"

	"blog/handlers"
	"blog/services"

	"github.com/gin-gonic/gin"
)

func main() {
	svc := services.NewPostService()
	h := handlers.NewPostHandler(svc)

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"add":      func(a, b int) int { return a + b },
		"subtract": func(a, b int) int { return a - b },
	})
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/", h.Home)
	r.GET("/post/:slug", h.View)
	r.GET("/new", h.NewForm)
	r.POST("/post", h.Create)
	r.GET("/edit/:slug", h.EditForm)
	r.POST("/post/:slug", h.Update)
	r.POST("/post/:slug/save-draft", h.SaveDraft)
	r.POST("/post/:slug/delete", h.Delete)
	r.GET("/tag/:tag", h.Tag)

	log.Println("Blog running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
