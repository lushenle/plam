package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/gzip"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/lushenle/plam/docs"
	"github.com/lushenle/plam/pkg/db"
	"github.com/lushenle/plam/pkg/token"
	"github.com/lushenle/plam/pkg/util"
	"github.com/penglongli/gin-metrics/ginmetrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// Server serves HTTP requests for Paramount Construction Machinery System
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
		URL: "http://localhost:8080/swagger/doc.json", // Update this URL based on project structure
	}
	router.GET("/swagger/*any", ginSwagger.CustomWrapHandler(swaggerConfig, swaggerFiles.Handler))

	apiV1 := router.Group("/v1")

	{
		apiV1.GET("/healthz", server.healthz)
		apiV1.POST("/users/signup", server.signupUser)
		apiV1.POST("/users/login", server.loginUser)
	}

	authRoutes := apiV1.Group("/").Use(authMiddleware(server.tokenMaker))

	// projects router
	{
		authRoutes.POST("/projects/all", server.listProjects)
		authRoutes.GET("/projects/:id", server.getProject)
		authRoutes.POST("/projects/search", server.searchProjects)
	}

	// incomes router
	{
		authRoutes.POST("/incomes/all", server.listIncomes)
		authRoutes.GET("/incomes/:id", server.getIncome)
		authRoutes.POST("/incomes/search", server.searchIncomes)
	}

	// loans router
	{
		authRoutes.POST("/loans/all", server.listLoans)
		authRoutes.GET("/loans/:id", server.getLoan)
		authRoutes.POST("/loans/search", server.searchLoans)
	}

	// pay_outs router
	{
		authRoutes.POST("/pay_outs/all", server.listPayOuts)
		authRoutes.GET("/pay_outs/:id", server.getPayOut)
		authRoutes.POST("/pay_outs/search", server.searchPayOuts)
	}

	authRoutes.Use(rbacMiddleware())

	{
		authRoutes.POST("/projects", server.createProject)
		authRoutes.DELETE("/projects/:id", server.deleteProject)

		authRoutes.POST("/incomes", server.createIncome)
		authRoutes.DELETE("/incomes/:id", server.deleteIncome)

		authRoutes.POST("/loans", server.createLoan)
		authRoutes.DELETE("/loans/:id", server.deleteLoan)

		authRoutes.POST("/pay_outs", server.createPayOut)
		authRoutes.DELETE("/pay_outs/:id", server.deletePayOut)
	}

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

// errorResponse is a struct that represents an error response.
//
// @swagger:model
type errorResponse struct {
	Error string `json:"error"`
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
