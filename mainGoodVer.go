package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type students struct {
	name         string
	databirthday string
	institute    string
	stipend      float64
	GPA          float64
}

func (st students) ToString() string {
	return fmt.Sprintf("Имя студента: %s\nДата рождения: %s\n"+
		"Институт: %s\nСтипендия: %f\nСредний балл: %f",
		st.name, st.databirthday, st.institute, st.stipend, st.GPA)
}

// функция добавления студента в слайс
func record(student *[]students) {
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
	info := []byte(name + "/" + data + "/" + inst + "/" + fmt.Sprint(sti) + "/" + fmt.Sprint(GPA))
	f, e := os.OpenFile("Base.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	*student = append(*student, students{name, data, inst, sti, GPA})
	if e != nil {
		fmt.Println("Ошибка чтения")
	}
	if len("Base.txt") == 0 {
		f.WriteString(string(info))
	} else {
		f.WriteString("\n" + string(info))
	}
}

// сортирует слайс студентов по среднему баллу по убыванию
func sort_by_grades(students_list []students) {
	slices.SortFunc(students_list, func(a, b students) int {
		switch {
		case a.GPA < b.GPA:
			return 1
		case a.GPA > b.GPA:
			return -1
		default:
			return 0
		}
	})
	fmt.Println("Данные отсортированы!")
}

// функция для вывода студентов
func print_all(student []students) {
	for i := 0; i < len(student); i++ {
		if i == 0 {
			fmt.Println("---")
		} else {
			fmt.Println(student[i].ToString())
			fmt.Println("---")
		}
	}
	fmt.Println("Всего студентов:", len(student))
	fmt.Println("---")
}

func main() {
	var a int64
	all_students := []students{}

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
	for _, line := range lines {
		line := strings.Split(line, "/")
		sti, err := strconv.ParseFloat(line[3], 64)
		gpa, err := strconv.ParseFloat(line[4], 64)
		if err != nil {
			fmt.Println("Ошибка чтения данных!", err)
			return
		}
		all_students = append(all_students, students{line[0], line[1], line[2], sti, gpa})
	}

	for {
		fmt.Println("1 - Показать список всех ")
		fmt.Println("2 - Запись студента в базу данных")
		fmt.Println("3 - Сортировать студентов по среднему баллу")
		fmt.Println("0 - Завершить работу")
		fmt.Println("Выберите пункт из меню управления: ")
		fmt.Scan(&a)
		if a == 1 {
			print_all(all_students)
		} else if a == 0 {
			break
		} else if a == 2 {
			record(&all_students)
		} else if a == 3 {
			sort_by_grades(all_students)
		}
	}
}
