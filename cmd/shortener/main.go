package main

import (
	"net/http"
)

// Сервер должен предоставлять два эндпоинта: POST / и GET /{id}.<br/>
//Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с
//кодом 201 и сокращённым URL в виде текстовой строки в теле.<br/>
//Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL
//и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.<br/>
//Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400.<br/>
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		q := r.URL.Path
		if q == "/" {
			http.Error(w, "The id parameter is missing", http.StatusBadRequest)
			return
		}
		idParam := q[1:len(q)]

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>Hello, World</h1><br/><h2>" + idParam + "</h2>"))
		return
	} else if r.Method == http.MethodPost {

	}
	http.Error(w, "Bad request", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", HandleRequest)

	http.ListenAndServe(":8080", nil)
}
