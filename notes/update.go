package notes

import (
	"fmt"
	"strings"
)

func Update() {
	var notes []Note
	result := DB.Find(&notes)
	if result.Error != nil {
		fmt.Println("Ошибка загрузки заметок:", result.Error)
		return
	}

	if len(notes) == 0 {
		fmt.Println("❌ Доступный заметок нет")
		return
	}

	fmt.Println("\n========== ДОСТУПНЫЕ ЗАМЕТКИ ==========")
	for _, note := range notes {
		fmt.Printf("[%d] %s\n", note.ID, note.Title)
	}
	fmt.Println("=======================================")

	var inp uint
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&inp)

	UpdateNote(inp)
}

func UpdateNote(id uint) {
	var note Note
	result := DB.First(&note, id)
	if result.Error != nil {
		fmt.Printf("❌ Заметка с номером %d не найдена.\n", id)
		return
	}
	fmt.Printf("Введите название заметки: ")
	note.Title, _ = Reader.ReadString('\n')
	note.Title = strings.TrimSpace(note.Title)

	fmt.Printf("Введите содержание заметки: ")
	note.Content, _ = Reader.ReadString('\n')
	note.Content = strings.TrimSpace(note.Content)

	DB.Save(note)
}
