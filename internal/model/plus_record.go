package model

import (
	"time"
)

type PlusRecord struct {
	ID uint64 `gorm:"primaryKey;autoIncrement;comment:主键 ID" json:"id"`
	// 需要关联用户
	UserID    uint64    `gorm:"not null;uniqueIndex;comment:用户ID" json:"userId"`
	Level     int       `gorm:"type:int;default:1;comment:会员等级（1=普通，2=高级等）" json:"level"`
	StartTime time.Time `gorm:"type:datetime(6);not null;comment:开始时间" json:"startTime"`
	EndTime   time.Time `gorm:"type:datetime(6);not null;comment:结束时间" json:"endTime"`

	CreateTime time.Time `gorm:"type:datetime(6);autoCreateTime;comment:创建时间" json:"createTime"`
	UpdateTime time.Time `gorm:"type:datetime(6);autoUpdateTime;comment:更新时间" json:"updateTime"`
}

func (PlusRecord) TableName() string {
	return "plus_records"
}
