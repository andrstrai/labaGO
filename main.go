package main

import (
	"fmt"
	"os"
)

type students struct {
	name         string
	databirthday string
	institute    string
	stipend      float64
	GPA          float64
}

// Странная функция
func record() {
	var name, data, inst string
	var sti, GPA float64
	fmt.Println("Укажите имя студента: ")
	fmt.Scan(&name)
	fmt.Println("Укажите дату рождения студента: ")
	fmt.Scan(&data)
	fmt.Println("Укажите институт студента: ")
	fmt.Scan(&inst)
	fmt.Println("Укажите стипендию студента: ")
	fmt.Scan(&sti)
	fmt.Println("Укажите средний балл студента: ")
	fmt.Scan(&GPA)
	student := students{name, data, inst, sti, GPA}
	info := []byte(student.name + "/" + student.databirthday + "/" + student.institute + "/" + fmt.Sprint(student.stipend) + "/" + fmt.Sprint(student.GPA))
	f, e := os.OpenFile("Base.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		fmt.Println("Ошибка чтения")
	}
	f.WriteString(string(info))
}
func main() {
	var a int64

	for {
		fmt.Println("1 - Показать список всех ")
		fmt.Println("2 - Запись студента в базу данных")
		fmt.Println("Выберите пункт из меню управления: ")
		fmt.Scan(&a)
		if a == 1 {
			record()
		}
		if a == 0 {
			break
		}
	}

}
