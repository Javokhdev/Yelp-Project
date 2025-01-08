// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	rediscache "github.com/golanguzb70/redis-cache"
	"github.com/Javokhdev/Yelp-Project/config"
	// _ "github.com/Javokhdev/Yelp-Project/docs"
	"github.com/Javokhdev/Yelp-Project/internal/controller/http/v1/handlers"
	"github.com/Javokhdev/Yelp-Project/internal/usecase"
	"github.com/Javokhdev/Yelp-Project/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description This is a sample server Go Clean Template server.
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(engine *gin.Engine, l *logger.Logger, config *config.Config, useCase *usecase.UseCase, redis rediscache.RedisCache) {
	// Options
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	handlerV1 := handlers.NewHandler(l, config, useCase, redis)

	// Initialize Casbin enforcer
	e := casbin.NewEnforcer("config/rbac.conf", "config/policy.csv")
	engine.Use(handlerV1.AuthMiddleware(e))

	// Swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// K8s probe
	engine.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routes
	v1 := engine.Group("/v1")

	user := v1.Group("/user")
	{
		user.POST("/", handlerV1.CreateUser)
		user.GET("/list", handlerV1.GetUsers)
		user.GET("/:id", handlerV1.GetUser)
		user.PUT("/", handlerV1.UpdateUser)
		user.DELETE("/:id", handlerV1.DeleteUser)
	}

	session := v1.Group("/session")
	{
		session.GET("/list", handlerV1.GetSessions)
		session.GET("/:id", handlerV1.GetSession)
		session.PUT("/", handlerV1.UpdateSession)
		session.DELETE("/:id", handlerV1.DeleteSession)
	}

	auth := v1.Group("/auth")
	{
		auth.POST("/logout", handlerV1.Logout)
		auth.POST("/register", handlerV1.Register)
		auth.POST("/verify-email", handlerV1.VerifyEmail)
		auth.POST("/login", handlerV1.Login)
	}

	business := v1.Group("/business")
	{
		business.POST("/", handlerV1.CreateBusiness)
		business.GET("/list", handlerV1.GetAllBusinesses)
		business.GET("/:id", handlerV1.GetBusinessByID)
		business.PUT("/", handlerV1.UpdateBusiness)
		business.DELETE("/:id", handlerV1.DeleteBusiness)
	}

	review := v1.Group("/review")
	{
		review.POST("/", handlerV1.CreateReview)
		review.GET("/list", handlerV1.GetAllReviews)
		review.GET("/:id", handlerV1.GetReviewByID)
		review.PUT("/", handlerV1.UpdateReview)
		review.DELETE("/:id", handlerV1.DeleteReview)
		review.GET("/rating/:id", handlerV1.GetRatingByBusinessID)
	}

	category := v1.Group("/category")
	{
		category.POST("/", handlerV1.CreateCategory)
		category.GET("/list", handlerV1.GetAllCategories)
		category.GET("/:id", handlerV1.GetCategoryByID)
		category.PUT("/", handlerV1.UpdateCategory)
		category.DELETE("/:id", handlerV1.DeleteCategory)
	}

	
}
