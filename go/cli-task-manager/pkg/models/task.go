package models

import (
	"time"
)

/*
Representa la prioridad de una posible tarea
*/
type Priority string

const (
	PriorityHigh   Priority = "high"
	PriorityMedium Priority = "medium"
	PriorityLow    Priority = "low"
)

/*
Representa los estados de una posible tarea
*/
type State string

const (
	StatePending    State = "pending"
	StateInProgress State = "in_progress"
	StateCompleted  State = "completed"
	StateCancelled  State = "cancelled"
)

/*
Representa una tarea en el sistema
*/
type Task struct {
	ID          int
	Description string
	Priority    Priority
	State       State
	CreatedAt   time.Time
	CompletedAt *time.Time
}

/*
"Constructor" crea un una tarea
*/
func NewTask(description string, priority Priority) *Task {
	return &Task{
		Description: description,
		Priority:    priority,
		State:       StatePending,
		CreatedAt:   time.Now(),
	}
}

func IsValidPriority(p Priority) bool {
	switch p {
	case PriorityHigh, PriorityMedium, PriorityLow:
		return true
	}
	return false
}

// IsValidState verifica si un estado es válido
func IsValidState(s State) bool {
	switch s {
	case StatePending, StateInProgress, StateCompleted, StateCancelled:
		return true
	}
	return false
}
