package notes

import "fmt"

func ShowAll() {
	var notes []Note
	notes = GetAllNotesFromCache()
	if len(notes) == 0 {
		result := DB.Find(&notes)
		if result.Error != nil {
			fmt.Println("Ошибка загрузки заметок:", result.Error)
			return
		}
	}

	if len(notes) == 0 {
		fmt.Println("❌ Заметок нет")
		return
	}

	fmt.Println("\n===Все заметки===")
	for _, note := range notes {
		fmt.Printf("\n[%d] %s\n", note.ID, note.Title)
		fmt.Printf("\t%s\n", note.Content)
		AddNoteToCache(note)
	}
	fmt.Println("=================")
}

func ShowNote(id uint) {
	var note Note
	note, exists := GetNoteFromCache(id)
	if !exists {
		result := DB.First(&note, id)
		if result.Error != nil {
			fmt.Printf("❌ Заметка с номером %d не найдена.\n", id)
			return
		}
		AddNoteToCache(note)
	}

	fmt.Printf("\n========== ЗАМЕТКА #%d ==========\n", note.ID)
	fmt.Printf("\n[%d] %s\n", note.ID, note.Title)
	fmt.Printf("\t%s\n", note.Content)
	fmt.Print("================================\n\n")
}

func ShowNoteByChoice() {
	exists, err := ShowAvailableNotes()
	if err != nil {
		fmt.Println("Ошибка при загрузке заметок:", err)
		return
	}
	if !exists {
		return
	}

	var inp uint
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&inp)

	ShowNote(inp)
}
