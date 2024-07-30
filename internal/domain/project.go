package domain

import (
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	ManagerID   uuid.UUID `json:"manager_id"`
}
