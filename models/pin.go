package models

import "time"

type Pin struct {
	ID          uint64 `gorm:"primaryKey"`
	Href        string `json:"href,omitempty"`
	Description string `json:"description,omitempty"`
	Extended    string `json:"extended,omitempty"`
	Meta        string `json:"meta,omitempty"`
	Hash        string `json:"hash,omitempty"`
	Time        string `json:"time,omitempty"`
	Shared      string `json:"shared,omitempty"`
	Toread      string `json:"toread,omitempty"`
	Tags        string `json:"tags,omitempty"`
}

func (p Pin) isShared() bool {
	if p.Shared == "yes" {
		return true
	} else {
		return false
	}
}

func (p Pin) isToRead() bool {
	if p.Toread == "no" {
		return false
	} else {
		return true
	}
}

func (p Pin) timeCreated() time.Time {
	// . "2022-11-21T01:49:01Z"

}
