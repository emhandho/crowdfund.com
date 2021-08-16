package main

import (
	"log"
	"net/http"
	"strings"

	"crowdfund.com/auth"
	"crowdfund.com/campaign"
	"crowdfund.com/handler"
	"crowdfund.com/helper"
	"crowdfund.com/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/startup_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository	:= user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	userService		:= user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService		:= auth.NewJwtService()

	userHandler		:= handler.NewUserHandler(userService, authService) 
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router	:= gin.Default()
	router.Static("/images", "./images")
	api		:= router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/campaigns", campaignHandler.GetCampaigns)

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
