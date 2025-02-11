package main

import (
	"fmt"
	"net/http"
)

func hSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Парсим форму
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Ошибка при обработке формы", http.StatusBadRequest)
			return
		}

		q1 := r.FormValue("q1")
		q2 := r.FormValue("q2")
		q3 := r.FormValue("q3")
		q4 := r.FormValue("q4")
		q5 := r.FormValue("q5")

		fmt.Fprintf(w, "<h1>Результаты тестирования:</h1>")
		fmt.Fprintf(w, "<p>Ответ на вопрос 1: %s</p>", q1)
		fmt.Fprintf(w, "<p>Ответ на вопрос 2: %s</p>", q2)
		fmt.Fprintf(w, "<p>Ответ на вопрос 3: %s</p>", q3)
		fmt.Fprintf(w, "<p>Ответ на вопрос 4: %s</p>", q4)
		fmt.Fprintf(w, "<p>Ответ на вопрос 5: %s</p>", q5)
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/submit", hSubmit)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
