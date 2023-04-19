package main

import (
	"log"
	"net/http"
	"nexon_quiz/components/appctx"
	"nexon_quiz/middleware"
	"os"

	answertransport "nexon_quiz/modules/answer/transport"
	questiontransport "nexon_quiz/modules/question/transport"
	usertransport "nexon_quiz/modules/user/transport"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading.env file:", err)
	}

	dsn := os.Getenv("MYSQL_CONNECTION_ROOT")

	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	db = db.Debug()

	if err := runService(db, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(
	db *gorm.DB,
	secretKey string,
) error {
	appContext := appctx.NewAppContext(db, secretKey)

	router := gin.Default()

	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	// To be able to send tokens to the server.
	// config.AllowAllOrigins = true
	// OPTIONS method for ReactJS
	// config.AddAllowMethods("OPTIONS")
	config.AddAllowHeaders("authorization")
	router.Use(cors.New(config))

	router.Use(middleware.Recover(appContext))

	v1 := router.Group("/v1")

	v1.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello 500 ae",
		})
	})

	auth := v1.Group("/auth")
	auth.POST("/register", usertransport.HandleRegisterUser(appContext))
	auth.POST("/authenticate", usertransport.HandleLoginUser(appContext))

	question := v1.Group(
		"/question",
		middleware.RequiredAuthorization(appContext),
		middleware.RequiredRole(appContext, "admin"))
	question.POST("/new", questiontransport.HandleCreateNewQuestion(appContext))

	answer := v1.Group(
		"/answer",
		middleware.RequiredAuthorization(appContext),
		middleware.RequiredRole(appContext, "admin"))
	answer.POST("/new", answertransport.HandleCreateNewAnswer(appContext))

	router.Run()

	return nil
}
