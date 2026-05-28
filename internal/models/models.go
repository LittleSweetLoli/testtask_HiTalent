package models

import "time"

type Department struct {
	ID        uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string       `gorm:"not null;size:200"        json:"name"`
	ParentID  *uint        `gorm:"index"                    json:"parent_id"`
	CreatedAt time.Time    `                                json:"created_at"`
	Children  []Department `gorm:"foreignKey:ParentID"      json:"children,omitempty"`
	Employees []Employee   `gorm:"foreignKey:DepartmentID"  json:"employees,omitempty"`
}

type Employee struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	DepartmentId int        `gorm:"not null;index"           json:"department_id"`
	FullName     string     `gorm:"size:200;not null"        json:"full_name"`
	Position     string     `gorm:"size:200;not null"        json:"position"`
	HiredAt      *time.Time `                                json:"hired_at"`
	CreatedAt    time.Time  `                                json:"created_at"`
}
