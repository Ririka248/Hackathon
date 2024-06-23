package models

import "time"

type Calendar struct {
	ID          int
	UserID      int
	Content     string
	HealthScore int
	Date        time.Time
	CreatedAt   time.Time
}

type SCalendar struct {
	ID        int
	UserID    int
	hours     int
	minutes   int
	Date      time.Time
	CreatedAt time.Time
}

type FCalendar struct {
	ID        int
	UserID    int
	Content   string
	Amount    int
	Date      time.Time
	CreatedAt time.Time
}

type ECalendar struct {
	ID        int
	UserID    int
	Content   string
	Date      time.Time
	CreatedAt time.Time
}
