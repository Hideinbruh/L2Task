package repository

import (
	"awesomeProject2/serverHttp"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateEvent(event *serverHttp.Event) (int, error) {
	var id int
	fmt.Printf("%+v\n", r)
	query := fmt.Sprintf("INSERT INTO calendar (description, date, title) values ($1, $2, $3) RETURNING id")
	row := r.db.QueryRow(query, event.Description, event.Date, event.Title)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) UpdateEvent(event *serverHttp.Event) error {
	query := fmt.Sprintf("UPDATE calendar SET description=$1, date=$2, title=$3 WHERE id=$4")
	_, err := r.db.Exec(query, event.Description, event.Date, event.Title, event.Id)
	return err
}

func (r *Repository) DeleteEvent(eventID int) error {
	query := fmt.Sprintf("DELETE FROM calendar WHERE id=$1")
	_, err := r.db.Exec(query, eventID)
	return err
}

func (r *Repository) EventsForDay(event *serverHttp.Event) ([]serverHttp.Event, error) {
	var events []serverHttp.Event
	query := fmt.Sprintf("SELECT * FROM calendar WHERE date between $1 and $2")
	err := r.db.Select(&events, query, event.MinDate, event.MaxDate)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *Repository) EventsForWeek(event *serverHttp.Event) ([]serverHttp.Event, error) {
	var events []serverHttp.Event
	query := fmt.Sprintf("SELECT * FROM calendar WHERE date between $1 and $2")
	err := r.db.Select(&events, query, event.MinDate, event.MaxDate)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *Repository) EventsForMonth(event *serverHttp.Event) ([]serverHttp.Event, error) {
	var events []serverHttp.Event
	query := fmt.Sprintf("SELECT * FROM calendar WHERE date between $1 and $2")
	err := r.db.Select(&events, query, event.MinDate, event.MaxDate)
	if err != nil {
		return nil, err
	}
	return events, nil
}
