package logger

// --------------------------------------------------------------------
// NOTE:
// This code forked from : https://github.com/gin-contrib/logger
// --------------------------------------------------------------------

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Config struct {
	// UTC a boolean stating whether to use UTC time zone or local.
	UTC            bool
	SkipPath       []string
	SkipPathRegexp *regexp.Regexp
}

// SetLogger initializes the logging middleware.
func SetLogger(config ...Config) gin.HandlerFunc {
	var newConfig Config
	if len(config) > 0 {
		newConfig = config[0]
	}
	var skip map[string]struct{}
	if length := len(newConfig.SkipPath); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range newConfig.SkipPath {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		c.Next()
		track := true

		if _, ok := skip[path]; ok {
			track = false
		}

		if track &&
			newConfig.SkipPathRegexp != nil &&
			newConfig.SkipPathRegexp.MatchString(path) {
			track = false
		}

		if track {
			end := time.Now()
			latency := end.Sub(start)
			if newConfig.UTC {
				end = end.UTC()
			}

			msg := "[Gin] Request"
			if len(c.Errors) > 0 {
				msg = c.Errors.String()
			}

			dumplogger := logrus.WithFields(logrus.Fields{
				"status":     c.Writer.Status(),
				"method":     c.Request.Method,
				"path":       path,
				"ip":         c.ClientIP(),
				"latency":    latency,
				"user-agent": c.Request.UserAgent(),
			})

			switch {
			case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
				{
					dumplogger.Warn(msg)
				}
			case c.Writer.Status() >= http.StatusInternalServerError:
				{
					dumplogger.Error(msg)
				}
			default:
				dumplogger.Info(msg)
			}
		}
	}
}
