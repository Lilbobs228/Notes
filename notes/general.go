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
var cacheTime = 10 * time.Second
var mu sync.RWMutex

var (
	isCacheClearing bool
	cacheStateMu    sync.RWMutex // Mutex для защиты флага состояния
)

type Note struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	LastCall time.Time
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
}

// Проверяет, очищается ли кэш в данный момент. Если да, то кэшем не пользуемся.
func ShouldBypassCache() bool {
	cacheStateMu.RLock()
	defer cacheStateMu.RUnlock()
	return isCacheClearing
}

// Устанавливает флаг очистки кэша
func signalCacheClearing(clearing bool) {
	cacheStateMu.Lock()
	defer cacheStateMu.Unlock()
	isCacheClearing = clearing
}
