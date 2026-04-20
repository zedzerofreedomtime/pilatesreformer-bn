package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/repository"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/store"
)

type AdminService struct {
	repo    *repository.Repository
	catalog *CatalogService
}

func NewAdminService(repo *repository.Repository, catalog *CatalogService) *AdminService {
	return &AdminService{
		repo:    repo,
		catalog: catalog,
	}
}

func (s *AdminService) CreateEquipment(ctx context.Context, equipment store.Equipment) (store.Equipment, error) {
	if equipment.ID == "" {
		equipment.ID = fmt.Sprintf("equipment-%s", uuid.NewString()[:8])
	}

	if err := s.repo.CreateEquipment(ctx, equipment); err != nil {
		return store.Equipment{}, err
	}
	s.catalog.Invalidate(ctx)
	return equipment, nil
}

func (s *AdminService) UpdateEquipment(ctx context.Context, equipment store.Equipment) error {
	if err := s.repo.UpdateEquipment(ctx, equipment); err != nil {
		return err
	}
	s.catalog.Invalidate(ctx)
	return nil
}

func (s *AdminService) DeleteEquipment(ctx context.Context, id string) error {
	if err := s.repo.DeleteEquipment(ctx, id); err != nil {
		return err
	}
	s.catalog.Invalidate(ctx)
	return nil
}

func (s *AdminService) CreateTrainer(ctx context.Context, trainer store.Trainer) (store.Trainer, error) {
	if trainer.ID == "" {
		trainer.ID = fmt.Sprintf("coach-%s", uuid.NewString()[:8])
	}
	if len(trainer.WeeklySchedule) == 0 {
		trainer.WeeklySchedule = store.DefaultWeeklySchedule()
	}
	trainer.AvailableSlots, trainer.BookedSlots = store.SummarizeSchedule(trainer.WeeklySchedule)
	if trainer.ScheduleWindow == "" {
		trainer.ScheduleWindow = "Sunday - Saturday 08:00 - 17:00"
	}

	if err := s.repo.CreateTrainer(ctx, trainer); err != nil {
		return store.Trainer{}, err
	}
	s.catalog.Invalidate(ctx)
	return trainer, nil
}

func (s *AdminService) UpdateTrainer(ctx context.Context, trainer store.Trainer) error {
	if len(trainer.WeeklySchedule) == 0 {
		trainer.WeeklySchedule = store.DefaultWeeklySchedule()
	}
	trainer.AvailableSlots, trainer.BookedSlots = store.SummarizeSchedule(trainer.WeeklySchedule)
	if trainer.ScheduleWindow == "" {
		trainer.ScheduleWindow = "Sunday - Saturday 08:00 - 17:00"
	}

	if err := s.repo.UpdateTrainer(ctx, trainer); err != nil {
		return err
	}
	s.catalog.Invalidate(ctx)
	return nil
}

func (s *AdminService) DeleteTrainer(ctx context.Context, id string) error {
	if err := s.repo.DeleteTrainer(ctx, id); err != nil {
		return err
	}
	s.catalog.Invalidate(ctx)
	return nil
}

func (s *AdminService) ApproveTrainerApplication(ctx context.Context, applicationID string) (store.Trainer, error) {
	schedule := store.DefaultWeeklySchedule()
	availableSlots, bookedSlots := store.SummarizeSchedule(schedule)

	trainer := store.Trainer{
		ID:             fmt.Sprintf("coach-%s", uuid.NewString()[:8]),
		Name:           "New Coach",
		Image:          "/images/trainer-pim.svg",
		Specialty:      "Personalized coaching",
		SessionRate:    1800,
		Availability:   "Onsite / Online",
		Summary:        "Profile created from approved trainer application",
		ScheduleWindow: "Sunday - Saturday 08:00 - 17:00",
		AvailableSlots: availableSlots,
		BookedSlots:    bookedSlots,
		MachineFocus:   []string{"Pilates Reformer"},
		ExerciseFocus:  []string{"Private training"},
		WeeklySchedule: schedule,
	}

	approvedTrainer, err := s.repo.ApproveTrainerApplication(ctx, applicationID, trainer)
	if err != nil {
		return store.Trainer{}, err
	}

	s.catalog.Invalidate(ctx)
	return approvedTrainer, nil
}

func (s *AdminService) RejectTrainerApplication(ctx context.Context, applicationID string) error {
	if err := s.repo.RejectTrainerApplication(ctx, applicationID); err != nil {
		return err
	}
	return nil
}

func (s *AdminService) SaveHomeContent(ctx context.Context, content map[string]any) error {
	if err := s.repo.SaveHomeContent(ctx, content); err != nil {
		return err
	}
	s.catalog.Invalidate(ctx)
	return nil
}
