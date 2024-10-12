package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// ServiceConfig holds the mapping of paths to service URLs
type ServiceConfig struct {
	ServiceMap map[string]string
}

func main() {
	r := gin.Default()

	// Initialize service configuration
	serviceConfig := ServiceConfig{
		ServiceMap: map[string]string{
			"/books": "http://localhost:8082", // BookService URL
			"/users": "http://localhost:8081", // UserService URL
			// Add other services here as needed
		},
	}

	// Route for dynamic proxying
	r.Any("/*proxyPath", func(c *gin.Context) {
		proxyPath := c.Param("proxyPath")
		log.Printf("Received request for: %s", proxyPath) // Log incoming request path

		// Check for the service in the configuration
		for route, serviceURL := range serviceConfig.ServiceMap {
			if strings.HasPrefix(proxyPath, route) {
				// Parse the service URL
				url, err := url.Parse(serviceURL)
				if err != nil {
					log.Printf("Error parsing service URL: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Service not available"})
					return
				}

				// Create the reverse proxy
				proxy := httputil.NewSingleHostReverseProxy(url)

				// Update the request URL
				c.Request.URL.Host = url.Host
				c.Request.URL.Scheme = url.Scheme
				// Keep the original path
				c.Request.URL.Path = proxyPath

				// Update the request Host header
				c.Request.Host = url.Host

				// Log the forwarding request
				log.Printf("Forwarding request to %s%s", url.String(), c.Request.URL.Path)

				// Serve the request
				proxy.ServeHTTP(c.Writer, c.Request)
				return
			}
		}

		// If no matching route found, return a 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
	})

	// Start the API Gateway
	if err := r.Run(":8080"); err != nil { // Port for the API Gateway
		log.Fatal(err)
	}
}
