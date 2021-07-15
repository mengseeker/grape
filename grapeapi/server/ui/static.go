package ui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed static
var ui embed.FS

func Mount(e *gin.Engine) {
	e.Handle("GET", "/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/ui/")
	})
	web, _ := fs.Sub(ui, "static")
	e.StaticFS("/ui", http.FS(web))
}
