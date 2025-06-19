package model

import "time"

type EmotionLog struct {
	ID         uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string      `gorm:"type:varchar(100);" json:"title"`
	Content    string      `gorm:"type:text;not null" json:"content"`
	UserID     uint64      `gorm:"not null;index" json:"userId"`
	Audios     []*Resource `gorm:"foreignKey:EmoId;constraint:OnDelete:CASCADE;" json:"audios"`
	Images     []*Resource `gorm:"foreignKey:EmoId;constraint:OnDelete:CASCADE;" json:"images"`
	Date       time.Time   `gorm:"type:datetime(6);" json:"date"`
	CreateTime time.Time   `gorm:"type:datetime(6);autoCreateTime;index:idx_user_create_time" json:"createTime"`
	UpdateTime time.Time   `gorm:"type:datetime(6);autoUpdateTime" json:"updateTime"`
	Emo        string      `gorm:"type:varchar(200)" json:"emo"`
	DeviceId   string      `gorm:"type:varchar(100);comment:设备id或者定义uuid" json:"deviceId"`
	IsBackUp   int         `gorm:"type:int;default:0;comment:是否备份，1备份0未备份" json:"isBackUp"`
	//AudioTexts []*string   `gorm:"type:text;comment:语音转文字，存储的文字" json:"audioTexts"` // 存放在resource中
	//IsPlus      int         `gorm:"default:0;comment:0-普通用户,1-plus用户" json:"isPlus"` // 0,1
	//PlusId      uint64      `json:"plusId"`
	//PlusRecords *PlusRecord `gorm:"foreignKey:PlusId;references:ID" json:"plusRecords"`
}

func (EmotionLog) TableName() string {
	return "emotion_logs"
}
