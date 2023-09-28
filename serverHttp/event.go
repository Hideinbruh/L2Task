package serverHttp

import "time"

type Event struct {
	Id          int    `json:"-"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date        string `json:"date" binding:"required"`
	ParsedDate  time.Time
	MinDate     time.Time
	MaxDate     time.Time
}

type ResultEvent struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Date        string `json:"date"`
	Done        bool   `json:"done" db:"done"`
}
