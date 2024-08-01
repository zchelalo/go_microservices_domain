package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	Id        string           `json:"id" gorm:"type:char(36);not null;primary_key"`
	UserId    string           `json:"user_id,omitempty" gorm:"type:char(36);not null"`
	User      *User            `json:"user,omitempty" gorm:"-"`
	CourseId  string           `json:"course_id,omitempty" gorm:"type:char(36);not null"`
	Course    *Course          `json:"course,omitempty" gorm:"-"`
	Status    EnrollmentStatus `json:"status" gorm:"type:char(2)"`
	CreatedAt *time.Time       `json:"-"`
	UpdatedAt *time.Time       `json:"-"`
	DeletedAt gorm.DeletedAt   `json:"-"`
}

type EnrollmentStatus string

const (
	Pending  EnrollmentStatus = "PP"
	Active   EnrollmentStatus = "AA"
	Studying EnrollmentStatus = "SS"
	Inactive EnrollmentStatus = "II"
)

func (enrollment *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {
	if enrollment.Id == "" {
		enrollment.Id = uuid.New().String()
	}
	return
}
