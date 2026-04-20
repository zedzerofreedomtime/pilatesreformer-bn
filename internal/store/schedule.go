package store

import "fmt"

func DefaultWeeklySchedule() []WeeklyScheduleDay {
	weekDays := []struct {
		ID         string
		Label      string
		ShortLabel string
	}{
		{ID: "sun", Label: "Sunday", ShortLabel: "Sun"},
		{ID: "mon", Label: "Monday", ShortLabel: "Mon"},
		{ID: "tue", Label: "Tuesday", ShortLabel: "Tue"},
		{ID: "wed", Label: "Wednesday", ShortLabel: "Wed"},
		{ID: "thu", Label: "Thursday", ShortLabel: "Thu"},
		{ID: "fri", Label: "Friday", ShortLabel: "Fri"},
		{ID: "sat", Label: "Saturday", ShortLabel: "Sat"},
	}

	result := make([]WeeklyScheduleDay, 0, len(weekDays))
	for _, day := range weekDays {
		slots := make([]Slot, 0, 9)
		for hour := 8; hour <= 16; hour++ {
			start := fmt.Sprintf("%02d:00", hour)
			end := fmt.Sprintf("%02d:00", hour+1)
			slots = append(slots, Slot{
				Key:    fmt.Sprintf("%s-%s", day.ID, start),
				Label:  fmt.Sprintf("%s - %s", start, end),
				Status: "available",
			})
		}

		result = append(result, WeeklyScheduleDay{
			ID:             day.ID,
			Label:          day.Label,
			ShortLabel:     day.ShortLabel,
			AvailableCount: len(slots),
			BookedCount:    0,
			Slots:          slots,
		})
	}

	return result
}

func SummarizeSchedule(schedule []WeeklyScheduleDay) (int, int) {
	availableSlots := 0
	bookedSlots := 0

	for dayIndex := range schedule {
		dayAvailable := 0
		dayBooked := 0
		for _, slot := range schedule[dayIndex].Slots {
			if slot.Status == "booked" {
				dayBooked++
				continue
			}
			dayAvailable++
		}
		schedule[dayIndex].AvailableCount = dayAvailable
		schedule[dayIndex].BookedCount = dayBooked
		availableSlots += dayAvailable
		bookedSlots += dayBooked
	}

	return availableSlots, bookedSlots
}
