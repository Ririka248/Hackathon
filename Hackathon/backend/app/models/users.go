package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	PassWord  string
	Maxscore  int
	Item1     int
	Item2     int
	Item3     int
	CreatedAt time.Time
	Heights   []Height
}

type Session struct {
	ID        int
	UUID      string
	Name      string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		password,
		maxscore,
		item1,
		item2,
		item3,
		created_at) values (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		Encrypt(u.PassWord),
		u.Maxscore,
		u.Item1,
		u.Item2,
		u.Item3,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, password, maxscore, item1, item2, item3, created_at
	from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.PassWord,
		&user.Maxscore,
		&user.Item1,
		&user.Item2,
		&user.Item3,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ? , maxscore = ? , item1 = ? , item2 = ? , item3 = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Maxscore, u.Item1, u.Item2, u.Item3, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByName(name string) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, password, maxscore, item1, item2, item3, created_at
	from users where name = ?`
	err = Db.QueryRow(cmd, name).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.PassWord,
		&user.Maxscore,
		&user.Item1,
		&user.Item2,
		&user.Item3,
		&user.CreatedAt)
	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions (
		uuid,
		name,
		user_id,
		created_at) values (?, ?, ?, ?)`

	_, err = Db.Exec(cmd1, createUUID(), u.Name, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, name, user_id, created_at
	from sessions where user_id = ? and name =?`

	err = Db.QueryRow(cmd2, u.ID, u.Name).Scan(
		&session.ID,
		&session.UUID,
		&session.Name,
		&session.UserID,
		&session.CreatedAt)

	return session, err
}

func (sess *Session) ChaeckSession() (valid bool, err error) {
	cmd := `select id, uuid, name, user_id, created_at
	from sessions where uuid = ?`

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Name,
		&sess.UserID,
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid =?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, created_at FROM users
	where id = ?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.CreatedAt)

	return user, err
}
