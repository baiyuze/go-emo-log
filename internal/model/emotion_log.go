package model

import "time"

type EmotionLog struct {
	ID         uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string      `gorm:"type:varchar(100);" json:"title"`
	Context    string      `gorm:"type:text;" json:"context"`
	UserID     uint64      `gorm:"not null;" json:"userId"`
	Audios     []*Resource `gorm:"foreignKey:EmoId;constraint:OnDelete:CASCADE;" json:"audios"`
	Images     []*Resource `gorm:"foreignKey:EmoId;constraint:OnDelete:CASCADE;" json:"images"`
	CreateTime time.Time   `gorm:"type:datetime(6);autoCreateTime;index;" json:"createTime"`
	UpdateTime time.Time   `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
	Emo        string      `gorm:"type:varchar(200)" json:"emo"`
	DeviceId   string      `gorm:"type:varchar(100);comment:设备id或者定义uuid" json:"deviceId"`

	//AudioTexts []*string   `gorm:"type:text;comment:语音转文字，存储的文字" json:"audioTexts"` // 存放在resource中
	//IsPlus      int         `gorm:"default:0;comment:0-普通用户,1-plus用户" json:"isPlus"` // 0,1
	//PlusId      uint64      `json:"plusId"`
	//PlusRecords *PlusRecord `gorm:"foreignKey:PlusId;references:ID" json:"plusRecords"`
}

func (EmotionLog) TableName() string {
	return "emotion_logs"
}
