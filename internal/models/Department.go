package models

import "time"

type Department struct {
	id         int
	name       string
	parent_id  int
	created_at time.Time
}
