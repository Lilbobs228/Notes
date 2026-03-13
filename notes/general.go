package notes

import (
	"bufio"
	"sync"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB
var Reader *bufio.Reader
var cache = make(map[uint]Note)
var cacheTime = 10 * time.Minute
var mu sync.Mutex

type Note struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
}
