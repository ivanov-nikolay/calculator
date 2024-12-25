package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ivanov-nikolay/calculator/internal/model"
)

func TestGetCalculationBadRequestCase(t *testing.T) {
	body := `{""}`
	r := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetCalculation(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("GetCalculation returned wrong status code: got %v want %v", res.StatusCode, http.StatusBadRequest)
	}
}

func TestGetCalculationSuccessCase(t *testing.T) {
	body := `{"expression":"(5+ 2) /4+(2+2)*8+(4-3)*10"}`
	want := "43.75"

	r := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", strings.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetCalculation(w, r)

	res := w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code: status code should be 200")
	}
	result, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading expression: %v", err)
	}
	a := model.Response{}
	if err := json.Unmarshal(result, &a); err != nil {
		t.Errorf("error unmarshalling expression: %v", err)
	}
	if a.Result != want {
		t.Errorf("wrong calculation: want %s, got %s", want, string(result))
	}
}

func TestGetCalculationUnprocessableEntityCase(t *testing.T) {
	type args struct {
		expression string
		wantStatus int
		wantResult string
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "Empty expression body",
			args: args{
				expression: "",
				wantStatus: http.StatusBadRequest,
				wantResult: `{"error":"invalid expression JSON"}`,
			},
		},
		{
			name: "Empty expression",
			args: args{
				expression: `{"expression":""}`,
				wantStatus: http.StatusUnprocessableEntity,
				wantResult: `{"error":"expression is required"}`,
			},
		},
		{
			name: "Divided by zero",
			args: args{
				expression: `{"expression":"10/0"}`,
				wantStatus: http.StatusUnprocessableEntity,
				wantResult: `{"error":"divided by zero"}`,
			},
		},
		{
			name: "Input no ASCII",
			args: args{
				expression: `{"expression":"10/ф"}`,
				wantStatus: http.StatusUnprocessableEntity,
				wantResult: `{"error":"input no ASCII"}`,
			},
		},
		{
			name: "Missing value",
			args: args{
				expression: `{"expression":"10//2"}`,
				wantStatus: http.StatusUnprocessableEntity,
				wantResult: `{"error":"popValue: invalid expression"}`,
			},
		},
		{
			name: "Invalid expression",
			args: args{
				expression: `{"expression":"10/v2"}`,
				wantStatus: http.StatusUnprocessableEntity,
				wantResult: `{"error":"expression is not valid"}`,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", strings.NewReader(tc.args.expression))
			r.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			GetCalculation(w, r)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.args.wantStatus {
				t.Errorf("GetCalculation returned wrong status code: got %v want %v\n", res.StatusCode, tc.args.wantStatus)
			}

			result, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("failed to read result: %v\n", err)
			}
			if strings.TrimSpace(string(result)) != tc.args.wantResult {
				t.Errorf("GetCalculation returned wrong result: got %v want %v\n", string(result), tc.args.wantResult)
			}
		})
	}
}
