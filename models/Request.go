package models

var operationTypes = []string{"+", "-", "*", "/"}

type Request struct {
	Number1 int
	Number2 int
	OperationType string
}

func ValidateOperationType(operationType string) bool {
	for i := 0; i < len(operationTypes); i++ {
		if operationType == operationTypes[i] {
			return true
		}
	}
	return false
}

