package model

import (
	"time"
)

type Version struct {
	ID          uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Version     string      `gorm:"type:varchar(100);not null" json:"version"`
	Feedbacks   []*Feedback `gorm:"foreignKey:VersionId"`
	Description string      `gorm:"type:varchar(200)" json:"description"`
	CreateTime  time.Time   `gorm:"type:datetime(6);autoCreateTime;index" json:"createTime"`
	UpdateTime  time.Time   `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
}

func (Version) TableName() string {
	return "versions"
}
