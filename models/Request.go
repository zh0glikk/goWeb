package models

var operationTypes = []string{"+", "-", "*", "/"}

type Request struct {
	Number1 int `json:"number_1"`
	Number2 int `json:"number_2"`
	OperationType string `json:"operation_type"`
}

func ValidateOperationType(operationType string) bool {
	for i := 0; i < len(operationTypes); i++ {
		if operationType == operationTypes[i] {
			return true
		}
	}
	return false
}

