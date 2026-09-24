package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"time"
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
		"Институт: %s\nСтипендия: %.2f\nСредний балл: %.2f",
		st.ID, st.name, st.databirthday, st.institute, st.stipend, st.GPA)
}

// функция для генерации id
func max_id(students_list []students) int {
	if len(students_list) == 0 {
		return 0
	}
	max := slices.MaxFunc(students_list, func(a, b students) int {
		return cmp.Compare(a.ID, b.ID)
	})
	return max.ID
}

// Проверка строки на соответствие шаблону даты
func is_valid_date(dateStr string) bool {
	// Шаблон: 02 - день, 01 - месяц, 2006 - год
	_, err := time.Parse("2.1.2006", dateStr)
	return err == nil
}

// функция добавления студента в срез
func record(student *[]students, reader *bufio.Reader) {
	id := max_id(*student) + 1

	fmt.Print("Укажите имя студента: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Удаляем символ переноса строки

	fmt.Print("Укажите дату рождения студента в формате ДД.ММ.ГГГГ: ")
	data, _ := reader.ReadString('\n')
	data = strings.TrimSpace(data)
	for !is_valid_date(data) {
		fmt.Print("Ошибка ввода даты. Введите еще раз (ДД.ММ.ГГГГ): ")
		data, _ = reader.ReadString('\n')
		data = strings.TrimSpace(data)
	}

	fmt.Print("Укажите институт студента: ")
	inst, _ := reader.ReadString('\n')
	inst = strings.TrimSpace(inst)

	fmt.Print("Укажите стипендию студента: ")
	stiStr, _ := reader.ReadString('\n')
	stiStr = strings.TrimSpace(stiStr)
	sti, err := strconv.ParseFloat(stiStr, 64)
	for err != nil || sti < 0 {
		fmt.Print("Ошибка ввода стипендии. Введите еще раз неотрицательное число: ")
		stiStr, _ = reader.ReadString('\n')
		stiStr = strings.TrimSpace(stiStr)
		sti, err = strconv.ParseFloat(stiStr, 64)
	}

	fmt.Print("Укажите средний балл студента: ")
	gpaStr, _ := reader.ReadString('\n')
	gpaStr = strings.TrimSpace(gpaStr)
	GPA, err := strconv.ParseFloat(gpaStr, 64)
	for err != nil || GPA < 2 || GPA > 5 {
		fmt.Print("Ошибка ввода среднего балла. Введите число от 2 до 5: ")
		gpaStr, _ = reader.ReadString('\n')
		gpaStr = strings.TrimSpace(gpaStr)
		GPA, err = strconv.ParseFloat(gpaStr, 64)
	}

	*student = append(*student, students{id, name, data, inst, sti, GPA})

	f, err := os.OpenFile("Base.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}
	defer f.Close()
	fmt.Fprintf(f, "%d/%s/%s/%s/%.2f/%.2f\n", id, name, data, inst, sti, GPA)
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
		}
		fmt.Println(student[i].ToString())
		fmt.Println("---")
	}
	if len(student) == 0 {
		fmt.Println("---")
	}
	fmt.Println("Всего студентов:", len(student))
	fmt.Println("---")
}

// функция для сохранения измененнного слайса в файл
func save_to_file(students_list []students) error {
	f, err := os.Create("Base.tmp")
	if err != nil {
		fmt.Println("Ошибка записи!")
		return err
	}
	w := bufio.NewWriter(f)
	for _, student := range students_list {
		fmt.Fprintf(w, "%d/%s/%s/%s/%.2f/%.2f\n",
			student.ID, student.name, student.databirthday, student.institute, student.stipend, student.GPA)
	}
	if err := w.Flush(); err != nil {
		f.Close()
		os.Remove("Base.tmp")
		fmt.Println("Ошибка записи!")
		return err
	}
	f.Close()
	if err := os.Rename("Base.tmp", "Base.txt"); err != nil {
		os.Remove("Base.tmp")
		fmt.Println("Ошибка записи!")
		return err
	}
	fmt.Println("Изменения сохранены!")
	return nil
}

// функция для очистки экрана (работает и на Windows, и на macOS/Linux)
func clear_screen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// основная функция
func main() {
	all_students := []students{}

	file, err := os.Open("Base.txt")
	if err != nil {
		f, err := os.Create("Base.txt")
		if err != nil {
			fmt.Println("Не удалось создать файл:", err)
			return
		}
		f.Close()
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fields := strings.Split(scanner.Text(), "/")
			if len(fields) < 6 {
				continue
			}
			id, err1 := strconv.Atoi(fields[0])
			sti, err2 := strconv.ParseFloat(fields[4], 64)
			gpa, err3 := strconv.ParseFloat(fields[5], 64)
			if err1 != nil || err2 != nil || err3 != nil {
				fmt.Println("Пропущена некорректная строка:", scanner.Text())
				continue
			}
			all_students = append(all_students, students{id, fields[1], fields[2], fields[3], sti, gpa})
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка чтения:", err)
		}
		file.Close()
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Для начала работы пожалуйста, нажмите Enter...")
		reader.ReadString('\n')
		clear_screen()

		fmt.Println("Меню команд:")
		fmt.Println("1 - Показать список всех студентов")
		fmt.Println("2 - Запись студента в базу данных")
		fmt.Println("3 - Сортировать студентов по среднему баллу (по убыванию)")
		fmt.Println("4 - Сортировать студентов по размеру стипендии (по убыванию)")
		fmt.Println("5 - Сохранить изменения в файл")
		fmt.Println("0 - Завершить работу")
		fmt.Println("Выберите пункт из меню управления: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		a, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Некорректный ввод! Пожалуйста, введите число.")
			continue
		}

		switch a {
		case 0:
			return
		case 1:
			print_all(all_students)
		case 2:
			record(&all_students, reader)
		case 3:
			sort_by_grades(all_students)
		case 4:
			sort_by_stipend(all_students)
		case 5:
			if err := save_to_file(all_students); err != nil {
				fmt.Println("Не удалось сохранить:", err)
			}
		default:
			fmt.Println("Нет такого пункта меню!")
		}
	}
}
