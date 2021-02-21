package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robert-inkpen/randeng_co_mirror/backend/config"
	"github.com/robert-inkpen/randeng_co_mirror/backend/db"
)

// Webserver will abstract the HTTP server and database that supports the API
type Webserver struct {
	g    *gin.Engine
	db   *db.Database
	Addr string
	Port string
}

// NewServer initializes and returns a Webserver object
func NewServer(cfg *config.Config) *Webserver {
	return &Webserver{
		g:    setupRoutes(),
		db:   &db.Database{},
		Addr: cfg.ServerConfig.Address,
		Port: cfg.ServerConfig.Port,
	}
}

// setupRoutes builds the underlying gin engine and setups all the route handlers and middleware
func setupRoutes() *gin.Engine {
	engine := gin.Default()
	router := engine.Group("/api")
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return engine
}

// Start starts the webserver on the configured port and address
func (s *Webserver) Start() error {
	return s.g.Run(fmt.Sprintf("%v:%v", s.Addr, s.Port))
}

// ServerHTTP wrapper around the gin internal ServeHTTP
func (s *Webserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.g.ServeHTTP(w, r)
	return
}
