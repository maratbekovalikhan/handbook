package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *ArticleService
}

func NewHandler(s *ArticleService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Articles(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {

		data := h.service.GetAll()
		json.NewEncoder(w).Encode(data)
		return
	}

	if r.Method == http.MethodPost {

		var a Article

		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			http.Error(w, "Invalid JSON", 400)
			return
		}

		result := h.service.Create(a)

		json.NewEncoder(w).Encode(result)
		return
	}

	http.Error(w, "Method not allowed", 405)
}

func (h *Handler) Article(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/articles/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	if r.Method == http.MethodGet {

		a, ok := h.service.GetByID(id)

		if !ok {
			http.Error(w, "Not found", 404)
			return
		}

		json.NewEncoder(w).Encode(a)
		return
	}

	if r.Method == http.MethodPut {

		var a Article

		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			http.Error(w, "Invalid JSON", 400)
			return
		}

		res, ok := h.service.Update(id, a)

		if !ok {
			http.Error(w, "Not found", 404)
			return
		}

		json.NewEncoder(w).Encode(res)
		return
	}

	if r.Method == http.MethodDelete {

		ok := h.service.Delete(id)

		if !ok {
			http.Error(w, "Not found", 404)
			return
		}

		w.WriteHeader(204)
		return
	}

	http.Error(w, "Method not allowed", 405)
}
