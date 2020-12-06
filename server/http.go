package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"github.com/synchthia/remonpi/logger"
	"github.com/synchthia/remonpi/models"
	"github.com/synchthia/remonpi/remote"
)

type httpServer struct {
	Remote *remote.Remote
}

type statikFileSystem struct {
	fs http.FileSystem
}

func (b *statikFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *statikFileSystem) Exists(prefix string, filepath string) bool {
	if _, err := b.fs.Open(filepath); err != nil {
		return false
	}
	return true
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

	r.GET("/healthz", h.getHealthz)

	// State
	r.GET("/api/v1/state", h.getState)

	// Remote
	r.GET("/api/v1/remote", h.getRemote)
	r.POST("/api/v1/remote", h.postRemote)

	// Template
	r.GET("/api/v1/template", h.getTemplate)

	statikFS, err := fs.New()
	if err != nil {
		logrus.WithError(err).Fatal("[Static]")
	}

	r.Use(static.Serve("/", &statikFileSystem{
		statikFS,
	}))

	// LocalFile : r.Use(static.Serve("/", static.LocalFile("./public", false)))

	return r
}

func (h *httpServer) getHealthz(c *gin.Context) {
	c.String(http.StatusOK, "OK")
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
		logrus.WithError(err).Errorf("Failed bind RemoteData")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Remote.Send(remoteData)
	if err != nil {
		logrus.WithError(err).Errorf("Failed send signal")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	state := h.Remote.GetState().ToRemoteData()
	c.JSON(http.StatusOK, state)
}

func (h *httpServer) getTemplate(c *gin.Context) {
	c.JSON(http.StatusOK, h.Remote.Template)
}
