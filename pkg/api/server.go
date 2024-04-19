package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	db "github.com/lushenle/plam/pkg/db"
	"github.com/lushenle/plam/pkg/token"
	"github.com/lushenle/plam/pkg/util"
	"github.com/penglongli/gin-metrics/ginmetrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// Server serves HTTP requests for our PanGu Operating System
type Server struct {
	router     *gin.Engine
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	logger     *zap.Logger
}

type ServerOption func(server *Server)

func NewServer(config util.Config, opts ...ServerOption) *Server {
	server := &Server{
		config: config,
	}

	for _, opt := range opts {
		opt(server)
	}

	server.setupRouter()

	return server
}

func WithStore(store db.Store) ServerOption {
	return func(server *Server) {
		server.store = store
	}
}

func WithTokenMaker(maker token.Maker) ServerOption {
	return func(server *Server) {
		server.tokenMaker = maker
	}
}

func WithLogger(logger *zap.Logger) ServerOption {
	return func(server *Server) {
		server.logger = logger
	}
}

func (server *Server) setupRouter() {
	router := gin.New()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log
	//   - Logs to stdout
	//   - RFC3339 with UTC time format
	router.Use(ginzap.Ginzap(server.logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info
	router.Use(ginzap.RecoveryWithZap(server.logger, true))

	router.Use(gzip.Gzip(gzip.BestSpeed))

	// Create a Prometheus instance
	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)

	// Configure the Swagger UI with basic authentication
	swaggerConfig := &ginSwagger.Config{
		URL: "http://localhost/swagger/doc.json", // Update this URL based on project structure
	}
	router.GET("/swagger/*any", ginSwagger.CustomWrapHandler(swaggerConfig, swaggerFiles.Handler))

	apiV1 := router.Group("/v1")

	{
		apiV1.GET("/healthz", server.healthz)
		apiV1.POST("/users/sign", server.signupUser)
		apiV1.POST("/users/login", server.loginUser)
	}

	authRoutes := apiV1.Group("/").Use(authMiddleware(server.tokenMaker))

	server.router = router
}

// Start runs the HTTP server in a specific address
func (server *Server) Start(addr string) error {
	// Set the connection timeout
	srv := &http.Server{
		Addr:         addr,
		Handler:      server.router,
		ReadTimeout:  360 * time.Second, // Maximum duration for reading the entire request
		WriteTimeout: 360 * time.Second, // Maximum duration before timing out writes of the response
		IdleTimeout:  720 * time.Second, // Maximum amount of time to wait for the next request when keep-alives are enabled
	}

	// Run the server
	return srv.ListenAndServe()
}

func (server *Server) healthz(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
