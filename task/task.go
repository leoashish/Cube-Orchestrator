package task

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	EndTime       time.Time
}

// TaskEvent struct
// To start and stop a task

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	TimeStamp time.Time
	Task      Task
}
