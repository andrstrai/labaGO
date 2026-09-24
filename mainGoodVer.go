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
	ID           int
	name         string
	databirthday string
	institute    string
	stipend      float64
	GPA          float64
}

// метод для вывода красивой строки стуктуры
func (st students) ToString() string {
	return fmt.Sprintf("ID студента: %d\nИмя студента: %s\nДата рождения: %s\n"+
		"Институт: %s\nСтипендия: %f\nСредний балл: %f",
		st.ID, st.name, st.databirthday, st.institute, st.stipend, st.GPA)
}

// функция для генерации id
func max_id(students_list []students) int {
	if len(students_list) == 0 {
		return 0
	}
	max := slices.MaxFunc(students_list, func(a, b students) int {
		return a.ID - b.ID
	})
	return max.ID
}

// функция добавления студента в слайс
func record(student *[]students) {
	var name, data, inst string
	var sti, GPA float64
	id := max_id(*student) + 1
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
	info := []byte(fmt.Sprint(id) + "/" + name + "/" + data + "/" + inst + "/" + fmt.Sprint(sti) + "/" + fmt.Sprint(GPA))
	f, e := os.OpenFile("Base.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	*student = append(*student, students{id, name, data, inst, sti, GPA})
	if e != nil {
		fmt.Println("Ошибка чтения")
	}
	f.WriteString(string(info) + "\n")
}

// функция, которая сортирует слайс студентов по среднему баллу по убыванию
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

// функция, которая сортирует слайс студентов по размеру стипендии по убыванию
func sort_by_stipend(students_list []students) {
	slices.SortFunc(students_list, func(a, b students) int {
		switch {
		case a.stipend < b.stipend:
			return 1
		case a.stipend > b.stipend:
			return -1
		default:
			return 0
		}
	})
	fmt.Println("Данные отсортированы!")
}

// функция для вывода студентов списком
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

// функция для сохранения измененнного слайса в файл
func save_to_file(students_list []students) error {
	f, err := os.OpenFile("Base.txt", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка записи!")
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, student := range students_list {
		fmt.Fprintf(w, "%d/%s/%s/%s/%f/%f\n",
			student.ID, student.name, student.databirthday, student.institute, student.stipend, student.GPA)
	}
	fmt.Println("Изменения сохранены!")
	return w.Flush()
}

// основная функция
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
		id, err1 := strconv.ParseInt(line[0], 10, 0)
		sti, err2 := strconv.ParseFloat(line[4], 64)
		gpa, err3 := strconv.ParseFloat(line[5], 64)
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Ошибка чтения данных!", err)
			return
		}
		all_students = append(all_students, students{int(id), line[1], line[2], line[3], sti, gpa})
	}

	for {
		fmt.Println("1 - Показать список всех студентов")
		fmt.Println("2 - Запись студента в базу данных")
		fmt.Println("3 - Сортировать студентов по среднему баллу (по убыванию)")
		fmt.Println("4 - Сортировать студентов по размеру стипендии (по убыванию)")
		fmt.Println("5 - Сохранить изменения в файл")
		fmt.Println("0 - Завершить работу")
		fmt.Println("Выберите пункт из меню управления: ")
		fmt.Scan(&a)
		if a == 0 {
			break
		} else if a == 1 {
			print_all(all_students)
		} else if a == 2 {
			record(&all_students)
		} else if a == 3 {
			sort_by_grades(all_students)
		} else if a == 4 {
			sort_by_stipend(all_students)
		} else if a == 5 {
			save_to_file(all_students)
		}
	}
}
