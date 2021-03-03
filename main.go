package main

import (
	"encoding/json"
	logrus "github.com/sirupsen/logrus"
	"goWeb/models"
	"html/template"
	"net/http"
	"os"
)

var log = logrus.New()

func frontPageHander(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "front")
}

func renderTemplate(w http.ResponseWriter, templateName string) {
	t, _ := template.ParseFiles("views/" + templateName + ".html")

	t.Execute(w, "")
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	var request models.Request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	isCorrect := models.ValidateOperationType(request.OperationType)

	if !isCorrect {
		log.WithFields(logrus.Fields{
			"number_1": request.Number1,
			"number_2": request.Number2,
			"operation_type" : request.OperationType,
		}).Warnf("Wrong operation type '%s'", request.OperationType)

		a, err := json.Marshal("Wrong data")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(a)
		return
	}

	log.WithFields(logrus.Fields{
		"number_1": request.Number1,
		"number_2": request.Number2,
		"operation_type" : request.OperationType,
	}).Info("Correct data")

	result := calculateResult(&request)

	log.WithFields(logrus.Fields{
		"result" : result,
	}).Info("Result")

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

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

func main() {
	http.HandleFunc("/", frontPageHander)
	http.HandleFunc("/calc", calcHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}




