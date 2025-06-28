package api

import (
	"GOLANG_CLEAN_WEB_API/src/api/middlewares"
	"GOLANG_CLEAN_WEB_API/src/api/routers"
	validations "GOLANG_CLEAN_WEB_API/src/api/validations"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/docs"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()


	RegisterSwagger(r, cfg)

	//r.Use(middlewares.Cors(cfg))
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())

	log.Println("InitServer - Middleware LimitByRequest registered.")

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}

	log.Println("InitServer - Routes are being initialized.")

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		heath := v1.Group("/health")
		test_router := v1.Group("/test")
		users := v1.Group("/users")


		routers.Health(heath)
		routers.TestRouter(test_router)
		routers.User(users, cfg)
	}

	log.Println("InitServer - API groups registered.")

	

	log.Println("InitServer - Running server on port:", cfg.Server.Port)
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	log.Println("[Swagger Setup] RegisterSwagger executed")

	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

//server := &https.Server{

//handler : r ,
//Addr : (fmt.Sprintf(":%s",cfg.Server.Port))
//ReadTimeout : time.Second * 10

//server.ListenAndServe()
