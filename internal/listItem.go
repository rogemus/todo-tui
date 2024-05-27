package internal

import "time"

type Status int

const (
	TODO Status = iota
	IN_PROGRESS
	DONE
)

type Item struct {
	title       string
	description string
	created_at  time.Time
	status      Status
}


