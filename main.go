package main

import (
	"encoding/json"
	"goWeb/models"
	"html/template"
	"log"
	"net/http"
)

func frontPageHander(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "front")
}

func renderTemplate(w http.ResponseWriter, templateName string) {
	t, _ := template.ParseFiles("views/" + templateName + ".html")

	t.Execute(w, "")
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	var d models.Request
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	isCorrect := models.ValidateOperationType(d.OperationType)

	if !isCorrect {
		a, err := json.Marshal("Wrong data")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(a)
		return
	}

	result := calculateResult(&d)

	a, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(a)
}

func calculateResult(r *models.Request) int{
	switch r.OperationType {
		case "+":
			return r.Number1 + r.Number2
		case "-":
			return r.Number1 - r.Number2
		case "*":
			return r.Number1 * r.Number2
		case "/":
			return r.Number1 / r.Number2
	}
	return 0
}

func main() {
	http.HandleFunc("/", frontPageHander)
	http.HandleFunc("/calc", calcHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}




