package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	f.WriteString(string(info) + "\n")
}

func print_all() {
	file, err := os.Open("Base.txt")
	if err != nil {
		fmt.Println("Ошибка открытия:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
	}

	for i, line := range lines {
		corr_line := strings.Split(line, "/")
		if i == 0 {
			fmt.Println("---")
		}
		fmt.Printf("Имя студента: %s\nДата рождения: %s\nИнститут: %s\nСтипендия: %s\nСредний балл: %s\n---\n",
			corr_line[0], corr_line[1], corr_line[2], corr_line[3], corr_line[4])
	}
}

func main() {
	var a int64

	for {
		fmt.Println("1 - Показать список всех ")
		fmt.Println("2 - Запись студента в базу данных")
		fmt.Println("0 - Завершить работу")
		fmt.Println("Выберите пункт из меню управления: ")
		fmt.Scan(&a)
		if a == 1 {
			print_all()
		} else if a == 0 {
			break
		} else if a == 2 {
			record()
		}
	}

}
