package controllers

import (
	"hello-world-api/domain/usecase"
	"hello-world-api/domain/usecase/impl"
	"net/http"
)

type NameController struct {
	useCase usecase.NameUseCaseInterface
}

func NewNameController() *NameController {
	return &NameController{
		useCase: impl.NewNameUseCase(),
	}

}

func (u *NameController) GetName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := r.URL.Query().Get("name")

	isValid, res := u.useCase.GetName(ctx, name)
	if isValid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"` + res + `"}`))
		return

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"Invalid Input"}`))
		return
	}
}
