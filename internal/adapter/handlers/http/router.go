package http

import (
	"log/slog"
	"net/http"

	_ "github.com/behrouz-rfa/kentech/docs"
	"github.com/behrouz-rfa/kentech/internal/adapter/config"
	"github.com/behrouz-rfa/kentech/internal/core/ports"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zsais/go-gin-prometheus"
)

// Router is a wrapper for the HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(
	config *config.Config,
	auth ports.Auth,
	userHandler UserHandler,
	filmHandler FilmHandler,
) (*Router, error) {
	router := gin.New()

	setupPrometheus(router)
	setupMiddleware(router)
	setupSwagger(router)
	defineRoutes(router, auth, userHandler, filmHandler)

	return &Router{router}, nil
}

// setupPrometheus sets up Prometheus metrics
func setupPrometheus(router *gin.Engine) {
	customMetrics := []*ginprometheus.Metric{
		{
			ID:          "test_metric",
			Name:        "test_metric",
			Description: "Counter test metric",
			Type:        "counter",
		},
		{
			ID:          "test_metric_2",
			Name:        "test_metric_2",
			Description: "Summary test metric",
			Type:        "summary",
		},
	}

	p := ginprometheus.NewPrometheus("gin", customMetrics)
	p.Use(router)
}

// setupMiddleware sets up middleware for the router
func setupMiddleware(router *gin.Engine) {
	router.Use(sloggin.New(slog.Default()), gin.Recovery())
}

// setupSwagger sets up Swagger documentation
func setupSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// defineRoutes defines the API routes
func defineRoutes(router *gin.Engine, auth ports.Auth, userHandler UserHandler, filmHandler FilmHandler) {
	// Disable debug mode in production
	if gin.Mode() == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/register", userHandler.Register)
			user.POST("/login", userHandler.Login)

			authUser := user.Use(authMiddleware(auth))
			{
				authUser.GET("", userHandler.ListUsers)
				authUser.GET("/:id", userHandler.GetUser)

				admin := authUser.Use(adminMiddleware())
				{
					admin.PUT("/:id", userHandler.UpdateUser)
					admin.DELETE("/:id", userHandler.DeleteUser)
				}
			}
		}
		films := v1.Group("/films").Use(authMiddleware(auth))
		{
			films.GET("", filmHandler.ListFilms)
			films.GET("/:id", filmHandler.GetFilm)
			films.PUT("/:id", filmHandler.UpdateFilm)
			films.DELETE("/:id", filmHandler.DeleteFilm)
			films.POST("", filmHandler.CreateFilm)
		}
	}
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Engine.ServeHTTP(w, req)
}
