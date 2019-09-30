package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/synchthia/remonpi/logger"
	"github.com/synchthia/remonpi/remote"
	"github.com/synchthia/remonpi/remote/mitsubishi/kgsa3c"
)

type httpServer struct {
	Remote *remote.Remote
}

// NewHTTPServer - Start HTTP Server
func NewHTTPServer(remote *remote.Remote) *gin.Engine {
	h := httpServer{
		Remote: remote,
	}
	r := gin.Default()
	r.Use(logger.SetLogger())

	r.GET("/api/v1/test", h.getTest)

	// Remote
	r.GET("/api/v1/remote", h.getRemote)
	r.POST("/api/v1/remote", h.postRemote)

	// Template
	r.GET("api/v1/template", h.getTemplate)

	return r
}

func (h *httpServer) getTest(c *gin.Context) {
	c.String(http.StatusOK, "Test, 123...")
}

func (h *httpServer) getRemote(c *gin.Context) {

}

func (h *httpServer) postRemote(c *gin.Context) {
	// Detect vendor / model
	if h.Remote.Vendor == "mitsubishi" && h.Remote.Model == "kgsa3-c" {
		var controller kgsa3c.Controller
		if err := c.ShouldBindJSON(&controller); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate & Send IR Signal
		if err := kgsa3c.Send(&controller); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}

func (h *httpServer) getTemplate(c *gin.Context) {
	if h.Remote.Vendor == "mitsubishi" && h.Remote.Model == "kgsa3-c" {
		c.JSON(http.StatusOK, kgsa3c.GetTemplate())
	}
}
