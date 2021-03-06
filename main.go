package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"crowdfund.com/auth"
	"crowdfund.com/campaign"
	"crowdfund.com/handler"
	"crowdfund.com/helper"
	"crowdfund.com/payment"
	"crowdfund.com/transaction"
	"github.com/joho/godotenv"

	// "crowdfund.com/transaction"
	"crowdfund.com/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connection() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Unable to load .env file %v", err)
	}

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := connection()
	if err != nil {
		log.Fatal(err.Error())
	}

	// user repo and service layer
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	//campaign repo and service layer
	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)

	// transaction repo and service layer
	transactionRepository := transaction.NewRepository(db)
	paymentService := payment.NewService() // create payment service
	transacrtionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	// jwt service generator object
	key := os.Getenv("APP_SECRET_KEY")
	secret_key := []byte(key)
	authService := auth.NewJwtService(secret_key)

	// All the Handler for the Entity layer
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transacrtionService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.POST("/avatars/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadCampaignImages)

	api.GET("/campaigns/:id/transactions", authMiddleware(authService, userService), transactionHandler.GetCampaignTransaction)
	api.GET("/transactions", authMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	router.Run()
}

// Authentication Middleware
// get value header Authorization: Bearer tokentokentoken
// form header Athorization, get only the token value
// validate the token
// get user_id
// get user from db by user_id using service
// set context with user
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayHeader := strings.Split(authHeader, " ")
		if len(arrayHeader) == 2 {
			tokenString = arrayHeader[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}

}
