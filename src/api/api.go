package api

import (
	"GOLANG_CLEAN_WEB_API/docs"
	"GOLANG_CLEAN_WEB_API/src/api/middlewares"
	"GOLANG_CLEAN_WEB_API/src/api/routers"
	validations "GOLANG_CLEAN_WEB_API/src/api/validations"
	"GOLANG_CLEAN_WEB_API/src/config"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/swag/example/basic/docs"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	RegisterSwagger(r, cfg)

	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(
		gin.Logger(),
		gin.RecoveryWithWriter(gin.DefaultWriter, func(c *gin.Context, err interface{}) {
			middlewares.ErrorHandler(c, err)//خط بیست و هشت و بیست و نه جدید اضافه شده اند
		}),
		middlewares.LimitByRequest(),
	)

	log.Println("InitServer - Middleware LimitByRequest registered.")

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
	}

	log.Println("InitServer - Routes are being initialized.")

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{   //Test
		heath := v1.Group("/health")
		test_router := v1.Group("/test")


        //User
		users := v1.Group("/users")

		
		//Base
		Countries :=v1.Group("/countries",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
		Cities :=v1.Group("/Cities",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
		files :=v1.Group("/files",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
		companies := v1.Group("/companies",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
		colors := v1.Group("/colors",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
        Years := v1.Group("/years",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))


        //property
		properties :=v1.Group("/properties",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
		propertyCategories :=v1.Group("/property-categories",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
        //Car
		carTypes := v1.Group("/car-types",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
		gearboxes := v1.Group("/gearboxes",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
        carModels := v1.Group("/car-models",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
        carModelColors := v1.Group("/car-model-colors",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
        carModelYears := v1.Group("/car-model-years",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
        carModelPriceHistories := v1.Group("/car-model-price-histories",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))
        carModelImage := v1.Group("/car-model-images",middlewares.Authentication(cfg),middlewares.Authorization([]string{"admin"}))


        //Test
		routers.Health(heath)
		routers.TestRouter(test_router)


        //User
		routers.User(users, cfg)


        //Base
		routers.Country(Countries, cfg)
		routers.City(Cities , cfg)
		routers.File(files,cfg)
		routers.Company(companies,cfg)
		routers.Color(colors,cfg)
		routers.Year(Years,cfg)


		//property
		routers.Property(properties,cfg)
		routers.PropertyCategory(propertyCategories,cfg)

		//Car
		routers.CarType(carTypes,cfg)
		routers.Gearbox(gearboxes,cfg)
		routers.CarModel(carModels,cfg)
		routers.CarModelColor(carModelColors,cfg)
		routers.CarModelYear(carModelYears,cfg)
		routers.CarModelPriceHistory(carModelPriceHistories,cfg)
        routers.CarModelImage(carModelImage,cfg)

		r.Static("/static" , "./uploads")
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
