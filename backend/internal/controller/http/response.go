package http

import (
	"backend/internal/models"
	"encoding/json"
	"net/http"
)

type errData[T any] struct {
	Data       T      `json:"data"`
	StatusCode int    `json:"status"`
	Message    string `json:"msg"`
}

func SuccessBatchRespond(batch []models.Paste, w http.ResponseWriter) {
	payload := errData[[]models.Paste]{
		Data:       batch,
		StatusCode: http.StatusOK,
		Message:    "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SuccessPasteRespond(paste *models.Paste, w http.ResponseWriter) {
	payload := errData[*models.Paste]{
		Data:       paste,
		StatusCode: http.StatusOK,
		Message:    "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SuccessUsersRespond(users []models.User, w http.ResponseWriter) {
	payload := errData[[]models.User]{
		Data:       users,
		StatusCode: http.StatusOK,
		Message:    "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SuccessUserRespond(user *models.User, w http.ResponseWriter) {
	payload := errData[*models.User]{
		Data:       user,
		StatusCode: http.StatusOK,
		Message:    "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SuccessDelete(id string, w http.ResponseWriter) {
	payload := struct {
		Id string `json:"id"`
	}{Id: id}
	err := json.NewEncoder(w).Encode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
	}
}

func SuccessfulCreation(id string, w http.ResponseWriter) {
	payload := struct {
		Id string `json:"id"`
	}{Id: id}
	err := json.NewEncoder(w).Encode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
	}
}

func SuccessLogin(token string, w http.ResponseWriter) {
	payload := errData[string]{
		Data:       token,
		StatusCode: http.StatusOK,
		Message:    "success",
	}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SuccessLogout(w http.ResponseWriter) {
	payload := struct {
		Status string `json:"status"`
	}{Status: "success"}
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
