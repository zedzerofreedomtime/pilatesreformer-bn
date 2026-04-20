package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/config"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/http/handlers"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/http/middleware"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/service"
)

func NewRouter(cfg config.Config, handler *handlers.Handler, auth *service.AuthService) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSOrigins,
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/healthz", handler.Health)

	v1 := router.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/login", handler.Login)

		authenticated := authGroup.Group("")
		authenticated.Use(middleware.Authenticate(auth))
		authenticated.GET("/me", handler.Me)
		authenticated.POST("/logout", handler.Logout)

		catalog := v1.Group("/catalog")
		catalog.GET("/bootstrap", handler.CatalogBootstrap)
		catalog.GET("/equipment", handler.ListEquipment)
		catalog.GET("/trainers", handler.ListTrainers)
		catalog.GET("/rental-plans", handler.ListRentalPlans)
		catalog.GET("/trainer-service-plans", handler.ListTrainerServicePlans)
		catalog.GET("/home-content", handler.HomeContent)

		bookings := v1.Group("/bookings")
		bookings.POST("/quote", handler.QuoteBooking)
		bookings.POST("/inquiries", handler.CreateBookingInquiry)

		trainer := v1.Group("/trainer")
		trainer.Use(middleware.Authenticate(auth), middleware.RequireRole("trainer"))
		trainer.GET("/clients", handler.TrainerClients)

		admin := v1.Group("/admin")
		admin.Use(middleware.Authenticate(auth), middleware.RequireRole("admin"))
		admin.GET("/trainer-applications", handler.AdminListTrainerApplications)
		admin.POST("/trainer-applications/:id/approve", handler.AdminApproveTrainerApplication)
		admin.POST("/trainer-applications/:id/reject", handler.AdminRejectTrainerApplication)
		admin.POST("/trainers", handler.AdminCreateTrainer)
		admin.PUT("/trainers/:id", handler.AdminUpdateTrainer)
		admin.DELETE("/trainers/:id", handler.AdminDeleteTrainer)
		admin.POST("/equipment", handler.AdminCreateEquipment)
		admin.PUT("/equipment/:id", handler.AdminUpdateEquipment)
		admin.DELETE("/equipment/:id", handler.AdminDeleteEquipment)
		admin.GET("/home-content", handler.AdminGetHomeContent)
		admin.PUT("/home-content", handler.AdminSaveHomeContent)
		admin.GET("/bookings", handler.AdminListBookings)
	}

	return router
}
