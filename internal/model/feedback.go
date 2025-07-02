package model

import (
	"time"
)

type Feedback struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Version     string    `gorm:"not null" json:"version"`
	VersionId   string    `gorm:"not null" json:"versionId"`
	Description string    `gorm:"type:varchar(200)" json:"description"`
	CreateTime  time.Time `gorm:"type:datetime(6);autoCreateTime;index" json:"createTime"`
	UpdateTime  time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
}

func (Feedback) TableName() string {
	return "feedbacks"
}
