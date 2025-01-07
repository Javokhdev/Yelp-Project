package app

import (
	"log"

	_ "github.com/Javokhdev/Yelp-Project/api/docs" // Import Swagger docs
	"github.com/Javokhdev/Yelp-Project/config"
	"github.com/Javokhdev/Yelp-Project/internal/controller/handlers"
	"github.com/Javokhdev/Yelp-Project/internal/core/repositories"
	"github.com/Javokhdev/Yelp-Project/internal/core/services"
	"github.com/Javokhdev/Yelp-Project/pkg/db"
	"github.com/Javokhdev/Yelp-Project/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Yalp API
// @version 1.0
// @description API documentation for Yalp
// @BasePath /api/v1
func Run(cfg *config.Config) {
	infoLog, errLog := logger.InitLogger()

	// Postgres Connection
	db, err := db.Connect(cfg)
	if err != nil {
		errLog.Println("Can't connect to database, details:", err, log.Ldate|log.Ltime|log.Lshortfile)
	}
	defer db.Close()
	infoLog.Println("Connected to Postgres")

	userRepo := repositories.NewUserRepo(db)
	businessRepo := repositories.NewBusinessRepo(db)
	reviewRepo := repositories.NewReviewRepo(db)
	category := repositories.NewCategoryRepo(db)

	userSvc := services.NewAuthService(*userRepo)
	businessSvc := services.NewBusinessService(*businessRepo)
	reviewSvc := services.NewReviewAndRatingService(*reviewRepo)
	categorySvc := services.NewCategoryService(*category)

	authHandler := handlers.NewAuthHandler(userSvc)
	businessHandler := handlers.NewBusinessHandler(businessSvc)
	reviewHandler := handlers.NewReviewHandler(reviewSvc)
	categoryHandler := handlers.NewCategoryHandler(categorySvc)

	router := gin.Default()
	router.Use(cors.Default())

	// Swagger endpoint
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/auth/register", authHandler.RegisterUser)
		v1.GET("/auth/:userID", authHandler.GetUserByID)
		v1.GET("/auth", authHandler.GetAllUsers)
		v1.DELETE("/auth/:userID", authHandler.DeleteUser)
		v1.POST("/business", businessHandler.CreateBusiness)
		v1.GET("/business", businessHandler.GetAllBusinesses)
		v1.GET("/business/:businessID", businessHandler.GetBusinessByID)
		v1.POST("/business/:businessID/reviews", reviewHandler.CreateReview)
		v1.GET("/business/:businessID/reviews/:reviewID", reviewHandler.GetReviewByID)
		v1.PUT("/business/:businessID/reviews/:reviewID", reviewHandler.UpdateReview)
		v1.DELETE("/business/:businessID/reviews/:reviewID", reviewHandler.DeleteReview)
		v1.GET("/business/:businessID/reviews", reviewHandler.GetAllReviews)
		v1.POST("/categories", categoryHandler.CreateCategory)
		v1.GET("/categories", categoryHandler.GetAllCategories)
		v1.GET("/categories/:categoryID", categoryHandler.GetCategoryByID)
		v1.DELETE("/categories/:categoryID", categoryHandler.DeleteCategory)
	}

	router.Run(":8080")
}
