package controllers

import (
	"backend/app/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/mypage", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		heights, _ := user.GetHeightsByUser()
		user.Heights = heights
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func heightNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "height_new")
	}
}

func heightSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		value := r.PostFormValue("value")
		f64value, _ := strconv.ParseFloat(value, 64)
		timeString := "2021-06-10T12:00:00+09:00"
		parsedTime, _ := time.Parse(time.RFC3339, timeString)
		if err := user.CreateHeight(f64value, parsedTime); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/mypage", 302)
	}
}

func edit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetHeight(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "edit")
	}
}

func update(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Println()
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		value := r.PostFormValue("value")
		f64value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Println(err)
		}
		h := &models.Height{ID: id, Heightvalue: f64value, UserID: user.ID}
		if err := h.UpdateHeight(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/mypage", 302)
	}
}

func delete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		h, err := models.GetHeight(id)
		if err != nil {
			log.Println(err)
		}
		if err := h.DeleteHeight(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/mypage", 302)
	}
}
