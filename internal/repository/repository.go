package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/zedzerofreedomtime/pilatesreformer/api/internal/store"
)

var ErrNotFound = errors.New("not found")

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(ctx context.Context, user store.User) error {
	_, err := r.db.Exec(ctx, `
		insert into users (id, name, email, phone, role_id, password_hash, trainer_id)
		values ($1, $2, $3, $4, $5, $6, $7)
	`, user.ID, user.Name, user.Email, user.Phone, user.RoleID, user.PasswordHash, user.TrainerID)
	return err
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (store.User, error) {
	var user store.User
	var phone sql.NullString
	var trainerID sql.NullString

	err := r.db.QueryRow(ctx, `
		select id, name, email, phone, role_id, password_hash, trainer_id, created_at
		from users
		where email = $1
	`, email).Scan(&user.ID, &user.Name, &user.Email, &phone, &user.RoleID, &user.PasswordHash, &trainerID, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return store.User{}, ErrNotFound
		}
		return store.User{}, err
	}

	if phone.Valid {
		user.Phone = phone.String
	}
	if trainerID.Valid {
		value := trainerID.String
		user.TrainerID = &value
	}

	return user, nil
}

func (r *Repository) ListEquipment(ctx context.Context) ([]store.Equipment, error) {
	rows, err := r.db.Query(ctx, `
		select id, name, image, badge, monthly_rate, trainer_mode, summary, ideal_for, footprint, features
		from equipment
		order by created_at, name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []store.Equipment{}
	for rows.Next() {
		var item store.Equipment
		var features []string
		if err := rows.Scan(&item.ID, &item.Name, &item.Image, &item.Badge, &item.MonthlyRate, &item.TrainerMode, &item.Summary, &item.IdealFor, &item.Footprint, &features); err != nil {
			return nil, err
		}
		item.Features = features
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *Repository) CreateEquipment(ctx context.Context, equipment store.Equipment) error {
	_, err := r.db.Exec(ctx, `
		insert into equipment (id, name, image, badge, monthly_rate, trainer_mode, summary, ideal_for, footprint, features)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`, equipment.ID, equipment.Name, equipment.Image, equipment.Badge, equipment.MonthlyRate, equipment.TrainerMode, equipment.Summary, equipment.IdealFor, equipment.Footprint, equipment.Features)
	return err
}

func (r *Repository) UpdateEquipment(ctx context.Context, equipment store.Equipment) error {
	commandTag, err := r.db.Exec(ctx, `
		update equipment
		set name = $2,
		    image = $3,
		    badge = $4,
		    monthly_rate = $5,
		    trainer_mode = $6,
		    summary = $7,
		    ideal_for = $8,
		    footprint = $9,
		    features = $10
		where id = $1
	`, equipment.ID, equipment.Name, equipment.Image, equipment.Badge, equipment.MonthlyRate, equipment.TrainerMode, equipment.Summary, equipment.IdealFor, equipment.Footprint, equipment.Features)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *Repository) DeleteEquipment(ctx context.Context, id string) error {
	commandTag, err := r.db.Exec(ctx, `delete from equipment where id = $1`, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *Repository) ListTrainers(ctx context.Context) ([]store.Trainer, error) {
	rows, err := r.db.Query(ctx, `
		select id, name, image, specialty, session_rate, availability, summary, schedule_window,
		       available_slots, booked_slots, machine_focus, exercise_focus, weekly_schedule
		from trainers
		order by created_at, name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []store.Trainer{}
	for rows.Next() {
		item, err := scanTrainer(rows)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *Repository) GetTrainerByID(ctx context.Context, trainerID string) (store.Trainer, error) {
	row := r.db.QueryRow(ctx, `
		select id, name, image, specialty, session_rate, availability, summary, schedule_window,
		       available_slots, booked_slots, machine_focus, exercise_focus, weekly_schedule
		from trainers
		where id = $1
	`, trainerID)

	item, err := scanTrainer(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return store.Trainer{}, ErrNotFound
		}
		return store.Trainer{}, err
	}

	return item, nil
}

func (r *Repository) CreateTrainer(ctx context.Context, trainer store.Trainer) error {
	weeklySchedule, err := json.Marshal(trainer.WeeklySchedule)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, `
		insert into trainers (
		  id, name, image, specialty, session_rate, availability, summary, schedule_window,
		  available_slots, booked_slots, machine_focus, exercise_focus, weekly_schedule
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`, trainer.ID, trainer.Name, trainer.Image, trainer.Specialty, trainer.SessionRate, trainer.Availability, trainer.Summary, trainer.ScheduleWindow, trainer.AvailableSlots, trainer.BookedSlots, trainer.MachineFocus, trainer.ExerciseFocus, weeklySchedule)
	return err
}

func (r *Repository) UpdateTrainer(ctx context.Context, trainer store.Trainer) error {
	weeklySchedule, err := json.Marshal(trainer.WeeklySchedule)
	if err != nil {
		return err
	}

	commandTag, err := r.db.Exec(ctx, `
		update trainers
		set name = $2,
		    image = $3,
		    specialty = $4,
		    session_rate = $5,
		    availability = $6,
		    summary = $7,
		    schedule_window = $8,
		    available_slots = $9,
		    booked_slots = $10,
		    machine_focus = $11,
		    exercise_focus = $12,
		    weekly_schedule = $13
		where id = $1
	`, trainer.ID, trainer.Name, trainer.Image, trainer.Specialty, trainer.SessionRate, trainer.Availability, trainer.Summary, trainer.ScheduleWindow, trainer.AvailableSlots, trainer.BookedSlots, trainer.MachineFocus, trainer.ExerciseFocus, weeklySchedule)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *Repository) DeleteTrainer(ctx context.Context, id string) error {
	commandTag, err := r.db.Exec(ctx, `delete from trainers where id = $1`, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *Repository) ListTrainerClients(ctx context.Context, trainerID string) ([]store.TrainerClient, error) {
	rows, err := r.db.Query(ctx, `
		select id, client_name, equipment_name, plan_name, next_session, contact, status
		from trainer_clients
		where trainer_id = $1
		order by created_at desc
	`, trainerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []store.TrainerClient{}
	for rows.Next() {
		var item store.TrainerClient
		if err := rows.Scan(&item.ID, &item.Name, &item.EquipmentName, &item.PlanName, &item.NextSession, &item.Contact, &item.Status); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *Repository) CreateTrainerApplication(ctx context.Context, application store.TrainerApplication) error {
	_, err := r.db.Exec(ctx, `
		insert into trainer_applications (id, name, email, phone, password_hash, specialty, machine_focus)
		values ($1, $2, $3, $4, $5, $6, $7)
	`, application.ID, application.Name, application.Email, application.Phone, application.PasswordHash, application.Specialty, application.MachineFocus)
	return err
}

func (r *Repository) ListPendingTrainerApplications(ctx context.Context) ([]store.TrainerApplication, error) {
	rows, err := r.db.Query(ctx, `
		select id, name, email, phone, password_hash, specialty, machine_focus, status, submitted_at
		from trainer_applications
		where status = 'pending'
		order by submitted_at desc
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []store.TrainerApplication{}
	for rows.Next() {
		var item store.TrainerApplication
		var machineFocus []string
		if err := rows.Scan(&item.ID, &item.Name, &item.Email, &item.Phone, &item.PasswordHash, &item.Specialty, &machineFocus, &item.Status, &item.SubmittedAt); err != nil {
			return nil, err
		}
		item.MachineFocus = machineFocus
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *Repository) ApproveTrainerApplication(ctx context.Context, applicationID string, trainer store.Trainer) (store.Trainer, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return store.Trainer{}, err
	}
	defer tx.Rollback(ctx)

	var application store.TrainerApplication
	var machineFocus []string
	err = tx.QueryRow(ctx, `
		select id, name, email, phone, password_hash, specialty, machine_focus, status, submitted_at
		from trainer_applications
		where id = $1 and status = 'pending'
	`, applicationID).Scan(&application.ID, &application.Name, &application.Email, &application.Phone, &application.PasswordHash, &application.Specialty, &machineFocus, &application.Status, &application.SubmittedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return store.Trainer{}, ErrNotFound
		}
		return store.Trainer{}, err
	}
	application.MachineFocus = machineFocus

	trainer.Name = application.Name
	trainer.Specialty = application.Specialty
	if len(application.MachineFocus) > 0 {
		trainer.MachineFocus = application.MachineFocus
	}

	weeklySchedule, err := json.Marshal(trainer.WeeklySchedule)
	if err != nil {
		return store.Trainer{}, err
	}

	_, err = tx.Exec(ctx, `
		insert into trainers (
		  id, name, image, specialty, session_rate, availability, summary, schedule_window,
		  available_slots, booked_slots, machine_focus, exercise_focus, weekly_schedule
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`, trainer.ID, trainer.Name, trainer.Image, trainer.Specialty, trainer.SessionRate, trainer.Availability, trainer.Summary, trainer.ScheduleWindow, trainer.AvailableSlots, trainer.BookedSlots, trainer.MachineFocus, trainer.ExerciseFocus, weeklySchedule)
	if err != nil {
		return store.Trainer{}, err
	}

	_, err = tx.Exec(ctx, `
		insert into users (id, name, email, phone, role_id, password_hash, trainer_id)
		values ($1, $2, $3, $4, 'trainer', $5, $6)
	`, fmt.Sprintf("user-%s", uuid.NewString()[:8]), trainer.Name, application.Email, application.Phone, application.PasswordHash, trainer.ID)
	if err != nil {
		return store.Trainer{}, err
	}

	_, err = tx.Exec(ctx, `update trainer_applications set status = 'approved' where id = $1`, applicationID)
	if err != nil {
		return store.Trainer{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return store.Trainer{}, err
	}

	return trainer, nil
}

func (r *Repository) RejectTrainerApplication(ctx context.Context, applicationID string) error {
	commandTag, err := r.db.Exec(ctx, `update trainer_applications set status = 'rejected' where id = $1`, applicationID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *Repository) ListRentalPlans(ctx context.Context) ([]store.RentalPlan, error) {
	rows, err := r.db.Query(ctx, `
		select id, name, months, discount, optional_sessions, required_sessions, note
		from rental_plans
		order by months
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []store.RentalPlan{}
	for rows.Next() {
		var item store.RentalPlan
		if err := rows.Scan(&item.ID, &item.Name, &item.Months, &item.Discount, &item.OptionalSessions, &item.RequiredSessions, &item.Note); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *Repository) ListTrainerServicePlans(ctx context.Context) ([]store.TrainerServicePlan, error) {
	rows, err := r.db.Query(ctx, `
		select id, name, sessions, discount, note
		from trainer_service_plans
		order by sessions
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []store.TrainerServicePlan{}
	for rows.Next() {
		var item store.TrainerServicePlan
		if err := rows.Scan(&item.ID, &item.Name, &item.Sessions, &item.Discount, &item.Note); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *Repository) GetHomeContent(ctx context.Context) (map[string]any, error) {
	var payload []byte
	err := r.db.QueryRow(ctx, `
		select payload
		from home_page_contents
		order by updated_at desc
		limit 1
	`).Scan(&payload)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	var content map[string]any
	if err := json.Unmarshal(payload, &content); err != nil {
		return nil, err
	}

	return content, nil
}

func (r *Repository) SaveHomeContent(ctx context.Context, content map[string]any) error {
	payload, err := json.Marshal(content)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, `
		insert into home_page_contents (id, payload)
		values ($1, $2)
	`, fmt.Sprintf("home-content-%s", uuid.NewString()[:8]), payload)
	return err
}

func (r *Repository) CreateBookingInquiry(ctx context.Context, inquiry store.BookingInquiry) error {
	_, err := r.db.Exec(ctx, `
		insert into booking_inquiries (
		  id, name, email, phone, mode_id, equipment_id, rental_plan_id, trainer_id,
		  trainer_service_plan_id, grand_total, notes, status
		)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`, inquiry.ID, inquiry.Name, inquiry.Email, inquiry.Phone, inquiry.ModeID, inquiry.EquipmentID, inquiry.RentalPlanID, inquiry.TrainerID, inquiry.TrainerServicePlanID, inquiry.GrandTotal, inquiry.Notes, inquiry.Status)
	return err
}

func (r *Repository) ListBookingInquiries(ctx context.Context) ([]store.BookingInquiry, error) {
	rows, err := r.db.Query(ctx, `
		select id, name, email, phone, mode_id, equipment_id, rental_plan_id, trainer_id,
		       trainer_service_plan_id, grand_total, notes, status, created_at
		from booking_inquiries
		order by created_at desc
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []store.BookingInquiry{}
	for rows.Next() {
		var item store.BookingInquiry
		var equipmentID, rentalPlanID, trainerID, trainerServicePlanID sql.NullString
		var notes sql.NullString

		if err := rows.Scan(&item.ID, &item.Name, &item.Email, &item.Phone, &item.ModeID, &equipmentID, &rentalPlanID, &trainerID, &trainerServicePlanID, &item.GrandTotal, &notes, &item.Status, &item.CreatedAt); err != nil {
			return nil, err
		}

		if equipmentID.Valid {
			value := equipmentID.String
			item.EquipmentID = &value
		}
		if rentalPlanID.Valid {
			value := rentalPlanID.String
			item.RentalPlanID = &value
		}
		if trainerID.Valid {
			value := trainerID.String
			item.TrainerID = &value
		}
		if trainerServicePlanID.Valid {
			value := trainerServicePlanID.String
			item.TrainerServicePlanID = &value
		}
		if notes.Valid {
			item.Notes = notes.String
		}

		items = append(items, item)
	}

	return items, rows.Err()
}

type scanner interface {
	Scan(dest ...any) error
}

func scanTrainer(row scanner) (store.Trainer, error) {
	var item store.Trainer
	var machineFocus []string
	var exerciseFocus []string
	var weeklySchedule []byte

	err := row.Scan(&item.ID, &item.Name, &item.Image, &item.Specialty, &item.SessionRate, &item.Availability, &item.Summary, &item.ScheduleWindow, &item.AvailableSlots, &item.BookedSlots, &machineFocus, &exerciseFocus, &weeklySchedule)
	if err != nil {
		return store.Trainer{}, err
	}

	item.MachineFocus = machineFocus
	item.ExerciseFocus = exerciseFocus
	if err := json.Unmarshal(weeklySchedule, &item.WeeklySchedule); err != nil {
		return store.Trainer{}, err
	}

	return item, nil
}
