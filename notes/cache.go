package notes

import (
	"time"
)

func GetNoteFromCache(id uint) (Note, bool) {
	mu.RLock()
	defer mu.RUnlock()
	note, exists := cache[id]
	return note, exists
}

func AddNoteToCache(note Note) {
	note.LastCall = time.Now()
	mu.Lock()
	cache[note.ID] = note
	mu.Unlock()
	SaveCacheToFile()
}

func RemoveNoteFromCache(id uint) {
	mu.Lock()
	delete(cache, id)
	mu.Unlock()
	SaveCacheToFile()
}

func RemoveAllNotesFromCache() {
	mu.Lock()
	cache = make(map[uint]Note)
	mu.Unlock()
	SaveCacheToFile()
}

func isNoteExpired(note Note) bool {
	duration := time.Since(note.LastCall)
	return duration > cacheTime
}

func deleteExpiredNote() bool {
	mu.Lock()
	var expiredID uint
	var found bool
	for id, note := range cache {
		if isNoteExpired(note) {
			expiredID = id
			found = true
			break
		}
	}

	if !found {
		mu.Unlock()
		return false
	}

	delete(cache, expiredID)
	mu.Unlock()
	SaveCacheToFile()
	return true
}

func ClearNoteFromCache() {
	// Сигнализируем о начале очистки устаревших заметок
	signalCacheClearing(true)

	for deleteExpiredNote() {
		time.Sleep(3 * time.Second)
	}

	// Даем время на обновление состояния всем горутинам
	time.Sleep(100 * time.Millisecond)

	// Сигнализируем об окончании очистки устаревших заметок
	signalCacheClearing(false)
}
