package dto

type EmotionLog struct {
	Title    string    `json:"title"`             // 标题，例如 "糟糕"
	Content  string    `json:"content,omitempty"` // 内容，例如 "每次都想到很早之前的事情..."
	UserID   uint64    `json:"userId"`            // 用户 ID，例如 7
	Date     string    `json:"date"`              // 日期字符串，"2025-06-19 11:28:30"
	Emo      string    `json:"emo"`               // 情绪标签，例如 "内耗"
	IsBackup int       `json:"isBackup"`          // 是否备份，0 表示否
	Images   []*string `json:"images,omitempty"`
	Audios   []*string `json:"audios,omitempty"`
}
