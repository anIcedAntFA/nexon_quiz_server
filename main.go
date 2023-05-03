package main

import (
	"log"
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	"nexon_quiz/middleware"
	"os"

	answertransport "nexon_quiz/modules/answer/transport"
	categorytransport "nexon_quiz/modules/category/transport"
	difficultytransport "nexon_quiz/modules/difficulty/transport"
	questiontransport "nexon_quiz/modules/question/transport"
	typetransport "nexon_quiz/modules/type/transport"
	usertransport "nexon_quiz/modules/user/transport"
	userroletransport "nexon_quiz/modules/userrole/transport"
	usersettingtransport "nexon_quiz/modules/usersetting/transport"

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

	// User Role API only by Root Admin
	role := v1.Group(
		"/roles",
		middleware.RequiredAuthorization(appContext),
		middleware.RequiredRole(appContext, common.RootAdminRole))
	role.POST("/new", userroletransport.HandleCreateNewUserRole(appContext))
	role.GET("/:id", userroletransport.HandleGetUserRole(appContext))

	// Auth API
	auth := v1.Group("/auth")
	auth.POST("/register", usertransport.HandleRegisterUser(appContext))
	auth.POST("/authenticate", usertransport.HandleLoginUser(appContext))

	// Question Type API
	types := v1.Group("/types", middleware.RequiredAuthorization(appContext))
	types.POST("/new",
		middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
		typetransport.HandleCreateNewType(appContext),
	)
	types.GET("", typetransport.HandleGetTypeList(appContext))
	types.GET("/:id", typetransport.HandleGetTypeById(appContext))
	types.PATCH(
		"/:id",
		middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
		typetransport.HandleUpdateTypeById(appContext),
	)
	types.DELETE(
		"/:id",
		middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
		typetransport.HandleDeleteTypeById(appContext),
	)

	// Question Difficulty API
	difficulties := v1.Group("/difficulties", middleware.RequiredAuthorization(appContext))
	difficulties.POST("",
		middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
		difficultytransport.HandleCreateNewDifficulty(appContext),
	)
	difficulties.GET("", difficultytransport.HandleGetDifficultyList(appContext))
	difficulties.GET("/:id", difficultytransport.HandleGetDifficultyById(appContext))

	// Question Category API
	categories := v1.Group("/categoriess", middleware.RequiredAuthorization(appContext))
	categories.POST(
		"/new",
		middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
		categorytransport.HandleCreateNewCategory(appContext),
	)
	categories.POST(
		"",
		middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
		categorytransport.HandleCreateCategoryList(appContext),
	)
	categories.GET("/:id", categorytransport.HandleGetCategory(appContext))
	categories.GET("", categorytransport.HandleGetCategoryList(appContext))

	// Question API
	questions := v1.Group(
		"/questions",
		middleware.RequiredAuthorization(appContext),
	)
	questions.POST(
		"/new",
		middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
		questiontransport.HandleCreateNewQuestion(appContext),
	)
	questions.POST(
		"",
		middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
		questiontransport.HandleCreateQuestionList(appContext),
	)
	questions.GET("", questiontransport.HandleGetQuestionList(appContext))
	questions.GET("/:id", questiontransport.HandleGetQuestionById(appContext))

	// Answer API
	answers := v1.Group(
		"/answers",
		middleware.RequiredAuthorization(appContext),
	)
	answers.POST("/new",
		middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
		answertransport.HandleCreateNewAnswer(appContext),
	)
	answers.POST("",
		middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
		answertransport.HandleCreateNewAnswer(appContext),
	)

	// User Setting API
	userSetting := v1.Group("user/setting", middleware.RequiredAuthorization(appContext))
	userSetting.POST("/new", usersettingtransport.HandleCreateNewUserSetting(appContext))
	// userSetting.GET("", typetransport.HandleGetTypeList(appContext))
	// userSetting.GET("/:id", typetransport.HandleGetTypeById(appContext))
	// userSetting.PATCH(
	// 	"/:id",
	// 	middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
	// 	typetransport.HandleUpdateTypeById(appContext),
	// )
	// userSetting.DELETE(
	// 	"/:id",
	// 	middleware.RequiredRole(appContext, common.RootAdminRole, common.AdminRole),
	// 	typetransport.HandleDeleteTypeById(appContext),
	// )

	router.Run()

	return nil
}
