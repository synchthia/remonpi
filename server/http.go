package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/synchthia/remonpi/controller/mitsubishi/kgsa3c"
	"github.com/synchthia/remonpi/logger"
	"github.com/synchthia/remonpi/models"
	"github.com/synchthia/remonpi/remote"
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
	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))
	r.Use(logger.SetLogger())

	r.GET("/api/v1/test", h.getTest)

	// State
	r.GET("/api/v1/state", h.getState)

	// Remote
	r.GET("/api/v1/remote", h.getRemote)
	r.POST("/api/v1/remote", h.postRemote)

	// Template
	r.GET("/api/v1/template", h.getTemplate)

	r.Use(static.Serve("/", static.LocalFile("./public", false)))

	return r
}

func (h *httpServer) getTest(c *gin.Context) {
	c.String(http.StatusOK, "Test, 123...")
}

func (h *httpServer) getState(c *gin.Context) {
	state := h.Remote.GetState()
	c.JSON(http.StatusOK, state)
}

func (h *httpServer) getRemote(c *gin.Context) {
	state := h.Remote.GetState()

	mode := c.Query("mode")
	if mode == "" {
		c.JSON(http.StatusOK, state.ToRemoteData())
	} else {
		c.JSON(http.StatusOK, state.ToRemoteDataByMode(mode))
	}
}

func (h *httpServer) postRemote(c *gin.Context) {
	remoteData := &models.RemoteData{}

	// Inject RemoteData
	if err := c.BindJSON(&remoteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Remote.Send(remoteData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	state := h.Remote.GetState().ToRemoteData()
	c.JSON(http.StatusOK, state)
}

func (h *httpServer) getTemplate(c *gin.Context) {
	if h.Remote.Vendor == "mitsubishi" && h.Remote.Model == "kgsa3-c" {
		c.JSON(http.StatusOK, kgsa3c.TemplateData)
	}
}
