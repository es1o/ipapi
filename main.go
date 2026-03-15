package main

import (
	"log/slog"
	"net/http"
	"net/netip"
	"os"

	"github.com/gin-gonic/gin"
	maxminddb "github.com/oschwald/maxminddb-golang/v2"
	sloggin "github.com/samber/slog-gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	r := gin.New()
	r.Use(sloggin.New(logger))
	r.Use(gin.Recovery())

	// Get API key from environment variable
	API_KEY := os.Getenv("API_KEY")
	if API_KEY == "" {
		logger.Warn("API_KEY is not set")
	} else {
		logger.Info("API_KEY is set", "API_KEY", API_KEY)
	}

	// Define a simple GET endpoint
	r.GET("/", func(c *gin.Context) {
		if API_KEY != "" && c.GetHeader("X-Api-Key") != API_KEY {
			logger.Error("Unauthorized", "API_KEY", API_KEY, "X-Api-Key", c.GetHeader("X-Api-Key"))
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		cfIP := c.GetHeader("Cf-Connecting-Ip")
		ip, err := netip.ParseAddr(cfIP)
		if err != nil {
			logger.Error("failed to parse IP", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to parse IP",
			})
			return
		}

		db, _ := maxminddb.Open("ip66.mmdb")
		var record any
		err = db.Lookup(ip).Decode(&record)
		if err != nil {
			logger.Error("failed to decode record", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to decode record",
			})
			return
		}
		var cityRecord any
		cityDb, err := maxminddb.Open("city.mmdb")
		if err != nil {
			logger.Error("failed to open city database", "error", err)
		} else {
			err = cityDb.Lookup(ip).Decode(&cityRecord)
			if err != nil {
				logger.Error("failed to decode city record", "error", err)
			}
		}
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"ip":     ip,
			"record": record,
			"city":   cityRecord,
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	if err := r.Run(); err != nil {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}
