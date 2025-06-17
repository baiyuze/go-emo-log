package model

import "time"

type Resource struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	URL        string    `json:"url"`
	Text       string    `gorm:"type:text" json:"text"`
	EmoId      uint64    `gorm:"index;comment:emotion的主键ID" json:"emoId"` // 对应 emotion_log_id
	Type       string    `gorm:"type:varchar(100);not null" json:"type"`  // "image" or "audio"
	CreateTime time.Time `gorm:"type:datetime(6);autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
}

func (Resource) TableName() string { return "resources" }
