package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/repository"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/service"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/store"
)

type Handler struct {
	repo    *repository.Repository
	auth    *service.AuthService
	catalog *service.CatalogService
	booking *service.BookingService
	admin   *service.AdminService
}

func New(repo *repository.Repository, auth *service.AuthService, catalog *service.CatalogService, booking *service.BookingService, admin *service.AdminService) *Handler {
	return &Handler{
		repo:    repo,
		auth:    auth,
		catalog: catalog,
		booking: booking,
		admin:   admin,
	}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) Register(c *gin.Context) {
	var payload service.RegisterInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := h.auth.Register(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *Handler) Login(c *gin.Context) {
	var payload service.LoginInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := h.auth.Login(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) Me(c *gin.Context) {
	user, ok := c.Get("authUser")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing auth context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) Logout(c *gin.Context) {
	token, _ := c.Get("sessionToken")
	tokenString, _ := token.(string)

	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing session token"})
		return
	}

	if err := h.auth.Logout(c.Request.Context(), tokenString); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func (h *Handler) CatalogBootstrap(c *gin.Context) {
	bootstrap, err := h.catalog.Bootstrap(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bootstrap)
}

func (h *Handler) ListEquipment(c *gin.Context) {
	items, err := h.repo.ListEquipment(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) ListTrainers(c *gin.Context) {
	items, err := h.repo.ListTrainers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) ListRentalPlans(c *gin.Context) {
	items, err := h.repo.ListRentalPlans(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) ListTrainerServicePlans(c *gin.Context) {
	items, err := h.repo.ListTrainerServicePlans(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) HomeContent(c *gin.Context) {
	content, err := h.repo.GetHomeContent(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, content)
}

func (h *Handler) QuoteBooking(c *gin.Context) {
	var payload service.QuoteInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := h.booking.Quote(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) CreateBookingInquiry(c *gin.Context) {
	var payload service.BookingInquiryInput
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	inquiry, err := h.booking.CreateInquiry(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, inquiry)
}

func (h *Handler) TrainerClients(c *gin.Context) {
	rawUser, ok := c.Get("authUser")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing auth context"})
		return
	}

	user := rawUser.(store.AuthenticatedUser)
	if user.TrainerID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "trainerId not attached to session"})
		return
	}

	items, err := h.repo.ListTrainerClients(c.Request.Context(), *user.TrainerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) AdminListTrainerApplications(c *gin.Context) {
	items, err := h.repo.ListPendingTrainerApplications(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) AdminApproveTrainerApplication(c *gin.Context) {
	trainer, err := h.admin.ApproveTrainerApplication(c.Request.Context(), c.Param("id"))
	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "trainer approved", "trainer": trainer})
}

func (h *Handler) AdminRejectTrainerApplication(c *gin.Context) {
	if err := h.admin.RejectTrainerApplication(c.Request.Context(), c.Param("id")); err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "trainer rejected"})
}

func (h *Handler) AdminCreateTrainer(c *gin.Context) {
	var payload store.Trainer
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	trainer, err := h.admin.CreateTrainer(c.Request.Context(), payload)
	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusCreated, trainer)
}

func (h *Handler) AdminUpdateTrainer(c *gin.Context) {
	var payload store.Trainer
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	payload.ID = c.Param("id")
	if err := h.admin.UpdateTrainer(c.Request.Context(), payload); err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "trainer updated"})
}

func (h *Handler) AdminDeleteTrainer(c *gin.Context) {
	if err := h.admin.DeleteTrainer(c.Request.Context(), c.Param("id")); err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "trainer deleted"})
}

func (h *Handler) AdminCreateEquipment(c *gin.Context) {
	var payload store.Equipment
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	equipment, err := h.admin.CreateEquipment(c.Request.Context(), payload)
	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusCreated, equipment)
}

func (h *Handler) AdminUpdateEquipment(c *gin.Context) {
	var payload store.Equipment
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	payload.ID = c.Param("id")
	if err := h.admin.UpdateEquipment(c.Request.Context(), payload); err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "equipment updated"})
}

func (h *Handler) AdminDeleteEquipment(c *gin.Context) {
	if err := h.admin.DeleteEquipment(c.Request.Context(), c.Param("id")); err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "equipment deleted"})
}

func (h *Handler) AdminGetHomeContent(c *gin.Context) {
	content, err := h.repo.GetHomeContent(c.Request.Context())
	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, content)
}

func (h *Handler) AdminSaveHomeContent(c *gin.Context) {
	var payload map[string]any
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.admin.SaveHomeContent(c.Request.Context(), payload); err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "home content saved"})
}

func (h *Handler) AdminListBookings(c *gin.Context) {
	items, err := h.repo.ListBookingInquiries(c.Request.Context())
	if err != nil {
		h.writeError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *Handler) writeError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, repository.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}
