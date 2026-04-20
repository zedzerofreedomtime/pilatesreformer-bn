package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/repository"
	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/store"
)

type CatalogService struct {
	repo     *repository.Repository
	redis    *redis.Client
	cacheTTL time.Duration
}

func NewCatalogService(repo *repository.Repository, redis *redis.Client, cacheTTL time.Duration) *CatalogService {
	return &CatalogService{
		repo:     repo,
		redis:    redis,
		cacheTTL: cacheTTL,
	}
}

func (s *CatalogService) Bootstrap(ctx context.Context) (store.CatalogBootstrap, error) {
	const cacheKey = "catalog:bootstrap"

	if raw, err := s.redis.Get(ctx, cacheKey).Bytes(); err == nil {
		var cached store.CatalogBootstrap
		if json.Unmarshal(raw, &cached) == nil {
			return cached, nil
		}
	}

	equipmentCatalog, err := s.repo.ListEquipment(ctx)
	if err != nil {
		return store.CatalogBootstrap{}, err
	}

	trainerCatalog, err := s.repo.ListTrainers(ctx)
	if err != nil {
		return store.CatalogBootstrap{}, err
	}

	rentalPlanCatalog, err := s.repo.ListRentalPlans(ctx)
	if err != nil {
		return store.CatalogBootstrap{}, err
	}

	trainerServicePlans, err := s.repo.ListTrainerServicePlans(ctx)
	if err != nil {
		return store.CatalogBootstrap{}, err
	}

	homePageContent, err := s.repo.GetHomeContent(ctx)
	if err != nil {
		return store.CatalogBootstrap{}, err
	}

	bootstrap := store.CatalogBootstrap{
		EquipmentCatalog:    equipmentCatalog,
		TrainerCatalog:      trainerCatalog,
		RentalPlanCatalog:   rentalPlanCatalog,
		TrainerServicePlans: trainerServicePlans,
		HomePageContent:     homePageContent,
		HomeStats:           buildHomeStats(homePageContent, len(trainerCatalog), len(equipmentCatalog)),
		TopRentedCards:      buildTopRentedCards(homePageContent, equipmentCatalog),
	}

	if payload, err := json.Marshal(bootstrap); err == nil {
		_ = s.redis.Set(ctx, cacheKey, payload, s.cacheTTL).Err()
	}

	return bootstrap, nil
}

func (s *CatalogService) Invalidate(ctx context.Context) {
	_ = s.redis.Del(ctx, "catalog:bootstrap").Err()
}

func buildHomeStats(homePageContent map[string]any, trainerCount, equipmentCount int) []map[string]any {
	rawStats, ok := homePageContent["stats"].([]any)
	if !ok {
		return []map[string]any{
			{"value": trainerCount, "label": "trainers"},
			{"value": equipmentCount, "label": "equipment"},
		}
	}

	result := make([]map[string]any, 0, len(rawStats))
	for index, rawItem := range rawStats {
		item, ok := rawItem.(map[string]any)
		if !ok {
			continue
		}

		copyItem := map[string]any{}
		for key, value := range item {
			copyItem[key] = value
		}

		if index == 0 {
			copyItem["value"] = trainerCount
		}
		if index == 1 {
			copyItem["value"] = equipmentCount
		}

		result = append(result, copyItem)
	}

	return result
}

func buildTopRentedCards(homePageContent map[string]any, equipmentCatalog []store.Equipment) []map[string]any {
	topRentals, ok := homePageContent["topRentals"].(map[string]any)
	if !ok {
		return []map[string]any{}
	}

	rawItems, ok := topRentals["items"].([]any)
	if !ok {
		return []map[string]any{}
	}

	equipmentByID := map[string]store.Equipment{}
	for _, equipment := range equipmentCatalog {
		equipmentByID[equipment.ID] = equipment
	}

	result := make([]map[string]any, 0, len(rawItems))
	for _, rawItem := range rawItems {
		item, ok := rawItem.(map[string]any)
		if !ok {
			continue
		}

		equipmentID, _ := item["equipmentId"].(string)
		equipment, ok := equipmentByID[equipmentID]
		if !ok {
			continue
		}

		card := map[string]any{}
		for key, value := range item {
			card[key] = value
		}
		card["equipment"] = equipment
		result = append(result, card)
	}

	return result
}
