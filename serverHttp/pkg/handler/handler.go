package handler

import (
	"awesomeProject2/serverHttp"
	"awesomeProject2/serverHttp/pkg/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.eventsForDay)
	mux.HandleFunc("/events_for_week", h.eventsForWeek)
	mux.HandleFunc("/events_for_month", h.eventsForMonth)

	return mux

}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	var input serverHttp.Event

	if r.Method != http.MethodPost {
		http.Error(w, "Неправильный метод запроса", http.StatusBadRequest)
	}

	if err := ValidateCreateEventParams(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(input)
	id, err := h.services.CreateEvent(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/JSON")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	var input serverHttp.Event

	if r.Method != http.MethodPost {
		http.Error(w, "Неправильный метод запроса", http.StatusBadRequest)
	}

	if err := ValidateCreateEventParams(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := h.services.UpdateEvent(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonResponse, err := json.Marshal(input.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/JSON")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	var input serverHttp.Event

	if r.Method != http.MethodPost {
		http.Error(w, "Неправильный метод запроса", http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := h.services.DeleteEvent(input.Id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonResponse, err := json.Marshal(input.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/JSON")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *Handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	var input serverHttp.Event
	if r.Method != http.MethodGet {
		http.Error(w, "Неправильный метод запроса", http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eventsForDay(&input)
	events, err := h.services.EventsForDay(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonResponse, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/JSON")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func (h *Handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	var input serverHttp.Event
	if r.Method != http.MethodGet {
		http.Error(w, "Неправильный метод запроса", http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eventsForWeek(&input)
	events, err := h.services.EventsForDay(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonResponse, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/JSON")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *Handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	var input serverHttp.Event
	if r.Method != http.MethodGet {
		http.Error(w, "Неправильный метод запроса", http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	eventsForMonth(&input)
	events, err := h.services.EventsForDay(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonResponse, err := json.Marshal(events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/JSON")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
