package notes

import (
	"time"
)

func GetNoteFromCache(id uint) (Note, bool) {
	note, exists := cache[id]
	return note, exists
}

func AddNoteToCache(note Note) {
	cache[note.ID] = note
	SaveCacheToFile()
}

func RemoveNoteFromCache(id uint) {
	delete(cache, id)
	SaveCacheToFile()
}

func GetAllNotesFromCache() []Note {
	notes := make([]Note, 0, len(cache))
	for _, note := range cache {
		notes = append(notes, note)
	}
	return notes
}

func isNoteExpired(note Note) bool {
	duration := time.Since(note.CreatedAt)
	return duration > cacheTime
}

func deleteExpiredNote() bool {
	notes := GetAllNotesFromCache()
	for _, note := range notes {
		if isNoteExpired(note) {
			mu.Lock()
			delete(cache, note.ID)
			mu.Unlock()
			SaveCacheToFile()
			return true
		}
	}
	return false
}

func ClearNoteFromCache() {
	for deleteExpiredNote() {
		time.Sleep(1 * time.Second)
	}
}

func ClearCache() {
	cache = make(map[uint]Note)
	SaveCacheToFile()
}
