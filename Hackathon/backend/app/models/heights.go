package models

import (
	"log"
	"time"
)

type Height struct {
	ID             int
	Heightvalue    float64
	NowHeightValue float64
	UserID         int
	Date           time.Time
	CreatedAt      time.Time
}

func (u *User) CreateHeight(heightvalue float64, date time.Time) (err error) {
	cmd := `insert into heights (
		heightvalue, 
		user_id, 
		date,
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd, heightvalue, u.ID, date, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetHeight(id int) (height Height, err error) {
	cmd := `select id, heightvalue, user_id, created_at
	from heights where id = ?`
	height = Height{}

	err = Db.QueryRow(cmd, id).Scan(
		&height.ID,
		&height.Heightvalue,
		&height.UserID,
		&height.CreatedAt)
	return height, err
}

func GetHeights() (heights []Height, err error) {
	cmd := `select id, heightvalue, user_id, created_at
	from heights`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var height Height
		err = rows.Scan(
			&height.ID,
			&height.Heightvalue,
			&height.UserID,
			&height.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		heights = append(heights, height)
	}
	rows.Close()

	return heights, err
}

func (u *User) GetHeightsByUser() (heights []Height, err error) {
	cmd := `select id, heightvalue, user_id, date, created_at
	from heights where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var height Height
		err = rows.Scan(
			&height.ID,
			&height.Heightvalue,
			&height.UserID,
			&height.Date,
			&height.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		heights = append(heights, height)
	}
	rows.Close()

	return heights, err
}

func (u *User) GetLatestHeightByUser() (height Height, err error) {
	cmd := `select id, heightvalue, user_id, date, created_at
	from heights group by user_id = ? having max(date)`
	height = Height{}
	err = Db.QueryRow(cmd, u.ID).Scan(
		&height.ID,
		&height.Heightvalue,
		&height.UserID,
		&height.Date,
		&height.CreatedAt)
	return height, err
}

func (h *Height) UpdateHeight() error {
	cmd := `update heights set heightvalue = ? , user_id = ? 
	where id = ?`
	_, err = Db.Exec(cmd, h.Heightvalue, h.UserID, h.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (h *Height) DeleteHeight() error {
	cmd := `delete from heights where id = ?`
	_, err = Db.Exec(cmd, h.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
