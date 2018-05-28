package main

import (
	"time"
)

type EffortType string

type Effort struct {
	Amount int
	Type   EffortType
}

type PriorityType string

type Priority struct {
	Amout int
	Type  PriorityType
}

type Task struct {
	DeadLine *time.Time
	DueDate  *time.Time
	Effort   *Effort //NOTE: Maybe this can be custom
	Name     string
	Priority *Priority //NOTE: Maybe this can be custom
	Status   *Status   //NOTE: Maybe this can be custom
}

type List struct {
	Name  *string
	Tasks []Task
}
