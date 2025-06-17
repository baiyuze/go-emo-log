package model

import (
	"time"
)

type Permission struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(30);not null" json:"name"`
	Roles       []*Role   `gorm:"many2many:role_permissions"`
	Description string    `gorm:"type:varchar(200);" json:"description"`
	CreateTime  time.Time `gorm:"type:datetime(6);autoCreateTime" json:"createTime"`
	UpdateTime  time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
}

func (Permission) TableName() string {
	return "permissions"
}
