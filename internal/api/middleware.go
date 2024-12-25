package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ivanov-nikolay/calculator/internal/model"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

// Write записывает результат вычисления арифметического выражения в body (это переопределение метода - кастомизация)
func (lrw *loggingResponseWriter) Write(p []byte) (n int, err error) {
	lrw.body = append(lrw.body, p...)
	return lrw.ResponseWriter.Write(p)
}

// WriteHeader записывает код ответа сервера в header (это переопределение метода - кастомизация)
func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

// LoggingMiddleWare записывает события (логирует) в файл
func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("./logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Printf("failed to open log file: %v\n", err)
			sendJSONError(w, "internal server error", http.StatusInternalServerError)
			return
		}
		defer file.Close()
		log.SetOutput(file)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("failed to read r.Body: %v\n", err)
			sendJSONError(w, "Bad request", http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		expr := model.Request{}
		if err = json.Unmarshal(body, &expr); err != nil {
			log.Printf("failed to parse expression: %v\n", err)
			sendJSONError(w, "internal server error", http.StatusInternalServerError)
			return
		}

		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(lrw, r)

		if len(lrw.body) == 0 {
			log.Println("response body is empty")
			return
		}

		var response map[string]string
		if err := json.Unmarshal(lrw.body, &response); err != nil {
			log.Printf("failed to parse response expression: %v\n", err)
		}

		answer, ok := response["result"]
		if !ok {
			log.Printf("the key 'result' is missing from the response: %v\n", response)
		}

		if expr.Expression != "" {
			if ok && answer != "" {
				log.Printf("%s = %s\n", expr.Expression, answer)
			} else {
				log.Printf("%s = processing error\n", expr.Expression)
			}
		} else {
			log.Println("the query does not contain an expression")
		}
	})
}
