package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockNameUseCase struct {
	isValid bool
	res     string
}

func (m *mockNameUseCase) GetName(ctx context.Context, name string) (bool, string) {
	return m.isValid, m.res
}

func TestNameController_GetName(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		isValid    bool
		res        string
		wantStatus int
		wantBody   map[string]string
	}{
		{
			name:       "valid name",
			query:      "name=John",
			isValid:    true,
			res:        "Hello, John!",
			wantStatus: http.StatusOK,
			wantBody:   map[string]string{"message": "Hello, John!"},
		},
		{
			name:       "invalid name",
			query:      "name=",
			isValid:    false,
			res:        "",
			wantStatus: http.StatusBadRequest,
			wantBody:   map[string]string{"error": "Invalid Input"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUC := &mockNameUseCase{isValid: tt.isValid, res: tt.res}
			controller := NameController{useCase: mockUC}

			req := httptest.NewRequest(http.MethodGet, "/?"+tt.query, nil)
			rr := httptest.NewRecorder()

			controller.GetName(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, rr.Code)
			}

			var got map[string]string
			if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}
			if got["message"] != tt.wantBody["message"] && got["error"] != tt.wantBody["error"] {
				t.Errorf("expected body %v, got %v", tt.wantBody, got)
			}
		})
	}
}
