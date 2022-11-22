package models

type Tag struct {
	ID    uint64 `gorm:"primaryKey"`
	Tag   string `json:"href,omitempty"`
	Count uint64 `json:"href,omitempty"`
}
