package repository

import (
	"fmt"
	"github.com/TrackTasks/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type PeoplePostgres struct {
	db *sqlx.DB
}

func NewPeoplePostgres(db *sqlx.DB) *PeoplePostgres {
	return &PeoplePostgres{db: db}
}

func (r *PeoplePostgres) Create(newPeople models.People) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createPeopleQuery := fmt.Sprintf("INSERT INTO people (name, surname, address, passportnumber) VALUES ($1,$2,$3,$4) RETURNING id")
	row := tx.QueryRow(createPeopleQuery, newPeople.Name, newPeople.Surname, newPeople.Address, newPeople.PassportNumber)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *PeoplePostgres) Update(id int, updatePeople models.People) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if updatePeople.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, updatePeople.Name)
		argId++
	}
	if updatePeople.Surname != "" {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, updatePeople.Surname)
		argId++
	}
	if updatePeople.Address != "" {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, updatePeople.Address)
		argId++
	}
	if updatePeople.PassportNumber != 0 {
		setValues = append(setValues, fmt.Sprintf("passportnumber=$%d", argId))
		args = append(args, updatePeople.PassportNumber)
		argId++
	}

	setQuery := strings.Join(setValues, ",")

	//query := fmt.Sprintf("UPDATE people SET %s WHERE id = $%d", setQuery, argId)
	query := "UPDATE people SET "
	query = query + setQuery
	query += fmt.Sprintf(" WHERE id = $%d", argId)
	args = append(args, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PeoplePostgres) GetAll() ([]models.People, error) {
	var peoples []models.People

	query := fmt.Sprintf("SELECT * FROM people")
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query people: %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var person models.People
		if err := rows.Scan(&person.Id, &person.Name, &person.Surname, &person.Address, &person.PassportNumber); err != nil {
			return nil, fmt.Errorf("failed to scan people: %s", err)
		}
		peoples = append(peoples, person)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading rows: %s", err)
	}
	return peoples, nil
}

func (r *PeoplePostgres) GetWithFilters(id int, name string, surname string, address string, passportNumber int) ([]models.People, error) {
	var peoples []models.People
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if id != 0 {
		setValues = append(setValues, fmt.Sprintf("id=$%d", argId))
		args = append(args, id)
		argId++
	}

	if name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, name)
		argId++
	}
	if surname != "" {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, surname)
		argId++
	}
	if address != "" {
		setValues = append(setValues, fmt.Sprintf("address=$%d", argId))
		args = append(args, address)
		argId++
	}
	if passportNumber != 0 {
		setValues = append(setValues, fmt.Sprintf("passportNumber=$%d", argId))
		args = append(args, passportNumber)
		argId++
	}

	query := fmt.Sprintf("SELECT * FROM people WHERE 1 = 1")
	if len(setValues) > 0 {
		query += " AND " + strings.Join(setValues, " AND ")
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query people: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var person models.People
		if err := rows.Scan(&person.Id, &person.Name, &person.Surname, &person.Address, &person.PassportNumber); err != nil {
			return nil, fmt.Errorf("failed to scan people: %s", err)
		}
		peoples = append(peoples, person)
	}
	return peoples, nil
}

func (r *PeoplePostgres) Delete(id int) (int64, error) {
	query := fmt.Sprintf("DELETE FROM people WHERE id = $1")
	res, err := r.db.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete people: %s", err)
	}
	return res.RowsAffected()
}
