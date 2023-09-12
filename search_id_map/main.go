// Поиск человека по id

package main

import "fmt"

func searchId(id int, people map[int]string) (string, error) {
	name, found := people[id]
	if !found {
		return "", fmt.Errorf("Человек с id: %d не найден.", id)
	}
	return name, nil
}

func main() {

	mans := map[int]string{
		1: "Mark",
		2: "Yulia",
		3: "Aleksandr",
		4: "Jhon",
	}
	id := 1
	name, err := searchId(id, mans)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Найден человек с id: %d. Его имя: %s", id, name)
	}
}
