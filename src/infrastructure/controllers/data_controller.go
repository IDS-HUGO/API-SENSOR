package controllers

import (
    "encoding/json"
    "go-hexagonal-api/src/application/usecases"
    "go-hexagonal-api/src/domain/entities"
    "net/http"
)

type DataController struct {
    UseCase *usecases.SendDataUseCase
}

func NewDataController(useCase *usecases.SendDataUseCase) *DataController {
    return &DataController{UseCase: useCase}
}

func (ctrl *DataController) HandlePostData(w http.ResponseWriter, r *http.Request) {
    var data entities.Data
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := ctrl.UseCase.Execute(&data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}