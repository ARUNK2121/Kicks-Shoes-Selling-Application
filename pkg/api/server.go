package http

import (
	"log"

	"github.com/gin-gonic/gin"

	handler "kicks/pkg/api/handler"
	"kicks/pkg/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	categoryHandler *handler.CategoryHandler,
	inventoryHandler *handler.InventoryHandler,
	otpHandler *handler.OtpHandler,
	orderHandler *handler.OrderHandler,
	cartHandler *handler.CartHandler,
	couponHandler *handler.CouponHandler,
	paymentHandler *handler.PaymentHandler,
	offerhandler *handler.OfferHandler,
	wishlistHandler *handler.WishlistHandler) *ServerHTTP {

	engine := gin.New()

	engine.LoadHTMLGlob("templates/*.html")

	// Use logger from Gin
	engine.Use(gin.Logger())

	//Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	engine.GET("/validate-token", adminHandler.ValidateRefreshTokenAndCreateNewAccess)

	routes.UserRoutes(engine.Group("/users"), userHandler, otpHandler, inventoryHandler, orderHandler, cartHandler, paymentHandler, wishlistHandler, categoryHandler, couponHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, inventoryHandler, userHandler, categoryHandler, orderHandler, couponHandler, offerhandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	err := sh.engine.Run(":3000")
	if err != nil {
		log.Fatal("gin engine couldn't start")
	}
}
