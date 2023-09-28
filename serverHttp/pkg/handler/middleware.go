package handler

import (
	"awesomeProject2/serverHttp"
	"fmt"
	"log"
	"net/http"
	"time"
)

func ValidateCreateEventParams(r *http.Request) error {
	title := r.FormValue("title")
	date := r.FormValue("date")
	description := r.FormValue("description")

	if title == "" || date == "" || description == "" {
		return fmt.Errorf("введите данные во все поля")
	}
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return fmt.Errorf("неправильный формат даты")
	}
	return nil
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func eventsForDay(event *serverHttp.Event) {
	year, month, day := event.ParsedDate.Date()
	event.MinDate = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	event.MaxDate = time.Date(year, month, day, 23, 59, 59, 999999999, time.UTC)
}

func eventsForWeek(event *serverHttp.Event) {
	WeekStart := weekBegin(event.ParsedDate.ISOWeek())
	event.MinDate = WeekStart
	event.MaxDate = WeekStart.AddDate(0, 0, 6)
}

func weekBegin(year, week int) time.Time {
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)
	if weekDay := t.Weekday(); weekDay == time.Sunday {
		t = t.AddDate(0, 0, 6)
	} else {
		t = t.AddDate(0, 0, -int(weekDay)+1)
	}
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)
	return t
}

func eventsForMonth(event *serverHttp.Event) {
	year, month, _ := event.ParsedDate.Date()
	event.MinDate = time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	event.MaxDate = time.Date(year, month+1, 0, 23, 59, 59, 999999999, time.UTC)
}
