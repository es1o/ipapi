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

	// Define a simple GET endpoint
	r.GET("/", func(c *gin.Context) {
		cfIP := c.GetHeader("Cf-Connecting-Ip")
		ip, err := netip.ParseAddr(cfIP)
		if err != nil {
			slog.Error("failed to parse IP", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to parse IP",
			})
			return
		}

		db, _ := maxminddb.Open("ip66.mmdb")
		var record any
		err = db.Lookup(ip).Decode(&record)
		if err != nil {
			slog.Error("failed to decode record", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to decode record",
			})
			return
		}
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"ip":     ip,
			"record": record,
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
