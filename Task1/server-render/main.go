package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageData := struct {
			Title   string
			Heading string
			Text    string
		}{
			Title:   "Это главная страница, пока что она одна =(",
			Heading: "Hellov. Hov are u?",
			Text:    "Это текст из файла main.go, можешь проверить в структуре. Он был создан, когда ты сделал запрос к сайту. Вот так то!",
		}

		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, "Ошибочка вышла: ", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, pageData)
	})

	fmt.Println("Server is running on 8080")

	http.ListenAndServe(":8080", nil)
}
