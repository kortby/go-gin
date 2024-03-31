package models

import (
	"time"
	"example/gingonic/db"
)

type Event struct {
    ID          int64     `json:"id"`
    Name        string    `json:"name" binding:"required"`
    Description string    `json:"description" binding:"required"`
    Location    string    `json:"location" binding:"required"`
    DateTime    time.Time `json:"datetime" binding:"required"`
    UserID      int       `json:"user_id"`
}


func (e *Event) Save() error {
    query := `INSERT INTO events(name, description, location, datetime, user_id)
              VALUES (?, ?, ?, ?, ?)`
    sqlStatement, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer sqlStatement.Close()

	ParsedDatetime := e.DateTime.Format("2006-01-02 15:04:05")

    result, err := sqlStatement.Exec(e.Name, e.Description, e.Location, ParsedDatetime, e.UserID)
    if err != nil {
        return err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return err
    }

    e.ID = id
    return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}
	
	return events, nil
}


func GetEventById(id int64) (*Event, error){
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	// defer rows.Close()

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, datetime = ? WHERE id = ?`
	sqlStatement, err := db.DB.Prepare(query)
    if err != nil {
		return err
    }
    defer sqlStatement.Close()
	
	_, err = sqlStatement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}
