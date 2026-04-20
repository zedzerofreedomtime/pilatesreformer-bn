package store

import "time"

type AuthenticatedUser struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	RoleID    string  `json:"roleId"`
	RoleLabel string  `json:"roleLabel"`
	TrainerID *string `json:"trainerId,omitempty"`
}

type User struct {
	ID           string
	Name         string
	Email        string
	Phone        string
	RoleID       string
	PasswordHash string
	TrainerID    *string
	CreatedAt    time.Time
}

type Equipment struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	Badge       string   `json:"badge"`
	MonthlyRate int      `json:"monthlyRate"`
	TrainerMode string   `json:"trainerMode"`
	Summary     string   `json:"summary"`
	IdealFor    string   `json:"idealFor"`
	Footprint   string   `json:"footprint"`
	Features    []string `json:"features"`
}

type Slot struct {
	Key    string `json:"key"`
	Label  string `json:"label"`
	Status string `json:"status"`
}

type WeeklyScheduleDay struct {
	ID             string `json:"id"`
	Label          string `json:"label"`
	ShortLabel     string `json:"shortLabel"`
	AvailableCount int    `json:"availableCount"`
	BookedCount    int    `json:"bookedCount"`
	Slots          []Slot `json:"slots"`
}

type Trainer struct {
	ID             string              `json:"id"`
	Name           string              `json:"name"`
	Image          string              `json:"image"`
	Specialty      string              `json:"specialty"`
	SessionRate    int                 `json:"sessionRate"`
	Availability   string              `json:"availability"`
	Summary        string              `json:"summary"`
	ScheduleWindow string              `json:"scheduleWindow"`
	AvailableSlots int                 `json:"availableSlots"`
	BookedSlots    int                 `json:"bookedSlots"`
	MachineFocus   []string            `json:"machineFocus"`
	ExerciseFocus  []string            `json:"exerciseFocus"`
	WeeklySchedule []WeeklyScheduleDay `json:"weeklySchedule"`
}

type TrainerClient struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	EquipmentName string `json:"equipmentName"`
	PlanName      string `json:"planName"`
	NextSession   string `json:"nextSession"`
	Contact       string `json:"contact"`
	Status        string `json:"status"`
}

type TrainerApplication struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Specialty    string    `json:"specialty"`
	MachineFocus []string  `json:"machineFocus"`
	Status       string    `json:"status"`
	SubmittedAt  time.Time `json:"submittedAt"`
	PasswordHash string    `json:"-"`
}

type RentalPlan struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Months           int     `json:"months"`
	Discount         float64 `json:"discount"`
	OptionalSessions int     `json:"optionalSessions"`
	RequiredSessions int     `json:"requiredSessions"`
	Note             string  `json:"note"`
}

type TrainerServicePlan struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Sessions int     `json:"sessions"`
	Discount float64 `json:"discount"`
	Note     string  `json:"note"`
}

type BookingInquiry struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Email                string    `json:"email"`
	Phone                string    `json:"phone"`
	ModeID               string    `json:"modeId"`
	EquipmentID          *string   `json:"equipmentId,omitempty"`
	RentalPlanID         *string   `json:"rentalPlanId,omitempty"`
	TrainerID            *string   `json:"trainerId,omitempty"`
	TrainerServicePlanID *string   `json:"trainerServicePlanId,omitempty"`
	GrandTotal           int       `json:"grandTotal"`
	Notes                string    `json:"notes,omitempty"`
	Status               string    `json:"status"`
	CreatedAt            time.Time `json:"createdAt"`
}

type CatalogBootstrap struct {
	EquipmentCatalog    []Equipment          `json:"equipmentCatalog"`
	TrainerCatalog      []Trainer            `json:"trainerCatalog"`
	RentalPlanCatalog   []RentalPlan         `json:"rentalPlanCatalog"`
	TrainerServicePlans []TrainerServicePlan `json:"trainerServicePlans"`
	HomePageContent     map[string]any       `json:"homePageContent"`
	HomeStats           []map[string]any     `json:"homeStats"`
	TopRentedCards      []map[string]any     `json:"topRentedCards"`
}
