package notes

import (
	"encoding/json"
	"os"
)

const notesFile = "cache.json"

func InitializeCacheFile() {
	if _, err := os.Stat(notesFile); os.IsNotExist(err) {
		var notes []Note
		data, _ := json.MarshalIndent(notes, "", "  ")
		os.WriteFile(notesFile, data, 0644)
	}
	ClearCacheFile()
}

func SaveCacheToFile() {
	notes := GetAllNotesFromCache()
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(notesFile, data, 0644)
}

func ClearCacheFile() {
	var notes []Note
	data, _ := json.MarshalIndent(notes, "", "  ")
	os.WriteFile(notesFile, data, 0644)
}
