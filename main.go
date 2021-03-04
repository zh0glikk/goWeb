package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"goWeb/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

func frontPageHander(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "front")
}

func renderTemplate(w http.ResponseWriter, templateName string) {
	t, _ := template.ParseFiles("views/" + templateName + ".html")

	t.Execute(w, "")
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	localLog := r.Context().Value("log").(*logrus.Logger)

	defer localLog.WithFields(logrus.Fields{
		"Duration" : time.Since(startTime),
	}).Info("Handling request")

	var request models.Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	isCorrect := models.ValidateOperationType(request.OperationType)

	if !isCorrect {
		a, err := json.Marshal("Wrong data")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(a)
		return
	}

	result := calculateResult(&request)

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
	m := http.NewServeMux()


	m.HandleFunc("/", frontPageHander)
	m.HandleFunc("/calc", calcHandler)

	wr := models.NewLogger(m)


	log.Fatal(http.ListenAndServe(":8080", wr))
}




