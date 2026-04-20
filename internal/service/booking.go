package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/repository"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/store"
)

type BookingService struct {
	repo *repository.Repository
}

func NewBookingService(repo *repository.Repository) *BookingService {
	return &BookingService{repo: repo}
}

type QuoteInput struct {
	ModeID               string `json:"modeId"`
	EquipmentID          string `json:"equipmentId"`
	RentalPlanID         string `json:"rentalPlanId"`
	TrainerID            string `json:"trainerId"`
	TrainerServicePlanID string `json:"trainerServicePlanId"`
}

type QuoteResult struct {
	RentalSubtotal        int `json:"rentalSubtotal"`
	InstallFee            int `json:"installFee"`
	BundleSessions        int `json:"bundleSessions"`
	BundleTrainerSubtotal int `json:"bundleTrainerSubtotal"`
	BundleGrandTotal      int `json:"bundleGrandTotal"`
	EquipmentOnlyTotal    int `json:"equipmentOnlyTotal"`
	TrainerOnlyTotal      int `json:"trainerOnlyTotal"`
}

type BookingInquiryInput struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	ModeID               string `json:"modeId"`
	EquipmentID          string `json:"equipmentId"`
	RentalPlanID         string `json:"rentalPlanId"`
	TrainerID            string `json:"trainerId"`
	TrainerServicePlanID string `json:"trainerServicePlanId"`
	Notes                string `json:"notes"`
}

func (s *BookingService) Quote(ctx context.Context, input QuoteInput) (QuoteResult, error) {
	equipmentCatalog, err := s.repo.ListEquipment(ctx)
	if err != nil {
		return QuoteResult{}, err
	}

	rentalPlanCatalog, err := s.repo.ListRentalPlans(ctx)
	if err != nil {
		return QuoteResult{}, err
	}

	trainerCatalog, err := s.repo.ListTrainers(ctx)
	if err != nil {
		return QuoteResult{}, err
	}

	trainerServicePlans, err := s.repo.ListTrainerServicePlans(ctx)
	if err != nil {
		return QuoteResult{}, err
	}

	equipment, ok := findByID(equipmentCatalog, input.EquipmentID, func(item store.Equipment) string { return item.ID })
	if !ok && input.ModeID != "trainer-only" {
		return QuoteResult{}, errors.New("equipment not found")
	}

	rentalPlan, ok := findByID(rentalPlanCatalog, input.RentalPlanID, func(item store.RentalPlan) string { return item.ID })
	if !ok && input.ModeID != "trainer-only" {
		return QuoteResult{}, errors.New("rental plan not found")
	}

	trainer, ok := findByID(trainerCatalog, input.TrainerID, func(item store.Trainer) string { return item.ID })
	if !ok && input.ModeID != "equipment-only" {
		return QuoteResult{}, errors.New("trainer not found")
	}

	trainerServicePlan, ok := findByID(trainerServicePlans, input.TrainerServicePlanID, func(item store.TrainerServicePlan) string { return item.ID })
	if !ok && input.ModeID == "trainer-only" {
		return QuoteResult{}, errors.New("trainer service plan not found")
	}

	rentalSubtotal := int(float64(equipment.MonthlyRate*rentalPlan.Months) * rentalPlan.Discount)
	installFee := 1500
	if rentalPlan.Months >= 3 {
		installFee = 0
	}

	bundleSessions := rentalPlan.OptionalSessions
	if equipment.TrainerMode == "required" {
		bundleSessions = rentalPlan.RequiredSessions
	}

	bundleTrainerSubtotal := bundleSessions * trainer.SessionRate
	trainerOnlyTotal := int(float64(trainer.SessionRate*trainerServicePlan.Sessions) * trainerServicePlan.Discount)

	return QuoteResult{
		RentalSubtotal:        rentalSubtotal,
		InstallFee:            installFee,
		BundleSessions:        bundleSessions,
		BundleTrainerSubtotal: bundleTrainerSubtotal,
		BundleGrandTotal:      rentalSubtotal + installFee + bundleTrainerSubtotal,
		EquipmentOnlyTotal:    rentalSubtotal + installFee,
		TrainerOnlyTotal:      trainerOnlyTotal,
	}, nil
}

func (s *BookingService) CreateInquiry(ctx context.Context, input BookingInquiryInput) (store.BookingInquiry, error) {
	quote, err := s.Quote(ctx, QuoteInput{
		ModeID:               input.ModeID,
		EquipmentID:          input.EquipmentID,
		RentalPlanID:         input.RentalPlanID,
		TrainerID:            input.TrainerID,
		TrainerServicePlanID: input.TrainerServicePlanID,
	})
	if err != nil {
		return store.BookingInquiry{}, err
	}

	inquiry := store.BookingInquiry{
		ID:         fmt.Sprintf("booking-%s", uuid.NewString()[:8]),
		Name:       input.Name,
		Email:      input.Email,
		Phone:      input.Phone,
		ModeID:     input.ModeID,
		GrandTotal: resolveGrandTotal(input.ModeID, quote),
		Notes:      input.Notes,
		Status:     "new",
	}

	if input.EquipmentID != "" {
		inquiry.EquipmentID = &input.EquipmentID
	}
	if input.RentalPlanID != "" {
		inquiry.RentalPlanID = &input.RentalPlanID
	}
	if input.TrainerID != "" {
		inquiry.TrainerID = &input.TrainerID
	}
	if input.TrainerServicePlanID != "" {
		inquiry.TrainerServicePlanID = &input.TrainerServicePlanID
	}

	if err := s.repo.CreateBookingInquiry(ctx, inquiry); err != nil {
		return store.BookingInquiry{}, err
	}

	return inquiry, nil
}

func resolveGrandTotal(modeID string, quote QuoteResult) int {
	switch modeID {
	case "equipment-only":
		return quote.EquipmentOnlyTotal
	case "trainer-only":
		return quote.TrainerOnlyTotal
	default:
		return quote.BundleGrandTotal
	}
}

func findByID[T any](items []T, id string, pickID func(T) string) (T, bool) {
	var zero T
	for _, item := range items {
		if pickID(item) == id {
			return item, true
		}
	}
	return zero, false
}
