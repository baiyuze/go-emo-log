package model

import (
	"time"
)

type Device struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint64    `gorm:"index" json:"userId,omitempty"`
	User        *User     `gorm:"foreignKey:UserID;references:ID" json:"user"`
	DeviceID    string    `gorm:"type:varchar(64);not null;uniqueIndex" json:"deviceId"`
	Platform    string    `gorm:"type:varchar(16)" json:"platform,omitempty"`
	Brand       string    `gorm:"type:varchar(32)" json:"brand,omitempty"`
	Model       string    `gorm:"type:varchar(32)" json:"model,omitempty"`
	OSVersion   string    `gorm:"type:varchar(32)" json:"osVersion,omitempty"`
	Resolution  string    `gorm:"type:varchar(32)" json:"resolution,omitempty"`
	AppVersion  string    `gorm:"type:varchar(16)" json:"appVersion,omitempty"`
	CreatedTime time.Time `gorm:"type:datetime(6);autoCreateTime" json:"createdTime"`
	UpdatedTime time.Time `gorm:"type:datetime(6);autoUpdateTime" json:"updatedTime"`
}

func (Device) TableName() string {
	return "devices"
}
