package notes

import (
	"fmt"
	"strings"
)

func Update() {
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

	UpdateNote(inp)
}

func UpdateNote(id uint) {
	var note Note
	note, exists := GetNoteFromCache(id)
	if !exists {
		result := DB.First(&note, id)
		if result.Error != nil {
			fmt.Printf("❌ Заметка с номером %d не найдена.\n", id)
			return
		}
	}
	fmt.Printf("Введите название заметки: ")
	note.Title, _ = Reader.ReadString('\n')
	note.Title = strings.TrimSpace(note.Title)

	fmt.Printf("Введите содержание заметки: ")
	note.Content, _ = Reader.ReadString('\n')
	note.Content = strings.TrimSpace(note.Content)

	fmt.Print("Вы уверены? Данные будут обновлены. (y/n): ")
	approve, _ := Reader.ReadString('\n')
	approve = strings.TrimSpace(approve)

	if approve == "y" || approve == "Y" {
		result := DB.Save(&note)
		if result.Error != nil {
			fmt.Println("Ошибка при обновлении заметки:", result.Error)
		} else {
			fmt.Println("Заметка успешно обновлена!")
			AddNoteToCache(note)
		}
	} else {
		fmt.Println("Удаление отменено")
	}
}
