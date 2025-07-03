package model

import (
	"time"
)

type User struct {
	ID          uint64        `gorm:"primaryKey;autoIncrement" json:"id"`
	DeviceId    uint64        `gorm:"index" json:"deviceId"`
	Device      uint64        `gorm:"foreignKey:UserId;references:ID" json:"device"`
	Name        string        `gorm:"type:varchar(30);not null" json:"name"`
	PlusRecord  []*PlusRecord `gorm:"foreignKey:UserID;references:ID" json:"plusRecord"`
	EmoRecords  []*EmotionLog `gorm:"foreignKey:UserID;references:ID" json:"emoRecords"`
	Account     string        `gorm:"type:varchar(30);not null;uniqueIndex:idx_account_create_time,sort:asc" json:"account"`
	CreateTime  time.Time     `gorm:"type:datetime(6);uniqueIndex:idx_account_create_time,sort:asc;index" json:"createTime"`
	Password    *string       `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	Roles       []*Role       `gorm:"many2many:user_roles"`
	Departments []*Department `gorm:"many2many:user_departments;" json:"departments,omitempty"`
	Phone       string        `gorm:"type:varchar(30)" json:"phone"`
	Email       string        `gorm:"type:varchar(100)" json:"email"`
	UpdateTime  time.Time     `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
}

func (User) TableName() string {
	return "users"
}
