package model

import (
	"time"

	"gorm.io/gorm"
)

type OperationLog struct {
	Base
	Category    uint8
	Action      string
	Object      string
	ObjectValue string
	Operator    string
	Detail      string
}

func (t OperationLog) TableName() string {
	return "operation_log"
}

func (t *OperationLog) BeforeCreate(*gorm.DB) (err error) {
	now := time.Now()
	t.CreateAt = &now
	t.UpdateAt = &now
	return
}

func (t *OperationLog) BeforeSave(*gorm.DB) (err error) {
	now := time.Now()
	t.UpdateAt = &now
	return
}

func (t *OperationLog) BeforeUpdate(*gorm.DB) (err error) {
	now := time.Now()
	t.UpdateAt = &now
	return
}
