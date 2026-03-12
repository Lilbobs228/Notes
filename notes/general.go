package notes

import "fmt"

var cache = make(map[uint]Note)

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

func ClearCache() {
	cache = make(map[uint]Note)
	SaveCacheToFile()
}

func GetAllNotesFromCache() []Note {
	notes := make([]Note, 0, len(cache))
	for _, note := range cache {
		notes = append(notes, note)
	}
	return notes
}

func ShowAvailableNotes() (bool, error) {
	var notes []Note
	notes = GetAllNotesFromCache()
	if len(notes) == 0 {
		result := DB.Find(&notes)
		if result.Error != nil {
			fmt.Println("Ошибка загрузки заметок:", result.Error)
			return false, result.Error
		}
	}

	if len(notes) == 0 {
		fmt.Println("❌ Доступный заметок нет")
		return false, nil
	}
	fmt.Println("\n========== ДОСТУПНЫЕ ЗАМЕТКИ ==========")
	for _, note := range notes {
		fmt.Printf("[%d] %s\n", note.ID, note.Title)
	}
	fmt.Println("=======================================")
	return true, nil
}
