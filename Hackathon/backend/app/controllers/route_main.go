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
		generateHTML(w, nil, "layout", "public_navbar", "top")
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

func eat(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "eating_manage")
	}
}

func sleep(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "sleep")
	}
}

func train(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		sports, _ := models.GetSports()
		user.Sports = sports
		generateHTML(w, user, "layout", "private_navbar", "training")
	}
}

func privateexplanation(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "explanation")
	}
}

func publicexplanation(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "explanation")
}

func setting(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		heights, _ := user.GetHeightsByUser()
		weights, _ := user.GetWeightsByUser()
		user.Heights = heights
		user.Weights = weights
		generateHTML(w, user, "layout", "private_navbar", "setting")
	}
}

func heightNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
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

		http.Redirect(w, r, "/mypage/setting", 302)
	}
}

func heightedit(w http.ResponseWriter, r *http.Request, id int) {
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
		generateHTML(w, t, "layout", "private_navbar", "heightedit")
	}
}

func heightupdate(w http.ResponseWriter, r *http.Request, id int) {
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
		http.Redirect(w, r, "/mypage/setting", 302)
	}
}

func heightdelete(w http.ResponseWriter, r *http.Request, id int) {
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
		http.Redirect(w, r, "/mypage/setting", 302)
	}
}

func weightNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "weight_new")
	}
}

func weightSave(w http.ResponseWriter, r *http.Request) {
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
		if err := user.CreateWeight(f64value, parsedTime); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/mypage/setting", 302)
	}
}

func weightedit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		t, err := models.GetWeight(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "weightedit")
	}
}

func weightupdate(w http.ResponseWriter, r *http.Request, id int) {
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
		h := &models.Weight{ID: id, Weightvalue: f64value, UserID: user.ID}
		if err := h.UpdateWeight(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/mypage/setting", 302)
	}
}

func weightdelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		h, err := models.GetWeight(id)
		if err != nil {
			log.Println(err)
		}
		if err := h.DeleteWeight(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/mypage/setting", 302)
	}
}

func game(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, user, "layout", "private_navbar", "game")
	}
}

func calendar(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, user, "layout", "private_navbar", "calendar")
	}
}

func itemcheck(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, user, "layout", "private_navbar", "item")
	}
}
