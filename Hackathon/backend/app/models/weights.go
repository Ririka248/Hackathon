package models

import (
	"log"
	"time"
)

type Weight struct {
	ID          int
	Weightvalue float64
	UserID      int
	Date        time.Time
	CreatedAt   time.Time
}

func (u *User) CreateWeight(weightvalue float64, date time.Time) (err error) {
	cmd := `insert into weights (
		weightvalue, 
		user_id, 
		date,
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd, weightvalue, u.ID, date, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetWeight(id int) (weight Weight, err error) {
	cmd := `select id, weightvalue, user_id, created_at
	from weights where id = ?`
	weight = Weight{}

	err = Db.QueryRow(cmd, id).Scan(
		&weight.ID,
		&weight.Weightvalue,
		&weight.UserID,
		&weight.CreatedAt)
	return weight, err
}

func GetWeights() (weights []Weight, err error) {
	cmd := `select id, weightvalue, user_id, created_at
	from weights`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var weight Weight
		err = rows.Scan(
			&weight.ID,
			&weight.Weightvalue,
			&weight.UserID,
			&weight.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		weights = append(weights, weight)
	}
	rows.Close()

	return weights, err
}

func (u *User) GetWeightsByUser() (weights []Weight, err error) {
	cmd := `select id, weightvalue, user_id, date, created_at
	from weights where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var weight Weight
		err = rows.Scan(
			&weight.ID,
			&weight.Weightvalue,
			&weight.UserID,
			&weight.Date,
			&weight.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		weights = append(weights, weight)
	}
	rows.Close()

	return weights, err
}

func (u *User) GetLatestWeightByUser() (weight Weight, err error) {
	cmd := `select id, weightvalue, user_id, date, created_at
	from weights group by user_id = ? having max(date)`
	weight = Weight{}
	err = Db.QueryRow(cmd, u.ID).Scan(
		&weight.ID,
		&weight.Weightvalue,
		&weight.UserID,
		&weight.Date,
		&weight.CreatedAt)
	return weight, err
}

func (h *Weight) UpdateWeight() error {
	cmd := `update weights set weightvalue = ? , user_id = ? 
	where id = ?`
	_, err = Db.Exec(cmd, h.Weightvalue, h.UserID, h.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (h *Weight) DeleteWeight() error {
	cmd := `delete from weights where id = ?`
	_, err = Db.Exec(cmd, h.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
