package notes

import "fmt"

func ShowAll() {
	var notes []Note
	notes = GetAllNotesFromCache()
	isFromCache := true
	if len(notes) == 0 {
		result := DB.Find(&notes)
		if result.Error != nil {
			fmt.Println("Ошибка загрузки заметок:", result.Error)
			return
		}
		isFromCache = false
	}

	if len(notes) == 0 {
		fmt.Println("❌ Заметок нет")
		return
	}

	fmt.Println("\n===Все заметки===")
	for _, note := range notes {
		fmt.Printf("\n[%d] %s\n", note.ID, note.Title)
		fmt.Printf("\t%s\n", note.Content)
		if !isFromCache {
			fmt.Printf("Заметка #%d загружена из базы данных и добавлена в кэш.\n", note.ID)
			AddNoteToCache(note)
		}
	}
	fmt.Println("=================")
}

func ShowNote(id uint) {
	var note Note
	note, exists := GetNoteFromCache(id)
	isFromCache := true
	if !exists {
		result := DB.First(&note, id)
		if result.Error != nil {
			fmt.Printf("❌ Заметка с номером %d не найдена.\n", id)
			return
		}
		isFromCache = false
	}

	if !isFromCache {
		fmt.Printf("Заметка #%d загружена из базы данных и добавлена в кэш.\n", note.ID)
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
