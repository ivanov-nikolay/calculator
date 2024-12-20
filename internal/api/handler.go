package api

import (
	"encoding/json"
	"fmt"
	"github.com/ivanov-nikolay/calculator/internal/model"
	"github.com/ivanov-nikolay/calculator/internal/service"
	"net/http"
)

// GetCalculation обработчик http-запроса вычисления арифметического выражения
func GetCalculation(w http.ResponseWriter, r *http.Request) {
	var resp model.Request

	err := json.NewDecoder(r.Body).Decode(&resp)
	if err != nil {
		sendJSONError(w, "invalid expression JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if resp.Expression == "" {
		sendJSONError(w, "expression is required", http.StatusUnprocessableEntity)
		return
	}

	result, err := service.Calculation(resp.Expression)
	if err != nil {
		sendJSONError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	answer := fmt.Sprintf("%.2f", result)

	response := model.Response{
		Result: answer,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		sendJSONError(w, "internal server error", http.StatusInternalServerError)
	}
}

// sendJSONError формирует json сообщение ошибки
func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
