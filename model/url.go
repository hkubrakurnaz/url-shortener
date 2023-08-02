package model

import (
	"time"
)

type Url struct {
	ID        uint64    `gorm:"type:bigint;primary_key" json:"id,omitempty"`
	LongUrl   string    `gorm:"varchar(100);not null" json:"longUrl,omitempty"`
	ShortUrl  string    `gorm:"varchar(64);" json:"shortUrl,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"createdAt,omitempty"`
}
